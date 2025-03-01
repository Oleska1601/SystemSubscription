package controller

import (
	"SystemSubscription/internal/entity"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 2
// Принимает запрос на выбор подписки. Если у пользователя нет активной подписки, возвращает токен на оплату. Если подписка активна, сообщает об этом.

// APIChooseSubscriptionHandler
// @Summary post page
// @Description choose subscriptionType and get paymentToken if last subscription is not active
// @Tags subscription
// @Accept json
// @Produce json
// @Param token header string true "jwt token for authentification"
// @Param user_id path int true "UserID"
// @Param subscription_type body entity.SubscriptionType true " ID, TypeName, Duration, Price"
// @Success 200 {string} string "paymentToken"
// @Failure 400 {string} string "choose subscription is impossible"
// @Failure 401 {string} string "not authorised or invalid token"
// @Failure 403 {string} string "token has been expired or current subscription is active"
// @Failure 500 {string} string "error of getting last subscription or error of adding payment"
// @Router /api/choose-subscription/{user_id} [post]
func (s *Server) APIChooseSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	userIDString := mux.Vars(r)["user_id"]
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		http.Error(w, "choose subscription is impossible", http.StatusBadRequest)
		s.logger.Error("controller-handlersAPI APIChooseSubscriptionHandler strconv.Atoi", slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	lastSubscription, err := s.u.GetLastSubscription(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller-handlersAPI APIChooseSubscriptionHandler s.u.GetLastSubscription", slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	if s.u.IsSubscriptionStatusActive(lastSubscription) {
		http.Error(w, "current subscription is active", http.StatusForbidden)
		s.logger.Error("controller-handlersAPI APIChooseSubscriptionHandler", slog.Int("status", http.StatusForbidden))
		return
	}

	//запрашиваем подписку: id, name, dur, price
	var subscriptionType entity.SubscriptionType
	if err := json.NewDecoder(r.Body).Decode(&subscriptionType); err != nil {
		http.Error(w, "choose subscription is impossible", http.StatusBadRequest)
		s.logger.Error("controller - handlersAPI APIChooseSubscriptionHandler json.NewDecoder(r.Body).Decode", slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	paymentToken, err := s.u.AddPayment(&subscriptionType, userID) //тк нет никаких скидок, то внутри subscriptionType.Price будет передаваться напрямую
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller - handlersAPI APIChooseSubscriptionHandler s.u.AddPayment", slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller - handlersAPI APIChooseSubscriptionHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(paymentToken)
}

// 3
// Принимает токен на оплату. Проверяет его существование и актуальность (не истек ли timeout в 10 секунд). Активирует подписку или возвращает ошибку
// возвращаем id подписки если все успешно отработало

// APIActivateSubscriptionHandler
// @Summary post page
// @Description add subscription if paymentToken has not been expired
// @Tags subscription
// @Accept json
// @Produce json
// @Param token header string true "jwt token for authentification"
// @Param paymentToken body entity.PaymentToken true "paymentToken"
// @Success 200 {int} int "subscriptionID"
// @Failure 400 {string} string "activate subscription is impossible"
// @Failure 401 {string} string "not authorised or invalid token"
// @Failure 403 {string} string "token has been expired or payment token has been expired"
// @Failure 500 {string} string "error of getting payment or error of updating payment or error of adding subscription"
// @Router /api/activate-subscription [post]
func (s *Server) APIActivateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var paymentToken entity.PaymentToken
	if err := json.NewDecoder(r.Body).Decode(&paymentToken); err != nil {
		http.Error(w, "activate subscription is impossible", http.StatusBadRequest)
		s.logger.Error("controller - handlersAPI APIActivateSubscriptionHandler json.NewDecoder(r.Body).Decode", slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	payment, err := s.u.GetPayment(paymentToken.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller - handlersAPI APIActivateSubscriptionHandler s.u.GetPayment", slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	if !s.u.IsPaymentActive(payment) {
		http.Error(w, "payment token has been expired", http.StatusForbidden)
		s.logger.Error("controller - handlersAPI APIActivateSubscriptionHandler", slog.Int("status", http.StatusForbidden))
		return
	}
	//change status: "token is generated" -> "paid"
	newPaymentStatus := "paid"
	err = s.u.UpdatePayment(payment.ID, newPaymentStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller - handlersAPI APIActivateSubscriptionHandler s.u.UpdatePayment", slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	subscriptionID, err := s.u.AddSubscription(payment.SubscriptionTypeName, payment.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller - handlersAPI APIActivateSubscriptionHandler s.u.AddSubscription", slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller - handlersAPI APIActivateSubscriptionHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(subscriptionID)

}

// 4
// Возвращает информацию о текущей активной подписке или об отсутствии активной подписки и информацию о последней

// APIGetLastSubscriptionHandler
// @Summary get page
// @Description get last subscription
// @Tags subscription
// @Accept json
// @Produce json
// @Param token header string true "jwt token for authentification"
// @Param user_id path int true "UserID"
// @Success 200 {object} entity.Subscription
// @Failure 400 {string} string "get last subscription is impossible"
// @Failure 401 {string} string "not authorised or invalid token"
// @Failure 403 {string} string "token has been expired"
// @Failure 500 {string} string "error of getting last subscription"
// @Router /api/last-subscription/{user_id} [get]
func (s *Server) APIGetLastSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	// извлекаем id пользователя из пути
	userIDString := mux.Vars(r)["user_id"]
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		http.Error(w, "get last subscription is impossible", http.StatusBadRequest)
		s.logger.Error("controller-handlersAPI APIGetLastSubscriptionHandler strconv.Atoi", slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	lastSubscription, err := s.u.GetLastSubscription(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller-handlersAPI APIGetLastSubscriptionHandler s.u.GetLastSubscription", slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	//если у пользователя вообще не было подписок -> какая ошибка должна быть 500 или 400 или еще какая-то
	if !s.u.IsLastSubscription(lastSubscription) {
		http.Error(w, "get last subscription is impossible", http.StatusBadRequest)
		s.logger.Error("controller-handlersAPI APIGetLastSubscriptionHandler s.u.IsLastSubscription", slog.Any("msg", "last subscription does not exist"), slog.Int("status", http.StatusBadRequest))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller-handlersAPI APIGetLastSubscriptionHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(lastSubscription)
}
