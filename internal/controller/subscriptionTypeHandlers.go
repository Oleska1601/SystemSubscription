package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// 1
// Возвращает список доступных подписок с их уникальными номерами.

// APIGetSubscriptionTypesHandler
// @Summary get subscriptionTypes
// @Description get subscriptionTypes
// @Tags subscriptionTypes
// @Accept       json
// @Produce      json
// @Success 200 {array} entity.SubscriptionType
// @Param token header string true "jwt token for authentification"
// @Failure 401 {string} string "not authorised or invalid token"
// @Failure 403 {string} string "token has been expired"
// @Failure 500 {string} string "error of getting subscription types"
// @Router /api/subscription-types [get]
func (s *Server) APIGetSubscriptionTypesHandler(w http.ResponseWriter, r *http.Request) {
	subscriptionTypes, err := s.u.GetSubscriptionTypes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller-handlersAPI APIGetSubscriptionTypesHandler s.u.GetSubscriptionTypes",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller-handlersAPI APIGetSubscriptionTypesHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(subscriptionTypes)
}
