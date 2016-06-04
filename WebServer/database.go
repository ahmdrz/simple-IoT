package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type JSON struct {
	Success   bool        `json:"success"`
	ErrorCode int8        `json:"errorcode"`
	Result    interface{} `json:"result"`
}

type Light struct {
	ID       int    `json:"id"`
	HostID   int    `json:"host"`
	Name     string `json:"name"`
	Value    int    `json:"currentvalue"`
	MaxValue int    `json:"maxvalue"`
	LastUse  int64  `json:"lastuse"`
}

type User struct {
	Token  string
	Name   string
	HostID int
}

type Lights []Light

var db *gorm.DB
var err error

func connectToDatabase() {
	db, err = gorm.Open("sqlite3", "database.db")
	if err != nil {
		log.Println(err)
		panic("Error in database connection")
	}
	db.LogMode(false)
	db.CreateTable(&Light{})
	db.CreateTable(&User{})
}

func isUserExists(token string) bool {
	user := User{}
	db.Table("users").Where("token = ?", token).Find(&user)
	return user.HostID > 0
}

func getList(token string) []Light {
	var list []Light
	db.Table("lights").Where("host_id = ?", getHostID(token)).Find(&list)
	return list
}

func getHostID(token string) int {
	var result User
	db.Table("users").Where("token = ?", token).Find(&result)
	return result.HostID
}

func getLight(token string, id int) Light {
	var result Light
	db.Table("lights").Where("host_id = ?", getHostID(token)).Where("id = ? ", id).Find(&result)
	return result
}

func isIDExists(id int) bool {
	var result Light
	db.Table("lights").Where("id = ?", id).Find(&result)
	return (result.ID > 0)
}

func updateLight(token string, light Light) {
	db.Table("lights").Model(&light).Update("max_value", light.MaxValue).Update("value", light.Value)
}
