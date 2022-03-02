package main

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name      string         `json:"name"`
	Interests datatypes.JSON `json:"interests"`
}

type Gift struct {
	gorm.Model
	Name       string         `json:"name"`
	Categories datatypes.JSON `json:"categories"`
}
