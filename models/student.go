package models

import (
    "github.com/jinzhu/gorm"
)

type Student struct {
    gorm.Model
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Grade string `json:"grade"`
}
