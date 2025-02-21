package controller

import (
	"SystemSubscription/internal/usecase"
	"SystemSubscription/pkg/logger"
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type Server struct {
	router    *mux.Router
	u         *usecase.Usecase
	logger    *logger.Logger
	secretKey []byte
}

func New(u *usecase.Usecase, l *logger.Logger) *Server {
	s := &Server{
		router: mux.NewRouter(),
		u:      u,
		logger: l,
	}
	s.router.HandleFunc("/home", s.HomeHandler)
	s.router.HandleFunc("/info", s.InfoHandler)
	s.router.HandleFunc("/login", s.LoginHandler).Methods("POST")
	s.router.HandleFunc("/register", s.RegisterHandler).Methods("POST")
	apiRouter := s.router.PathPrefix("/api").Subrouter()
	apiRouter.Use(s.checkToken)
	s.AddAPIRouters(apiRouter)
	//swagger
	return s
}

func (s *Server) Run(port string) {
	s.logger.Info("http://127.0.0.1:" + port)
	if err := http.ListenAndServe("localhost:"+port, s.router); err != nil {
		s.logger.Error("fatal error", slog.Any("error", err))
		return
	}
}

func (s *Server) checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("token")
		if tokenString == "" {
			http.Error(w, "not authorised", http.StatusUnauthorized)
			s.logger.Error("controller-httpserver checkToken", slog.String("msg", "not authorised"), slog.Int("status", http.StatusUnauthorized))
			return
		}
		claims := Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) { return s.secretKey, nil })
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			s.logger.Error("controller-httpserver checkToken jwt.ParseWithClaims", slog.Any("error", err), slog.Int("status", http.StatusUnauthorized))
			return
		}
		now := time.Now().Unix()
		if now > claims.ExpiresAt {
			http.Error(w, "token has been expired", http.StatusForbidden)
			s.logger.Error("controller-httpserver checkToken", slog.String("msg", "token has been expired"), slog.Int("status", http.StatusForbidden))
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(context.Background(), "user_id", claims.userID)))
	})
}
