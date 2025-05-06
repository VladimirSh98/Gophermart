package operation

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	operationRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"
	operationMock "github.com/VladimirSh98/Gophermart.git/mocks/operation"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetByUser(t *testing.T) {
	type expect struct {
		status   int
		response any
	}
	type testRequest struct {
		operations []operationRepo.Operation
		err        error
	}
	tests := []struct {
		description string
		expect      expect
		testRequest testRequest
	}{
		{
			description: "Test #1. DB error",
			expect: expect{
				status:   http.StatusInternalServerError,
				response: nil,
			},
			testRequest: testRequest{
				operations: []operationRepo.Operation{},
				err:        errors.New("database error"),
			},
		},
		{
			description: "Test #2. No content",
			expect: expect{
				status:   http.StatusNoContent,
				response: nil,
			},
			testRequest: testRequest{
				operations: []operationRepo.Operation{},
				err:        nil,
			},
		},
		{
			description: "Test #3. Success",
			expect: expect{
				status: http.StatusOK,
				response: []GetByUserResponse{
					{
						Order:       "1",
						Sum:         5.8,
						ProcessedAt: time.Date(2025, time.April, 28, 14, 30, 0, 0, time.UTC),
					},
				},
			},
			testRequest: testRequest{
				operations: []operationRepo.Operation{
					{
						ID:        "1",
						CreatedAt: time.Date(2025, time.April, 28, 14, 30, 0, 0, time.UTC),
						UserID:    1,
						Value:     5.8,
					},
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/api/user/withdrawals", nil)
			w := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserService := operationMock.NewMockServiceInterface(ctrl)
			mockUserService.EXPECT().
				GetByUser(gomock.Any()).
				Return(test.testRequest.operations, test.testRequest.err).AnyTimes()
			customHandler := NewHandler(mockUserService)
			ctx := context.WithValue(request.Context(), authorization.UserIDKey, 1)
			customHandler.GetByUser(w, request.WithContext(ctx))
			var result *http.Response
			result = w.Result()
			defer result.Body.Close()
			var body []byte
			var err error
			body, err = io.ReadAll(result.Body)
			assert.NoError(t, err)
			assert.Equal(t, test.expect.status, result.StatusCode)
			if test.expect.response != nil {
				var realResponse []GetByUserResponse
				err = json.Unmarshal(body, &realResponse)
				assert.NoError(t, err)
				assert.Equal(t, test.expect.response, realResponse)
			}
		})
	}
}
