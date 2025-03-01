package controller

import (
	"SystemSubscription/internal/entity"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 5
// Проверяет наличие активной подписки у пользователя. Если подписка активна, возвращает текст «новость». Если нет, сообщает о её отсутствии.

// APIGetNewsHandler
// @Summary get page
// @Description get news if last subscription is active
// @Tags news
// @Accept json
// @Produce json
// @Param token header string true "jwt token for authentification"
// @Param user_id path int true "UserID"
// @Success 200 {object} entity.News
// @Failure 400 {string} string "get news is impossible"
// @Failure 401 {string} string "not authorised or invalid token"
// @Failure 403 {string} string "token has been expired or current subscription is not active"
// @Failure 500 {string} string "error of getting last subscription"
// @Router /api/news/{user_id} [get]
func (s *Server) APIGetNewsHandler(w http.ResponseWriter, r *http.Request) {
	userIDString := mux.Vars(r)["user_id"]
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		http.Error(w, "get news is impossible", http.StatusBadRequest)
		s.logger.Error("controller-handlersAPI APIGetNewsHandler strconv.Atoi", slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	lastSubscription, err := s.u.GetLastSubscription(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller-handlersAPI APIGetNewsHandler s.u.GetLastSubscription", slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	if !s.u.IsSubscriptionStatusActive(lastSubscription) {
		http.Error(w, "current subscription is not active", http.StatusForbidden)
		s.logger.Error("controller-handlersAPI APIGetNewsHandler", slog.Int("status", http.StatusForbidden))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller-handlersAPI APIGetNewsHandler", slog.Int("status", http.StatusOK))
	news := entity.News{Message: "news"}
	json.NewEncoder(w).Encode(news)
}
