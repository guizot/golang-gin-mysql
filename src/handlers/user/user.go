package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guizot/go-gin-mysql/config"
	model_user "github.com/guizot/go-gin-mysql/src/models/user"
)

// Constant Variable
const DateFormat = "20060102150405"

// Get DB from Mysql Config
func MysqlConfig() *sql.DB {
	db, err := config.GetMysql()
	if err != nil {
		fmt.Println("Error Get Database")
	}
	fmt.Println("MySQL is Running..")
	return db
}

// Get All User Endpoint
func GetAllUser(c *gin.Context) {
	db := *MysqlConfig()
	defer db.Close()

	sql := "SELECT * FROM users"
	results, err := db.Query(sql)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get All User",
		})
		return
	}

	var user model_user.User
	var users model_user.Users
	for results.Next() {
		err = results.Scan(&user.Id, &user.Name, &user.Address, &user.Age, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "Error Scan User",
			})
			return
		}
		users = append(users, user)
	}

	c.JSON(200, gin.H{
		"message": "Success Get All User",
		"user":    &users,
	})
}

// Get User Endpoint
func GetUser(c *gin.Context) {
	db := *MysqlConfig()
	defer db.Close()

	id := c.Param("id")
	sql := "SELECT * FROM users where id = " + id
	results, err := db.Query(sql)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get User",
		})
		return
	}

	var user model_user.User
	for results.Next() {
		err = results.Scan(&user.Id, &user.Name, &user.Address, &user.Age, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "Error Scan User",
			})
			return
		}
	}

	if user.Id == 0 {
		c.JSON(200, gin.H{
			"message": "Success Get User",
			"user":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success Get User",
		"user":    &user,
	})
}

// Create User Endpoint
func CreateUser(c *gin.Context) {
	db := *MysqlConfig()
	defer db.Close()

	user := model_user.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}
	user.CreatedAt = time.Now().Format(DateFormat)
	user.UpdatedAt = time.Now().Format(DateFormat)

	sql := fmt.Sprintf("INSERT INTO users VALUES(?, ?, ?, ?, ?, ?)")
	insert, err := db.Query(sql, user.Id, user.Name, user.Address, user.Age, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Insert User",
		})
		return
	}

	defer insert.Close()

	c.JSON(200, gin.H{
		"message": "Success Create Users",
		"user":    &user,
	})
}
