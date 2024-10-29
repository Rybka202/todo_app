package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"
	"todo_app"
	"todo_app/pkg/service"
	mock_service "todo_app/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_signUp(t *testing.T){
	type mockBehavior func(s *mock_service.MockAuthorization, user todo.User)

	testTable := []struct{
		name string
		inputBody string
		inputUser todo.User
		mockbehavior mockBehavior
		expectedStatusCode int
		expectedRequestBody string 
	}{
		{
			name: "OK",
			inputBody: `{"name":"Test","username":"test","password":"qwerty"}`,
			inputUser: todo.User{
				Name: "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockbehavior: func(s *mock_service.MockAuthorization, user todo.User) {
				s.EXPECT().CreateUser(user).Return(1,nil)
			},
			expectedStatusCode: 201,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Empty input data",
			inputBody: `{"username":"test","password":"qwerty"}`,
			mockbehavior: func(s *mock_service.MockAuthorization, user todo.User) {},
			expectedStatusCode: 400,
			expectedRequestBody: `{"message":"error invalid input data"}`,
		},
		{
			name: "Service Failure",
			inputBody: `{"name":"Test","username":"test","password":"qwerty"}`,
			inputUser: todo.User{
				Name: "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockbehavior: func(s *mock_service.MockAuthorization, user todo.User) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("service failure"))
			},
			expectedStatusCode: 500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable{
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockbehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
