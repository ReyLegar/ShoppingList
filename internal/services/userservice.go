package services

import (
	"time"

	"github.com/ReyLegar/ShoppingList/internal/database"
	"github.com/ReyLegar/ShoppingList/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("pass_secret_one")

type UserService interface {
	RegisterUser(user *models.User) (int64, error)
	LoginUser(cred *models.Credentials) map[string]string
}

type userServiceImpl struct {
	db database.Repository
}

func NewUserService(db database.Repository) UserService {
	return &userServiceImpl{db: db}
}

func (u *userServiceImpl) RegisterUser(user *models.User) (int64, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {

		return -1, err
	}

	user.Password = string(hashedPassword)
	id, err := u.db.User().Register(user)
	return id, err
}

func (u *userServiceImpl) LoginUser(cred *models.Credentials) map[string]string {
	user := u.db.User().Login(cred)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))

	if err != nil {
		return nil
	}

	accessExpirationTime := time.Now().Add(15 * time.Minute)

	accessClaims := &models.Claims{
		Login: cred.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpirationTime),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(JwtKey)
	if err != nil {
		return nil
	}

	refreshTokenUUID, err := uuid.NewRandom()
	if err != nil {
		return nil
	}
	refreshTokenString := refreshTokenUUID.String()
	refreshExpirationTime := time.Now().Add(7 * 24 * time.Hour) // 7 дней

	u.db.User().CreateRefreshToken(&user, refreshTokenString, refreshExpirationTime)

	return map[string]string{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	}
}

/* func (u *userServiceImpl) GetIDFromLogin() int64 {
	id := u.db.User().GetIDFromLogin()

	return id
}
*/
