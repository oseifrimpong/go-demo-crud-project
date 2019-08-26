package main

import (
	"net/http"

	"github.com/obed/demo-project/config"
	"github.com/obed/demo-project/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := NewRouter()
	r.Run(":8080")
}

// Router
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/createUser", createUser)
	return router
}

func createUser(c *gin.Context) {
	var newUser models.User
	c.ShouldBindJSON(&newUser)
	if config.Database.Create(&newUser).Error != nil {
		c.SecureJSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"Status": "User registered Successfully"})
	}
}

// func SignUp(c *gin.Context) {

// 	c.JSON(http.StatusOK, gin.H{"Status": "User registered Successfully"})
// }
