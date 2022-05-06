package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "tobuy-app/api/server/docs"
)

const (
	basePath = "/api/v1"
)

type MainRouter interface {
	setupRouting() *mux.Router
	StartWebServer() error
}

type mainRouter struct {
	appR    AppRouter
	authR   AuthRouter
	groupsR GroupsRouter
	itemsR  ItemsRouter
	usersR  UsersRouter
}

func NewMainRouter(appR AppRouter, authR AuthRouter, groupsR GroupsRouter, itemsR ItemsRouter, usersR UsersRouter) MainRouter {
	return &mainRouter{appR, authR, groupsR, itemsR, usersR}
}

func (mainRouter *mainRouter) setupRouting() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	mainRouter.appR.SetAppRouting(router)
	mainRouter.authR.SetAuthRouting(router)
	mainRouter.groupsR.SetGroupsRouting(router)
	mainRouter.itemsR.SetItemsRouting(router)
	mainRouter.usersR.SetUsersRouting(router)

	// Swagger設定
	router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("full"),
		httpSwagger.DomID("#swagger-ui"),
	)).Methods(http.MethodGet)
	log.Info().Msg("Start Swagger API")

	return router
}

func (mainRouter *mainRouter) StartWebServer() error {
	log.Info().Msg("Start ToBuyApp server")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), mainRouter.setupRouting())
}
