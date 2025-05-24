package auth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	rewardMock "github.com/VladimirSh98/Gophermart.git/mocks/reward"
	userMock "github.com/VladimirSh98/Gophermart.git/mocks/user"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	type expect struct {
		status      int
		contentType string
		err         error
		checkCookie bool
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
				status:      http.StatusBadRequest,
				contentType: "",
				err:         nil,
				checkCookie: false,
			},
			testRequest: testRequest{
				body: struct{}{},
			},
		},
		{
			description: "Test #2. Uknown database error",
			expect: expect{
				status:      http.StatusInternalServerError,
				contentType: "",
				err:         sql.ErrNoRows,
				checkCookie: false,
			},
			testRequest: testRequest{
				body: RegisterRequest{
					Login:    "testUser",
					Password: "pass",
				},
			},
		},
		{
			description: "Test #3. Existing user",
			expect: expect{
				status:      http.StatusConflict,
				contentType: "",
				err: &pgconn.PgError{
					Code: pgerrcode.UniqueViolation,
				},
				checkCookie: false,
			},
			testRequest: testRequest{
				body: RegisterRequest{
					Login:    "testUser",
					Password: "pass",
				},
			},
		},
		{
			description: "Test #4. Success",
			expect: expect{
				status:      http.StatusOK,
				contentType: "application/json",
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
				http.MethodPost, "/api/user/register", bytes.NewReader(jsonBody),
			)
			w := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserService := userMock.NewMockServiceInterface(ctrl)
			mockUserService.EXPECT().
				Create(gomock.Any(), "testUser", gomock.Any()).
				Return(0, test.expect.err).AnyTimes()
			mockRewardService := rewardMock.NewMockServiceInterface(ctrl)
			mockRewardService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			mockHandler := Handler{User: mockUserService, Reward: mockRewardService}
			mockHandler.Register(w, request)
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
