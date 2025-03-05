package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// 5
// Проверяет наличие активной подписки у пользователя.
// Если подписка активна, возвращает текст «новость». Если нет, сообщает о её отсутствии.

// APIGetNewsHandler
// @Summary get page
// @Description get news if last subscription is active
// @Tags news
// @Accept json
// @Produce json
// @Param token header string true "jwt token for authentification"
// @Success 200 {object} entity.News
// @Failure 400 {string} string "get news is impossible"
// @Failure 401 {string} string "not authorised or invalid token"
// @Failure 403 {string} string "token has been expired or current subscription is not active"
// @Failure 500 {string} string "error of getting news"
// @Router /api/news [get]
func (s *Server) APIGetNewsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int64)
	if userID == 0 {
		http.Error(w, "get news is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIGetNewsHandler",
			slog.String("msg", "no user_id in r.Context().Value"), slog.Int("status", http.StatusBadRequest))
		return
	}
	news, err := s.u.GetNews(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller APIGetNewsHandler s.u.GetNews",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetNewsHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(news)
}
