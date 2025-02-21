package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 1
// Возвращает список доступных подписок с их уникальными номерами.
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

// 2
// Принимает запрос на выбор подписки. Если у пользователя нет активной подписки, возвращает токен на оплату. Если подписка активна, сообщает об этом.
func (s *Server) APIInsertSubscriptionHandler(w http.ResponseWriter, r *http.Request) {

}

// 4
// Возвращает информацию о текущей активной подписке или об отсутствии активной подписки и информацию о последней
func (s *Server) APIGetLastSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	// извлекаем id пользователя из пути
	userIDString := mux.Vars(r)["user_id"]
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		http.Error(w, "get last subscription is impossible", http.StatusBadRequest)
		s.logger.Error("controller-handlersAPI GetLastSubscriptionHandler strconv.Atoi", slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	//если у пользователя вообще не было подписок, то будет internal server error, тк sql.QueryNoRows
	//или этот случай не ошибочный и его как-то отдельно обработать-???
	lastSubscription, err := s.u.GetLastSubscription(int64(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller-handlersAPI GetLastSubscriptionHandler s.u.GetLastSubscription", slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller-handlersAPI GetLastSubscriptionHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(lastSubscription)
}

// 5
// Проверяет наличие активной подписки у пользователя. Если подписка активна, возвращает текст «новость». Если нет, сообщает о её отсутствии.
func (s *Server) APIGetNewsHandler(w http.ResponseWriter, r *http.Request) {
	userIDString := mux.Vars(r)["user_id"]
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		http.Error(w, "get news is impossible", http.StatusBadRequest)
		s.logger.Error("controller-handlersAPI APIGetNewsHandler strconv.Atoi", slog.Any("error", err), slog.Int("status", http.StatusOK))
		return
	}
	err = s.u.GetNews(int64(userID))
	if err != nil {
		if err.Error() == "internal server error" {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.logger.Error("controller-handlersAPI APIGetNewsHandler s.u.GetNews", slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
			return
		}
		http.Error(w, err.Error(), http.StatusForbidden)
		s.logger.Error("controller-handlersAPI APIGetNewsHandler s.u.GetNews", slog.Any("error", err), slog.Int("status", http.StatusForbidden))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller-handlersAPI APIGetNewsHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode("news")
}
