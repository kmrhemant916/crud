package controllers

import (
    "errors"
	"gorm.io/gorm"
    "crud/models"
)

func DeleteUser(db *gorm.DB, id string) (int64, error) {
    var deletedUser models.User
    result := db.Where("id = ?", id).Delete(&deletedUser)
    if result.RowsAffected == 0 {
        return 0, errors.New("user data not update")
    }
    return result.RowsAffected, nil
}