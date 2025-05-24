package auth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/VladimirSh98/Gophermart.git/internal/app/repository/user"
	userMock "github.com/VladimirSh98/Gophermart.git/mocks/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	type expect struct {
		status       int
		contentType  string
		expectedUser user.User
		err          error
		checkCookie  bool
	}
	type testRequest struct {
		username string
		body     any
	}
	tests := []struct {
		description string
		expect      expect
		testRequest testRequest
	}{
		{
			description: "Test #1. Wrong body",
			expect: expect{
				status:       http.StatusBadRequest,
				contentType:  "",
				expectedUser: user.User{},
				err:          nil,
				checkCookie:  false,
			},
			testRequest: testRequest{
				body: struct{}{},
			},
		},
		{
			description: "Test #2. No user in database",
			expect: expect{
				status:       http.StatusUnauthorized,
				contentType:  "",
				expectedUser: user.User{},
				err:          sql.ErrNoRows,
				checkCookie:  false,
			},
			testRequest: testRequest{
				body: RegisterRequest{
					Login:    "testUser",
					Password: "pass",
				},
			},
		},
		{
			description: "Test #3. Wrong password",
			expect: expect{
				status:      http.StatusUnauthorized,
				contentType: "",
				expectedUser: user.User{
					ID:    1,
					Login: "testUser",
					Hash:  "$2a$10$R7tYatRL4Egf0uY9xCX2.OCIaX34QYbEQEtl4499C/7uGqJIDoCIW",
				},
				err:         nil,
				checkCookie: false,
			},
			testRequest: testRequest{
				body: RegisterRequest{
					Login:    "testUser",
					Password: "wrongPass",
				},
			},
		},
		{
			description: "Test #4. Success",
			expect: expect{
				status:      http.StatusOK,
				contentType: "application/json",
				expectedUser: user.User{
					ID:    1,
					Login: "testUser",
					Hash:  "$2a$10$R7tYatRL4Egf0uY9xCX2.OCIaX34QYbEQEtl4499C/7uGqJIDoCIW",
				},
				err:         nil,
				checkCookie: true,
			},
			testRequest: testRequest{
				body: RegisterRequest{
					Login:    "testUser",
					Password: "pass",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			jsonBody, _ := json.Marshal(test.testRequest.body)
			request := httptest.NewRequest(
				http.MethodPost, "/api/user/login", bytes.NewReader(jsonBody),
			)
			w := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := userMock.NewMockServiceInterface(ctrl)
			mockService.EXPECT().
				GetByLogin(gomock.Any(), "testUser", false).
				Return(test.expect.expectedUser, test.expect.err).AnyTimes()
			mockHandler := Handler{User: mockService}
			mockHandler.Login(w, request)
			result := w.Result()
			assert.Equal(t, test.expect.status, result.StatusCode, "Неверный код ответа")
			defer result.Body.Close()
			assert.Equal(t, test.expect.contentType, result.Header.Get("Content-Type"), "Неверный тип контента в хедере")
			if test.expect.checkCookie {
				cookies := result.Cookies()
				authCookie := false
				for _, cookie := range cookies {
					if cookie.Name == "Authorization" {
						authCookie = true
						break
					}
				}
				if !authCookie {
					t.Errorf("Отсутствует печенька авторизации")
				}
			}
		})
	}
}
