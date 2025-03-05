package controller

/*
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
*/
/*
func TestAPIChooseSubscriptionHandlerWithNoErr(t *testing.T) {
	serverWithNoErr := GetTestHTTPServer(NewUseCaseTestWithNoErr())

	tests := []struct {
		name           string
		userID         string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "validUserID1",
			userID:         "3",
			requestBody:    `{"subscription_type_id": "1", "type_name": "1 second", type_name": "1", price": "10"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   "", //как получить expected body если paymentToken генерируется рандомно???
		},
		{
			name:           "validUserID2",
			userID:         "4",
			requestBody:    `{"subscription_type_id": "1", "type_name": "1 second", type_name": "1", price": "10"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest("POST", "/choose-subscription/"+test.userID, nil)
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()
			router.HandleFunc("/choose-subscription/{user_id}", serverWithNoErr.APIChooseSubscriptionHandler)
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


func TestAPIChooseSubscriptionHandlerWithErr(t *testing.T) {
	serverWithErr := GetTestHTTPServer(NewUseCaseTestWithErr())

	tests := []struct {
		name           string
		userID         string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "invalidUserID1",
			userID:         "3",
			requestBody:    `{"subscription_type_id": "1", "type_name": "1 second", type_name": "1", price": "10"}`,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalidUserID2",
			userID:         "4",
			requestBody:    `{"subscription_type_id": "1", "type_name": "1 second", type_name": "1", price": "10"}`,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalidUserID3",
			userID:         "4",
			requestBody:    ``,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalidUserID4",
			userID:         "4",
			requestBody:    `{"subscription_type_id": "1", type_name": "1", price": "10"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest("POST", "/choose-subscription/"+test.userID, nil)
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()
			router.HandleFunc("/choose-subscription/{user_id}", serverWithErr.APIChooseSubscriptionHandler)
			router.ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}

		})
	}
}
*/
