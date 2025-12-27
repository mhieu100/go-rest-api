package model

import "time"

type User struct {
    ID        uint      `gorm:"column:id;primaryKey"`
    Email     string    `gorm:"column:email"`
    Name      string    `gorm:"column:name"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
    return "users"
}
