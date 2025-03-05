package controller

import (
	"SystemSubscription/internal/entity"
	"encoding/json"
	"log/slog"
	"net/http"
)

// 2
// Принимает запрос на выбор подписки Если у пользователя нет активной подписки возвращает токен на оплату Если подписка активна сообщает об этом

// APISetNewSubscriptionHandler
// @Summary post page
// @Description choose subscriptionType and get paymentToken if last subscription is not active
// @Tags subscription
// @Accept json
// @Produce json
// @Param token header string true "jwt token for authentification"
// @Param subscription_type body entity.SubscriptionType true " ID, TypeName, Duration, Price"
// @Success 200 {string} string "paymentToken"
// @Failure 400 {string} string "set new subscription is impossible"
// @Failure 401 {string} string "not authorised or invalid token"
// @Failure 403 {string} string "token has been expired"
// @Failure 500 {string} string "error of setting new subscription"
// @Router /api/set-new-subscription [post]
func (s *Server) APISetNewSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int64)
	if userID == 0 {
		http.Error(w, "set new subscription is impossible", http.StatusBadRequest)
		s.logger.Error("controller APISetNewSubscriptionHandler",
			slog.String("msg", "no user_id in r.Context().Value"), slog.Int("status", http.StatusBadRequest))
		return
	}
	var subscriptionType entity.SubscriptionType
	if err := json.NewDecoder(r.Body).Decode(&subscriptionType); err != nil {
		http.Error(w, "set new subscription is impossible", http.StatusBadRequest)
		s.logger.Error("controller APISetNewSubscriptionHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}

	paymentToken, err := s.u.SetNewSubscription(userID, subscriptionType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller APISetNewSubscriptionHandler s.u.SetNewSubscription",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APISetNewSubscriptionHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(paymentToken)
}

// 3
// Принимает токен на оплату Проверяет его существование и актуальность не истек ли в течение 100 секунд
// Активирует подписку или возвращает ошибку
// возвращаем  подписки если все успешно отработало

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
// @Failure 403 {string} string "token has been expired"
// @Failure 500 {string} string "error of activating subscription"
// @Router /api/activate-subscription [post]
func (s *Server) APIActivateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var paymentToken entity.PaymentToken
	if err := json.NewDecoder(r.Body).Decode(&paymentToken); err != nil {
		http.Error(w, "activate subscription is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIActivateSubscriptionHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	err := s.u.ActivateSubscription(paymentToken.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller APIActivateSubscriptionHandler s.u.ActivateSubscription",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIActivateSubscriptionHandler", slog.Int("status", http.StatusOK))

}

// 4
// Возвращает информацию о текущей активной подписке или об отсутствии активной подписки и информацию о последней

// APIGetLastSubscriptionInfoHandler
// @Summary get page
// @Description get last subscription info
// @Tags subscription
// @Accept json
// @Produce json
// @Param token header string true "jwt token for authentification"
// @Success 200 {object} entity.SubscriptionInfo
// @Failure 400 {string} string "get last subscription info is impossible"
// @Failure 401 {string} string "not authorised or invalid token"
// @Failure 403 {string} string "token has been expired"
// @Failure 500 {string} string "error of getting last subscription info"
// @Router /api/last-subscription-info [get]
func (s *Server) APIGetLastSubscriptionInfoHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int64)
	if userID == 0 {
		http.Error(w, "get last subscription info is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIGetLastSubscriptionInfoHandler",
			slog.String("msg", "no user_id in r.Context().Value"), slog.Int("status", http.StatusBadRequest))
		return
	}
	lastSubscriptionInfo, err := s.u.GetLastSubscriptionInfo(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller APIGetLastSubscriptionInfoHandler s.u.GetLastSubscriptionInfo",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetLastSubscriptionInfoHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(lastSubscriptionInfo)
}
