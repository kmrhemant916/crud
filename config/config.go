package config

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

const (
    Host     = "127.0.0.1"
    User     = "root"
    Password = "test"
    Name     = "users"
    Port     = "3306"
)

func Setup() (*gorm.DB, error) {
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		User,
        Password,
        Host,
        Port,
        Name,
    )

    db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}