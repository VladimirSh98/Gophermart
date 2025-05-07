package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
	accrualMock "github.com/VladimirSh98/Gophermart.git/mocks/accrual"
	orderMock "github.com/VladimirSh98/Gophermart.git/mocks/order"
	rewardMock "github.com/VladimirSh98/Gophermart.git/mocks/reward"
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
		orders []orderRepo.Order
		err    error
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
				orders: []orderRepo.Order{},
				err:    errors.New("database error"),
			},
		},
		{
			description: "Test #2. No content",
			expect: expect{
				status:   http.StatusNoContent,
				response: nil,
			},
			testRequest: testRequest{
				orders: []orderRepo.Order{},
				err:    nil,
			},
		},
		{
			description: "Test #3. Success",
			expect: expect{
				status: http.StatusOK,
				response: []GetByUserResponse{
					{
						Number:     "1",
						Accrual:    5.8,
						UploadedAt: time.Date(2025, time.April, 28, 14, 30, 0, 0, time.UTC),
						Status:     "PROCESSED",
					},
					{
						Number:     "1",
						UploadedAt: time.Date(2025, time.April, 27, 14, 30, 0, 0, time.UTC),
						Status:     "INVALID",
					},
				},
			},
			testRequest: testRequest{
				orders: []orderRepo.Order{
					{
						ID:         "1",
						UploadedAt: time.Date(2025, time.April, 28, 14, 30, 0, 0, time.UTC),
						UserID:     1,
						Value: sql.NullFloat64{
							Float64: 5.8,
							Valid:   true,
						},
						Status: "PROCESSED",
					},
					{
						ID:         "1",
						UploadedAt: time.Date(2025, time.April, 27, 14, 30, 0, 0, time.UTC),
						UserID:     1,
						Value: sql.NullFloat64{
							Float64: 0,
							Valid:   false,
						},
						Status: "INVALID",
					},
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/api/user/orders", nil)
			w := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserService := orderMock.NewMockServiceInterface(ctrl)
			mockUserService.EXPECT().
				GetByUser(gomock.Any()).
				Return(test.testRequest.orders, test.testRequest.err).AnyTimes()
			mockAccrualService := accrualMock.NewMockServiceInterface(ctrl)
			mockRewardService := rewardMock.NewMockServiceInterface(ctrl)
			customHandler := NewHandler(mockUserService, mockAccrualService, mockRewardService)
			ctx := context.WithValue(request.Context(), authorization.UserIDKey, 1)
			customHandler.GetByUser(w, request.WithContext(ctx))
			result := w.Result()
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
