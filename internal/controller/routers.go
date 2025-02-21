package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) AddAPIRouters(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/subscription-types", s.APIGetSubscriptionTypesHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/last-subscription/{user_id}", s.APIGetLastSubscriptionHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/news/{user_id}", s.APIGetNewsHandler).Methods(http.MethodGet)
}
