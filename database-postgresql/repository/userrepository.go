package user_repository

import (
	"errors"
	database_postgresql "user-api-rest/database-postgresql"
	"user-api-rest/database-postgresql/model"
)

func Save(user model.User) error {
	storedUser := database_postgresql.Connection_db.Create(&user)
	if storedUser.Error != nil {
		return errors.New("Error saving user")
	}
	return nil
}

func FindAll() []model.User {
	var users[] model.User
	database_postgresql.Connection_db.Find(&users)
	return users
}

func FindById(id string) model.User {
	var user model.User
	database_postgresql.Connection_db.Find(&user, id)
	return user
}

func Delete(id string) error {
	user := FindById(id)
	if user.ID == 0 {
		return errors.New("User not found")
	}
	database_postgresql.Connection_db.Unscoped().Delete(&user)
	return nil
}
