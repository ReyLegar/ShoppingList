package postgres

import (
	"fmt"
	"time"

	"github.com/ReyLegar/ShoppingList/internal/models"
)

type UserDatabase struct {
	database *Database
}

func (u *UserDatabase) Register(user *models.User) (int64, error) {
	sqlStatemant := `INSERT INTO users (name, surname, login, password) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64
	if err := u.database.db.Ping(); err != nil {
		fmt.Println("Не удалось подключиться к базе данных:", err)
		return -1, err
	}
	err := u.database.db.QueryRow(sqlStatemant, user.Name, user.Surname, user.Login, user.Password).Scan(&id)

	if err != nil {
		fmt.Println("Ошибка", err)
		return -1, err
	}
	return id, nil
}

var MaxRefreshTokenPerUser = 5

func (u *UserDatabase) Login(cred *models.Credentials) models.User {

	var user models.User

	sqlStatemant := `SELECT id, password FROM users WHERE login = $1`

	err := u.database.db.QueryRow(sqlStatemant, cred.Login).Scan(&user.ID, &user.Password)

	if err != nil {
		panic(err)
	}

	return user
}

func (u *UserDatabase) GetIDFromLogin() int64 {

	var id *int64

	sqlStatemant := `SELECT id FROM users WHERE login = $1`

	err := u.database.db.QueryRow(sqlStatemant, u.Login).Scan(&id)

	if err != nil {
		panic(err)
	}

	return *id

}

func (u *UserDatabase) CreateRefreshToken(user *models.User, refreshTokenString string, refreshExpirationTime time.Time) {

	_, err := u.database.db.Exec("INSERT INTO refresh_sessions (user_id, refresh_token, expires_at) VALUES ($1, $2, $3)", user.ID, refreshTokenString, refreshExpirationTime)
	if err != nil {
		return
	}
}
