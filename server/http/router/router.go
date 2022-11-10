package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"server/server/http/middleware"
	"server/server/services/wallet/api"
)

func Router(mw *middleware.Middleware, service *api.WalletService) http.Handler {
	root := mux.NewRouter().StrictSlash(true).PathPrefix("/").Subrouter()

	root.Use(mw.RecoverMiddleware)

	service.Router(root)

	return root
}
