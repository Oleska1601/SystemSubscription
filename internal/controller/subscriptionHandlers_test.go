package controller

import (
	"SystemSubscription/internal/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestAPIGetLastSubscriptionHandlerWithNoErr(t *testing.T) {
	serverWithNoErr := GetTestHTTPServer(NewUseCaseTestWithNoErr())

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedBody   entity.Subscription
	}{
		{"validUserID1", "1", http.StatusOK, entity.Subscription{}},
		{"validUserID2", "2", http.StatusOK, entity.Subscription{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest("GET", "/last-subscription/"+test.userID, nil)
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()
			router.HandleFunc("/last-subscription/{user_id}", serverWithNoErr.APIGetLastSubscriptionHandler)
			router.ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}
			var responseRecorderBody entity.Subscription
			err = json.NewDecoder(responseRecorder.Body).Decode(&responseRecorderBody)
			if err != nil || responseRecorderBody != test.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", responseRecorderBody, test.expectedBody)
			}
		})
	}
}

func TestAPIGetLastSubscriptionHandlerWithErr(t *testing.T) {
	serverWithErr := GetTestHTTPServer(NewUseCaseTestWithErr())

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedBody   string
	}{

		{"invalidUserID3", "3", http.StatusBadRequest, "get last subscription is impossible\n"}, //пользователь отсуствует в системе
		{"invalidUserID4", "abc", http.StatusBadRequest, "get last subscription is impossible\n"},
		{"invalidUserID5", "ab1", http.StatusBadRequest, "get last subscription is impossible\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest("GET", "/last-subscription/"+test.userID, nil)
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()
			router.HandleFunc("/last-subscription/{user_id}", serverWithErr.APIGetLastSubscriptionHandler)
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
