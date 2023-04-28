package models

import (
    "time"
)

type User struct {
    ID          string
    Name        string
    Password    string
    Created     time.Time `gorm:"autoCreateTime"`
    Updated     time.Time `gorm:"autoUpdateTime"`
}