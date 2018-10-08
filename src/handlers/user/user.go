package user

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/guizot/go-gin-mysql/config"
	model_user "github.com/guizot/go-gin-mysql/src/models/user"
)

// Get DB from Mysql Config
func MysqlConfig() *sql.DB {
	db, err := config.GetMysql()
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// Get All User Endpoint
func GetAllUser(c *gin.Context) {
	db := *MysqlConfig()
	fmt.Println("MYSQL RUNNING: ", db)
	defer db.Close()

	sql := "SELECT * FROM " + "users"
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
		"message": "Success Get All Users",
		"user":    &users,
	})
}

// Get User Endpoint
func GetUser(c *gin.Context) {
	db := *MysqlConfig()
	fmt.Println("MYSQL RUNNING: ", db)
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

	c.JSON(200, gin.H{
		"message": "Success Get Users",
		"user":    &user,
	})
}
