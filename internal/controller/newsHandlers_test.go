package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestAPIGetNewsHandlerWithNoErr(t *testing.T) {
	serverWithNoErr := GetTestHTTPServer(NewUseCaseTestWithNoErr())

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedBody   string
	}{
		{"validUserID1", "1", http.StatusOK, "\"news\"\n"},
		{"validUserID2", "2", http.StatusOK, "\"news\"\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			request, err := http.NewRequest("GET", "/news/"+test.userID, nil)
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()

			router.HandleFunc("/news/{user_id}", serverWithNoErr.APIGetNewsHandler)
			router.ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}

			if responseRecorder.Body.String() != test.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), test.expectedBody)
			}
		})
	}
}

func TestAPIGetNewsHandlerWithErr(t *testing.T) {
	serverWithErr := GetTestHTTPServer(NewUseCaseTestWithErr())

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedBody   string
	}{
		{"invalidUserID3", "3", http.StatusBadRequest, "get news is impossible\n"}, //пользователь отсуствует в системе
		{"invalidUserID4", "abc", http.StatusBadRequest, "get news is impossible\n"},
		{"invalidUserID5", "ab1", http.StatusBadRequest, "get news is impossible\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			request, err := http.NewRequest("GET", "/news/"+test.userID, nil)
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()

			router.HandleFunc("/news/{user_id}", serverWithErr.APIGetNewsHandler)
			router.ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}

			if responseRecorder.Body.String() != test.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), test.expectedBody)
			}
		})
	}
}
