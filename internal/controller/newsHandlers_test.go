package controller

/*
func TestAPIGetNewsHandlerWithNoErr(t *testing.T) {
	serverWithNoErr := GetTestHTTPServer(NewUseCaseTestWithNoErr())

	tests := []struct {
		name           string
		expectedStatus int
		expectedBody   entity.News
	}{
		{"validUserID1", http.StatusOK, entity.News{}},
		{"validUserID2", http.StatusOK, entity.News{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			request, err := http.NewRequest("GET", "/news", nil)
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()

			router.HandleFunc("/news", serverWithNoErr.APIGetNewsHandler)
			router.ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}
			var responseRecorderBody entity.News
			json.NewDecoder(responseRecorder.Body).Decode(&responseRecorderBody)
			if responseRecorderBody.Message != test.expectedBody.Message {
				t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body, test.expectedBody)
			}
		})
	}
}

func TestAPIGetNewsHandlerWithErr(t *testing.T) {
	serverWithErr := GetTestHTTPServer(NewUseCaseTestWithErr())

	tests := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{"invalidUserID1", http.StatusBadRequest, "get news is impossible\n"},
		{"invalidUserID2", http.StatusBadRequest, "get news is impossible\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			request, err := http.NewRequest("GET", "/news", nil)
			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
			router := mux.NewRouter()

			router.HandleFunc("/news", serverWithErr.APIGetNewsHandler)
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
