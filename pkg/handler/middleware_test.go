package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo_app/pkg/service"
	"todo_app/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_userIdentity(t *testing.T){
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	testTable := []struct{
		name string
		headerName string
		headreValue string
		token string
		mockbehavior mockBehavior
		expectedStatusCode int
		expectedResponseBody string
	}{
		{
			name: "OK",
			headerName: "Authorization",
			headreValue: "Bearer token",
			token: "token",
			mockbehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: "1",
		},
		{
			name: "Empty header",
			mockbehavior: func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name: "Invalid struct of header",
			headerName: "Authorization",
			headreValue: "token",
			mockbehavior: func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name: "Invalid Bearer",
			headerName: "Authorization",
			headreValue: "Bearr token",
			mockbehavior: func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponseBody: `{"message":"invalid bearer header"}`,
		},
	}

	for _, testCase := range testTable{
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockbehavior(auth, testCase.token)

			service := &service.Service{Authorization: auth}
			handler := NewHandler(service)

			r := gin.New()
			r.GET("/protected", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(http.StatusOK, fmt.Sprintf("%d", id.(int)))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headreValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}