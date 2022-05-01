package main

import (
	"fmt"
	"time"
	"tobuy-app/api/controllers"
	"tobuy-app/api/db"
	"tobuy-app/api/repositories"
	"tobuy-app/api/router"
	"tobuy-app/api/services"
	"tobuy-app/api/utils"
	"tobuy-app/api/utils/logic"

	"github.com/rs/zerolog/log"
)

// @title           Swagger ToBuyApp API
// @version         1.0
// @description     This is the ToBuyApp server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

// @securitydefinitions.oauth2.application  OAuth2Application
// @tokenUrl                                https://example.com/oauth/token
// @scope.write                             Grants write access
// @scope.admin                             Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit  OAuth2Implicit
// @authorizationurl                     https://example.com/oauth/authorize
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.password  OAuth2Password
// @tokenUrl                             https://example.com/oauth/token
// @scope.read                           Grants read access
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode  OAuth2AccessCode
// @tokenUrl                               https://example.com/oauth/token
// @authorizationurl                       https://example.com/oauth/authorize
// @scope.admin                            Grants read and write access to administrative information
func main() {
	// log設定
	log.Logger = *utils.CreateLogger(fmt.Sprintf("./logs/access.%s.log", time.Now().Local().Format("20060102")))

	// DB接続
	db := db.Init()

	// logic層
	// authLogic :=  logic.NewAuthLogic()
	// userLogic :=  logic.NewUserLogic()
	itemsLogic := logic.NewItemsLogic()
	responseLogic := logic.NewResponseLogic()
	// jwtLogic := logic.NewJWTLogic()

	// repository層
	// userRepo := repositories.NewUserRepository(db)
	itemsRepo := repositories.NewItemsRepository(db)

	// service層
	// authService := services.NewAuthService(userRepo, authLogic, userLogic, responseLogic, jwtLogic, authValidate)
	itemsService := services.NewItemsService(itemsRepo, itemsLogic, responseLogic)

	// controller層
	appController := controllers.NewAppController()
	authController := controllers.NewAuthController()
	groupsController := controllers.NewGroupsController()
	itemsController := controllers.NewItemsController(itemsService)
	usersController := controllers.NewUsersController()

	// router設定
	appRouter := router.NewAppRouter(appController)
	authRouter := router.NewAuthRouter(authController)
	groupsRouter := router.NewGroupsRouter(groupsController)
	itemsRouter := router.NewItemsRouter(itemsController)
	usersRouter := router.NewUsersRouter(usersController)

	mainRouter := router.NewMainRouter(appRouter, authRouter, groupsRouter, itemsRouter, usersRouter)

	// API起動
	log.Fatal().Err(mainRouter.StartWebServer()).Msg("Startup failed")
}
