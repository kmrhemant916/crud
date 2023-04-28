package main

import (
    // "errors"
    "fmt"
    "log"
    "crud/config"
    // "gorm.io/gorm"
    "crud/models"
    "crud/controllers"
)

func main() {
    db, err := config.Setup()
    if err != nil {
        log.Panic(err)
        return
    }
    fmt.Println("Connected")
    db.AutoMigrate(models.User{})
    fmt.Println("Migrated")
    user := models.User{
        ID: "zz",
        Name:        "User for item #1",
        Password:      "PENDING",
    }
    result, err := controllers.CreateUser(db, user)
    if err != nil {
        log.Panic(err)
        return
    }
    fmt.Println("User created", result)
    id := "zz"
    payment, _ := controllers.ReadUserWIthId(db, id)
    fmt.Println("Your user is", payment.Name)
    updatedUser, _ := controllers.UpdateUser(db, id, models.User{
        Name: "PAID",
    })
    fmt.Println("Your payment status now is ", updatedUser)
    controllers.DeleteUser(db, id)
    fmt.Println("Your user now is deleted")
}