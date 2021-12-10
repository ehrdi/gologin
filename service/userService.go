package service

import (
	"example/login/dbconnection"
	"example/login/model"
)

func AddUser(email string, password string) (model.User, error) {
	user := model.User{Email: email, Password: password}
	rowsAffected, lastInsertedId, err := dbconnection.AddUser(user)
	if err == nil && rowsAffected > 0 {
		user.ID = lastInsertedId
	}
	return user, err
}

func GetUser(userid uint) (model.User, error) {
	user, err := dbconnection.GetUser(userid)
	return user, err
}

func GetUserByEmail(email string) (model.User, error) {
	user, err := dbconnection.GetUserByEmail(email)
	return user, err
}
