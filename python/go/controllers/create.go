package controllers

import (
    "errors"
	"gorm.io/gorm"
    "crud/models"
)

func CreateUser(db *gorm.DB, user models.User) (int64, error) {
    result := db.Create(&user)
    if result.RowsAffected == 0 {
        return 0, errors.New("user not created")
    }
    return result.RowsAffected, nil
}