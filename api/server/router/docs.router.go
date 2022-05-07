package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type IDocsRouter interface {
	SetDocsRouting(router *mux.Router)
}

type DocsRouter struct {
}

func NewDocsRouter() *DocsRouter {
	return &DocsRouter{}
}

func (dr *DocsRouter) SetDocsRouting(router *mux.Router) {
	router.PathPrefix(basePath + "/docs").Handler(httpSwagger.WrapHandler)
}
