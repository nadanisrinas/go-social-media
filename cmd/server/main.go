package main

import (
	"fmt"
	"gosocialmedia/controllers"
	"gosocialmedia/models"
	"gosocialmedia/routes"
	"gosocialmedia/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	server   *gin.Engine
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "gosocialmedia"
	db       *gorm.DB
	//comment
	commentService         services.CommentService
	commentRouteController *routes.CommentRouteController
	commentController      *controllers.CommentController
	//user
	userService         services.UserService
	userRouteController *routes.UserRouteController
	userController      *controllers.UserController
	err                 error
	//photo
	photoService         services.PhotoService
	photoRouteController *routes.PhotoRouteController
	photoController      *controllers.PhotoController
	//SocialMedia
	socialMediaService          services.SocialMediaService
	socialMediaRouterController *routes.SocialMediaRouteController
	socialMediaController       *controllers.SocialMediaController
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("err conn to db: ", err)

	}
	errMigrate := db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if errMigrate != nil {
		log.Fatal("error migrate db", errMigrate)
	}
	fmt.Println("db successfully migrated...")
}

func init() {
	StartDB()
	//init service
	commentService = services.NewCommentService(db)
	userService = services.NewUserService(db)
	photoService = services.NewPhotoService(db)
	socialMediaService = services.NewSocialmediaService(db)

	//init controller
	commentController = controllers.NewCommentController(commentService)
	userController = controllers.NewUserController(userService, photoService)
	photoController = controllers.NewPhotoService(photoService)
	socialMediaController = controllers.NewSocialMediaService(socialMediaService)

	//initroutes
	commentRouteController = routes.NewCommentRouteController(commentController)
	userRouteController = routes.NewUserRouteController(userController)
	photoRouteController = routes.NewPhotoRouteController(photoController, userController)
	socialMediaRouterController = routes.NewsocialMediaRouteController(socialMediaController, userController)

}

func GetDB() *gorm.DB {
	return db
}

func main() {

	server = gin.Default()
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "helath check"})
	})
	commentRouteController.CommentRoute(router, commentService)
	userRouteController.UserRoute(router, userService)
	photoRouteController.PhotoRoute(router, photoService, userService)
	socialMediaRouterController.SocialMediaRoute(router, socialMediaService, userService)

	fmt.Println("routes running")
	server.Run(":" + "8080")

}
