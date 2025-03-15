package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

type MiddlewareFunc func(handler http.Handler) http.Handler

func RunHTTPServerOnAddr(addr string, handler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	rootRouter := chi.NewRouter()
	rootRouter.Use(middleware.DefaultLogger)

	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", handler(apiRouter))

	logrus.Info("Starting HTTP server", addr)
	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}

}

func RunHTTPServerOnAddrWithMiddlewares(addr string, handler func(router chi.Router) http.Handler, rootRouter chi.Router, middlewares ...MiddlewareFunc) {
	// config middlewares only for api routes
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", apiRouterWithMiddlewares(handler, middlewares))
	logrus.Info("Starting HTTP server", addr)
	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}

}
func apiRouterWithMiddlewares(handler func(router chi.Router) http.Handler, middlewares []MiddlewareFunc) http.Handler {
	router := chi.NewRouter()

	for _, m := range middlewares {
		router.Use(m)
	}
	return handler(router)
}
