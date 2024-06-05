package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

var db *sql.DB

type User struct {
	ID       string ``
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {

	dsn := "root:test@123!@tcp(127.0.0.1:3306)/golang_microservice"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	fmt.Println("connected")

	router := gin.Default()

	router.POST("/register", register)

	router.Run(":8080")
}
func register(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	//generate new uuid for the userID
	newUser.ID = uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash the password"})
		return
	}

	_, err = db.Exec("INSERT INTO users(id,username,email,password) VALUES (?, ?, ?, ?)", newUser.ID, newUser.Name, newUser.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":  "User created successfully",
		"id":       newUser.ID,
		"password": newUser.Password,
	})

}
