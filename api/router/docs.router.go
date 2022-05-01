package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type DocsRouter interface {
	SetDocsRouting(router *mux.Router)
}

type docsRouter struct {
}

func NewDocsRouter() DocsRouter {
	return &docsRouter{}
}

func (dor *docsRouter) SetDocsRouting(router *mux.Router) {
	router.PathPrefix(basePath + "/docs").Handler(httpSwagger.WrapHandler)
}
