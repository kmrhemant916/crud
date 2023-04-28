package controllers

import (
    "errors"
	"gorm.io/gorm"
    "crud/models"
)

func ReadUserWIthId(db *gorm.DB, id string) (models.User, error) {
    var user models.User
    result := db.First(&user, "id = ?", id)
    if result.RowsAffected == 0 {
        return models.User{}, errors.New("user data not found")
    }
    return user, nil
}