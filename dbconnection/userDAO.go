package dbconnection

import "example/login/model"

func AddUser(user model.User) (int64, uint, error) {
	result := GetDBConnection().Create(&user)
	if result.Error != nil {
		return 0, 0, result.Error
	}
	lastInsertedId := user.ID
	return result.RowsAffected, lastInsertedId, result.Error
}

func GetUser(userid uint) (model.User, error) {
	user := model.User{}
	result := GetDBConnection().First(&user, userid)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, result.Error
}

func GetUserByEmail(email string) (model.User, error) {
	user := model.User{}
	result := GetDBConnection().Where("email = ?", email).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, result.Error
}
