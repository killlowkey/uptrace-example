package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"colum:name;size:255;not null;unique"`
	Age  int    `gorm:"colum:age;default:18"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ToJson() string {
	data, _ := json.Marshal(u)
	return string(data)
}
