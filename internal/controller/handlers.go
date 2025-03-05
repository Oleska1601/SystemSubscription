package controller

import (
	"SystemSubscription/internal/entity"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home"))
}

func (s *Server) InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("info"))
}

// RegisterHandler
// @Summary register page
// @Description register by login and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body entity.Credentials true "Login, Password"
// @Success 201 {string} string "register successful"
// @Failure 400 {string} string "register is impossible"
// @Failure 500 {string} string "error with register or incorrect login or password"
// @Router /register [post]
func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "register is impossible", http.StatusBadRequest)
		s.logger.Error("controller RegisterHandler json.NewDecoder(r.Body).Decode", slog.Any("error", err),
			slog.Int("status", http.StatusBadRequest))
		return
	}
	if err := s.u.RegisterUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller RegisterHandler s.u.RegisterUser", slog.Any("error", err),
			slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusCreated)
	s.logger.Info("controller RegisterHandler", slog.String("user", user.Login), slog.Int("status", http.StatusCreated))
}

// LoginHandler
// @Summary login page
// @Description login by login and password
// @Tags users
// @Param user body entity.Credentials true "Login, Password"
// @Success 200 {string} string "login successful"
// @Failure 400 {string} string "login is impossible"
// @Failure 500 {string} string "error with login or incorrect login or password or login is impossible"
// @Router /login [post]
func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "login is impossible", http.StatusBadRequest)
		s.logger.Error("controller LoginHandler json.NewDecoder(r.Body).Decode", slog.Any("error", err),
			slog.Int("status", http.StatusBadRequest))
		return
	}
	loginUser, err := s.u.LoginUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller LoginHandler s.u.LoginUser", slog.Any("error", err),
			slog.Int("status", http.StatusInternalServerError))
		return
	}
	token, err := s.GenerateToken(loginUser.ID, "simple token for user"+strconv.Itoa(int(loginUser.ID)))
	if err != nil {
		http.Error(w, "login is impossible", http.StatusInternalServerError)
		s.logger.Error("controller LoginHandler s.GenerateToken", slog.Any("error", err),
			slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.Header().Set("token", token)
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller LoginHandler", slog.String("user", user.Login), slog.Int("status", http.StatusOK))
}
