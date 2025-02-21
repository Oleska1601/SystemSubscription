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

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "register is impossible", http.StatusBadRequest)
		s.logger.Error("controller-handlers RegisterHandler json.NewDecoder(r.Body).Decode", slog.Any("error", err),
			slog.Int("status", http.StatusBadRequest))
		return
	}
	if err := s.u.RegisterUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller-handlers RegisterHandler s.u.RegisterUser", slog.Any("error", err),
			slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusCreated)
	s.logger.Info("controller-handlers RegisterHandler", slog.String("msg", "register successful"), slog.String("user", user.Login), slog.Int("status", http.StatusCreated))
	json.NewEncoder(w).Encode(user.Secret)
}

// login + password
func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "login is impossible", http.StatusBadRequest)
		s.logger.Error("controller-handlers LoginHandler json.NewDecoder(r.Body).Decode", slog.Any("error", err),
			slog.Int("status", http.StatusBadRequest))
		return
	}
	loginUser, err := s.u.LoginUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Error("controller-handlers LoginHandler s.u.LoginUser", slog.Any("error", err),
			slog.Int("status", http.StatusInternalServerError))
		return
	}
	token, err := s.GenerateToken(loginUser.ID, "simple token for user"+strconv.Itoa(int(loginUser.ID)))
	if err != nil {
		http.Error(w, "login is impossible", http.StatusInternalServerError)
		s.logger.Error("controller-handlers LoginHandler s.GenerateToken", slog.Any("error", err),
			slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.Header().Set("token", token)
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller-handlers LoginHandler", slog.String("msg", "login successful"), slog.String("user", user.Login), slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(loginUser.Secret)
}
