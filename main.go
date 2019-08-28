package main

import (
	"github.com/obed/demo-project/userAuth/delivery"
	"github.com/obed/demo-project/userAuth/repository"
	service "github.com/obed/demo-project/userAuth/services"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	var logger *zap.Logger

	defer logger.Sync()

	r := gin.New()
	repo := repository.NewStdRepository(logger)
	service := service.NewStdService(repo, logger)
	delivery.NewGinDelivery(r, service, logger)

	logger.Fatal(r.Run(":8080").Error())

	// comment below
	// r := NewRouter()

}

// Router
// func NewRouter() *gin.Engine {
// 	router := gin.Default()
// 	router.POST("/createUser", createUser)
// 	return router
// }

// func createUser(c *gin.Context) {
// 	var newUser model.User
// 	c.ShouldBindJSON(&newUser)
// 	if config.Database.Create(&newUser).Error != nil {
// 		c.SecureJSON(http.StatusInternalServerError,
// 			gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
// 	} else {
// 		c.SecureJSON(http.StatusOK, gin.H{"Status": "User registered Successfully"})
// 	}
// }

// func SignUp(c *gin.Context) {

// 	c.JSON(http.StatusOK, gin.H{"Status": "User registered Successfully"})
// }
