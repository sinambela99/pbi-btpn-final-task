package main

import (
	"crowdfunding/auth"
	"crowdfunding/handler"
	"crowdfunding/helper"
	"crowdfunding/photo"
	"crowdfunding/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=ian password=12345 dbname=btpn port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to Database Successful")

	userRepository := user.NewRepository(db)
	photoRepository := photo.NewRepository(db)

	userService := user.NewService(userRepository)
	photoService := photo.NewService(photoRepository)
	
	authService := auth.NewService()
	
	userHandler := handler.NewUserHandler(userService, authService)
	photoHandler := handler.NewPhotoHandler(photoService)
	
	router := gin.Default()
	router.Use(cors.Default())
	
	api := router.Group("api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.PUT("/update/:id", authMiddleware(authService, userService), userHandler.UpdateUser)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)
	api.DELETE("/user/:id", authMiddleware(authService, userService), userHandler.DeleteUser)

	api.POST("/photo", authMiddleware(authService, userService), photoHandler.CreatePhoto)
	api.PUT("/photo/:id", authMiddleware(authService, userService), photoHandler.UpdatePhoto)
	api.GET("/photo", photoHandler.GetPhotos)
	api.DELETE("/photo/:id", authMiddleware(authService, userService), photoHandler.Delete)

	router.Run()

	}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1] 
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,  "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims) 

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,  "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64)) 

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,  "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

}
}





	


	

	

	


	