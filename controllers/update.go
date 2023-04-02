package controllers

import (
    "errors"
	"gorm.io/gorm"
    "crud/models"
)

func UpdateUser(db *gorm.DB, id string, user models.User) (models.User, error) {
    var updateUser models.User
    result := db.Model(&updateUser).Where("id = ?", id).Updates(user)
    if result.RowsAffected == 0 {
        return models.User{}, errors.New("user data not update")
    }
    return updateUser, nil
}