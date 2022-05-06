package main

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"tobuy-app/api/server/controllers"
	"tobuy-app/api/server/db"
	"tobuy-app/api/server/repositories"
	"tobuy-app/api/server/router"
	"tobuy-app/api/server/services"
	"tobuy-app/api/server/utils"
	"tobuy-app/api/server/utils/logic"
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

	// 環境変数設定
	if err := godotenv.Load(".env"); err != nil {
		log.Error().Err(err).Msg("環境変数を読み込めませんでした。")
	}

	// DB接続
	db := db.Init()

	// logic層
	authLogic := logic.NewAuthLogic()
	itemsLogic := logic.NewItemsLogic()
	responseLogic := logic.NewResponseLogic()
	jwtLogic := logic.NewJWTLogic()

	// repository層
	userRepo := repositories.NewUserRepository(db)
	itemsRepo := repositories.NewItemRepository(db)

	// service層
	authService := services.NewAuthService(userRepo, authLogic, responseLogic, jwtLogic)
	itemsService := services.NewItemsService(itemsRepo, itemsLogic, responseLogic)

	// controller層
	appController := controllers.NewAppController()
	authController := controllers.NewAuthController(authService)
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
	log.Fatal().Err(mainRouter.StartWebServer()).Msg("Fail to start ToBuyApp server")
}
