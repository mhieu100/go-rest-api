package model

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Email     string    `gorm:"uniqueIndex;not null"`
    Name      string    `gorm:"size:255"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
