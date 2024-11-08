package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
	ID uint 'json:"id" gorm:"primaryKey"'  // primaryKey makes this a unique id
	Name string 'json:"name"' // name of user
	Username string 'json:"username"' // username of user 
	Email string 'json:"email"' // email of user
}

var db *gorm.DB 
var err error
	
	func main () {
		router := gin.Default()

		// intializing database

		db, err = gom.Open(sqlite.Open("test.db").&gorm.Config{})
		if err != nil {
			panic("failed to open database")
		}
	}

	db.AutoMigrate(&Users{}) //auto-migrate User struct to update/create the database schema

	//routes
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUserByID)
	router.PUT("/users/:id", updateUser)
	router.DELETE ("/Users/:id", deleterUser)

	// run server 
	router.RUN("localhost:8080")