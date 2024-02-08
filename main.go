package main

import (
	"fmt"
	"log"
	"os"

	// "test_dealls/app/controllers"
	"test_dealls/handlers"
	authHandler "test_dealls/handlers/auth"
	"test_dealls/middlewares"
	"test_dealls/repositories"
	"test_dealls/services"
	authService "test_dealls/services/auth"

	"test_dealls/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	appName := os.Getenv("APP_NAME")
	modeCors := os.Getenv("MODE_CORS")

	fmt.Println("Value of MY_ENV_VARIABLE:", port)

	app := gin.Default()

	appMode := os.Getenv("APP_MODE")

	if appMode == "local" || appMode == "development" {
		gin.SetMode(gin.DebugMode)
	} else if appMode == "staging" || appMode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	configs.Database()

	app.Use(cors.Default())

	output := fmt.Sprintf(`
	┌───────────────────────────────────────────────────┐ 
	│                    %s                    │ 
	│               http://127.0.0.1:%s               │ 
	│       (bound on host 0.0.0.0 and port %s)       │ 
	│                                                   │ 
	│ Mode ........... %s  Mode CORS .......... %s │
	└───────────────────────────────────────────────────┘ 
	`, appName, port, port, appMode, modeCors)

	fmt.Print(output)
	
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to " + appName)
	})
	
	groupRoute := app.Group("/api")
	{
		groupRoute.GET("/", func(c *gin.Context) {
			c.JSON(200, "Welcome to " + appName)
		})


		// Setup repository
		userRepo := repositories.NewUserRepository(configs.DB)
		swipeRepo := repositories.NewSwipeRepository(configs.DB)

		userService := services.NewUserService(userRepo)
		userHandler := handlers.NewUserHandler(userService)

		registerService := authService.NewRegisterService(userRepo)
		registerHandler := authHandler.NewRegisterHandler(registerService)
		
		loginService := authService.NewLoginService(userRepo)
		loginHandler := authHandler.NewLoginHandler(loginService)
		
		profileService := services.NewProfileService(userRepo, swipeRepo)
		profileHandler := handlers.NewProfileHandler(profileService)
		
		swipeActionService := services.NewSwipeActionService(userRepo, swipeRepo)
		swipeActionHandler := handlers.NewSwipeActionHandler(swipeActionService)
		
		upgradePackageService := services.NewUpgradePackageService(userRepo, swipeRepo)
		upgradePackageHandler := handlers.NewUpgradePackageHandler(upgradePackageService)

		groupRoute.GET("/users", userHandler.GetAllUsers)
		groupRoute.GET("/register", registerHandler.Register)
		groupRoute.POST("/register", registerHandler.Register)
		groupRoute.POST("/login", loginHandler.Login)
		
		groupRouteSecure := groupRoute.Group("/", middlewares.JWTValidate(), middlewares.UserPackageMiddleware())
		groupRouteSecure.GET("/get-profile", profileHandler.Profile)
		groupRouteSecure.POST("/swipe", swipeActionHandler.SwipeAction)
		groupRouteSecure.GET("/upgrade-package", upgradePackageHandler.UpgradePackage)

	}

	app.Run(":" + port)
}
