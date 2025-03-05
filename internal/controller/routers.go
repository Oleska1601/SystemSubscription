package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) AddAPIRouters(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/subscription-types", s.APIGetSubscriptionTypesHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/set-new-subscription", s.APISetNewSubscriptionHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/activate-subscription", s.APIActivateSubscriptionHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/last-subscription-info", s.APIGetLastSubscriptionInfoHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/news", s.APIGetNewsHandler).Methods(http.MethodGet)
}
