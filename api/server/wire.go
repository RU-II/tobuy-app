package main

import (
	"github.com/google/wire"

	"tobuy-app/api/server/controllers"
	"tobuy-app/api/server/db"
	"tobuy-app/api/server/repositories"
	"tobuy-app/api/server/router"
	"tobuy-app/api/server/services"
	"tobuy-app/api/server/utils/logic"
)

var superSet = wire.NewSet(
	// DB
	db.Init(),

	// Logic
	logic.NewAuthLogic,
	logic.NewItemsLogic,
	logic.NewResponseLogic,
	logic.NewJWTLogic,
	wire.Bind(new(logic.IAuthLogic), new(*logic.AuthLogic)),
	wire.Bind(new(logic.IItemsLogic), new(*logic.ItemsLogic)),
	wire.Bind(new(logic.IResponseLogic), new(*logic.ResponseLogic)),
	wire.Bind(new(logic.IJWTLogic), new(*logic.JWTLogic)),

	// Repository
	repositories.NewUserRepository,
	repositories.NewItemRepository,
	wire.Bind(new(repositories.IUserRepository), new(*repositories.UserRepository)),
	wire.Bind(new(repositories.IItemRepository), new(*repositories.ItemRepository)),

	// Service
	services.NewAuthService,
	services.NewItemsService,
	wire.Bind(new(services.IAuthService), new(*services.AuthService)),
	wire.Bind(new(services.IItemsService), new(*services.ItemsService)),

	// Controller
	controllers.NewAppController,
	controllers.NewAuthController,
	controllers.NewGroupsController,
	controllers.NewItemsController,
	controllers.NewUsersController,
	wire.Bind(new(controllers.IAppController), new(*controllers.AppController)),
	wire.Bind(new(controllers.IAuthController), new(*controllers.AuthController)),
	wire.Bind(new(controllers.IGroupsController), new(*controllers.GroupsController)),
	wire.Bind(new(controllers.IItemsController), new(*controllers.ItemsController)),
	wire.Bind(new(controllers.IUsersController), new(*controllers.UsersController)),

	// Router
	router.NewAppRouter,
	router.NewAuthRouter,
	router.NewGroupsRouter,
	router.NewItemsRouter,
	router.NewUsersRouter,
	wire.Bind(new(router.IAppRouter), new(*router.AppRouter)),
	wire.Bind(new(router.IAuthRouter), new(*router.AuthRouter)),
	wire.Bind(new(router.IGroupsRouter), new(*router.GroupsRouter)),
	wire.Bind(new(router.IItemsRouter), new(*router.ItemsRouter)),
	wire.Bind(new(router.IUsersRouter), new(*router.UsersRouter)),

	router.NewMainRouter,
	wire.Bind(new(router.IMainRouter), new(*router.MainRouter)),
)

func Initialize() *router.MainRouter {
	wire.Build(superSet)

	return &router.MainRouter{}
}

// func Initialize() router.MainRouter {

// 	wire.Build(
// 		logic.NewAuthLogic,
// 		logic.NewItemsLogic,
// 		logic.NewResponseLogic,
// 		logic.NewJWTLogic,
// 		repositories.NewUserRepository, repositories.NewItemRepository,services.NewAuthService,
// 		services.NewItemsService,
// 		controllers.NewAppController, controllers.NewAuthController, controllers.NewGroupsController, controllers.NewItemsController, controllers.NewUsersController,
// 		router.NewAppRouter,
// 		router.NewAuthRouter,
// 		router.NewGroupsRouter,
// 		router.NewItemsRouter,
// 		router.NewUsersRouter,
// 	)

// 	return nil
// }
