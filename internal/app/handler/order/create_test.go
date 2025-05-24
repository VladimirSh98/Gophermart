package order

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	accrualClient "github.com/VladimirSh98/Gophermart.git/internal/app/client/accrual"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
	accrualMock "github.com/VladimirSh98/Gophermart.git/mocks/accrual"
	orderMock "github.com/VladimirSh98/Gophermart.git/mocks/order"
	rewardMock "github.com/VladimirSh98/Gophermart.git/mocks/reward"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	type expect struct {
		status int
	}
	type testRequest struct {
		body             any
		getOrder         orderRepo.Order
		getOrderErr      error
		createOrderErr   error
		updateOrderErr   error
		accrualRewardErr error
		accrual          *accrualClient.Calculations
		accrualErr       error
	}
	tests := []struct {
		description string
		expect      expect
		testRequest testRequest
	}{
		{
			description: "Test #1. Not valid body",
			expect: expect{
				status: http.StatusUnprocessableEntity,
			},
			testRequest: testRequest{
				body: struct{}{},
			},
		},
		{
			description: "Test #2. Not valid number",
			expect: expect{
				status: http.StatusUnprocessableEntity,
			},
			testRequest: testRequest{
				body: "1",
			},
		},
		{
			description: "Test #3. Existing order",
			expect: expect{
				status: http.StatusOK,
			},
			testRequest: testRequest{
				body: 4539682995824395,
				getOrder: orderRepo.Order{
					ID:     "4539682995824395",
					UserID: 1,
				},
			},
		},
		{
			description: "Test #4. Existing order by another user",
			expect: expect{
				status: http.StatusConflict,
			},
			testRequest: testRequest{
				body: 4539682995824395,
				getOrder: orderRepo.Order{
					ID:     "4539682995824395",
					UserID: 11,
				},
			},
		},
		{
			description: "Test #5. Get order error",
			expect: expect{
				status: http.StatusInternalServerError,
			},
			testRequest: testRequest{
				body:        4539682995824395,
				getOrder:    orderRepo.Order{},
				getOrderErr: errors.New("error"),
			},
		},
		{
			description: "Test #6. Success",
			expect: expect{
				status: http.StatusAccepted,
			},
			testRequest: testRequest{
				body:        4539682995824395,
				getOrder:    orderRepo.Order{},
				getOrderErr: sql.ErrNoRows,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			jsonBody, _ := json.Marshal(test.testRequest.body)
			request := httptest.NewRequest(http.MethodPost, "/api/user/orders", bytes.NewReader(jsonBody))
			w := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockOrderService := orderMock.NewMockServiceInterface(ctrl)
			mockOrderService.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(
				test.testRequest.getOrder, test.testRequest.getOrderErr,
			).AnyTimes()
			mockOrderService.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(
				test.testRequest.createOrderErr,
			).AnyTimes()
			mockOrderService.EXPECT().UpdateByID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
				test.testRequest.updateOrderErr,
			).AnyTimes()

			mockAccrualService := accrualMock.NewMockServiceInterface(ctrl)
			mockAccrualService.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(
				test.testRequest.accrual, test.testRequest.accrualErr,
			).AnyTimes()

			mockRewardService := rewardMock.NewMockServiceInterface(ctrl)
			mockRewardService.EXPECT().AccrueReward(gomock.Any(), gomock.Any(), gomock.Any()).Return(
				test.testRequest.accrualRewardErr,
			).AnyTimes()

			customHandler := NewHandler(mockOrderService, mockAccrualService, mockRewardService)
			ctx := context.WithValue(request.Context(), authorization.UserIDKey, 1)
			customHandler.Create(w, request.WithContext(ctx))
			result := w.Result()
			defer result.Body.Close()
			assert.Equal(t, test.expect.status, result.StatusCode)
		})
	}
}
