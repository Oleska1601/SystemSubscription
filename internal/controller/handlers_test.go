package controller

/*
import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRegisterHandlerWithNoErr(t *testing.T) {
	serverWithNoErr := GetTestHTTPServer(NewUseCaseTestWithNoErr())

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "validUser1",
			requestBody:    `{"login": "user_user", "password": "qwerty1"}`,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "validUser2",
			requestBody:    `{"login": "User2", "password": "123"}`,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "validUser3",
			requestBody:    `{"login": "123", "password": "123"}`,
			expectedStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestBody := []byte(test.requestBody)
			request, err := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()
			router.HandleFunc("/register", serverWithNoErr.RegisterHandler)
			router.ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}

		})
	}
}

func TestRegisterHandlerWithErr(t *testing.T) {
	serverWithErr := GetTestHTTPServer(NewUseCaseTestWithErr())

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "invalidUser1",
			requestBody:    `{"login": "user", "password": "qwerty1"}`, //user already exists
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalidUser2",
			requestBody:    ``,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalidUser3",
			requestBody:    `{"login": "", "password": "123"}`,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalidUser4",
			requestBody:    `{"login": "user1", "password": ""}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestBody := []byte(test.requestBody)
			request, err := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()
			router.HandleFunc("/register", serverWithErr.RegisterHandler)
			router.ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}

		})
	}
}

// login
func TestLoginHandlerWithNoErr(t *testing.T) {
	serverWithNoErr := GetTestHTTPServer(NewUseCaseTestWithNoErr())
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "validUser1",
			requestBody:    `{"login": "user", "password": "qwerty"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "validUser2",
			requestBody:    `{"login": "oleska", "password": "1234"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "validUser3",
			requestBody:    `{"login": "oleska2", "password": "1234"}`,
			expectedStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestBody := []byte(test.requestBody)
			request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()
			router.HandleFunc("/login", serverWithNoErr.LoginHandler)
			router.ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}

		})
	}
}

func TestLoginHandlerWithErr(t *testing.T) {
	serverWithErr := GetTestHTTPServer(NewUseCaseTestWithErr())

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "invalidUser1",
			requestBody:    `{"login": "user", "password": "qwerty1"}`, //incorrect password
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalidUser2",
			requestBody:    `{"login": "user1", "password": "qwerty"}`, //incorrect login
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalidUser3",
			requestBody:    `{"login": "user1", "password": "qwerty1"}`, //incorrect login and password
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalidUser4",
			requestBody:    ``,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalidUser5",
			requestBody:    `{"login": "", "password": "123"}`,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalidUser6",
			requestBody:    `{"login": "user1", "password": ""}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestBody := []byte(test.requestBody)
			request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()
			router.HandleFunc("/login", serverWithErr.LoginHandler)
			router.ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}

		})
	}
}
*/
