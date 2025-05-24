package operation

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	rewardRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
	operationMock "github.com/VladimirSh98/Gophermart.git/mocks/operation"
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
		body               any
		reward             rewardRepo.Reward
		getRewardErr       error
		operationCreateErr error
		updateRewardErr    error
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
				body:         struct{}{},
				reward:       rewardRepo.Reward{},
				getRewardErr: nil,
			},
		},
		{
			description: "Test #2. Not valid order number",
			expect: expect{
				status: http.StatusUnprocessableEntity,
			},
			testRequest: testRequest{
				body: CreateRequest{
					Order: "1",
					Sum:   1,
				},
				reward:       rewardRepo.Reward{},
				getRewardErr: nil,
			},
		},
		{
			description: "Test #3. Negative sum",
			expect: expect{
				status: http.StatusUnprocessableEntity,
			},
			testRequest: testRequest{
				body: CreateRequest{
					Order: "4539682995824395",
					Sum:   -31,
				},
				reward:       rewardRepo.Reward{},
				getRewardErr: nil,
			},
		},
		{
			description: "Test #4. Too much sum",
			expect: expect{
				status: http.StatusPaymentRequired,
			},
			testRequest: testRequest{
				body: CreateRequest{
					Order: "4539682995824395",
					Sum:   31,
				},
				reward: rewardRepo.Reward{
					ID:        1,
					UserID:    1,
					Balance:   0,
					Withdrawn: 0,
				},
				getRewardErr: nil,
			},
		},
		{
			description: "Test #5. Reward get user error",
			expect: expect{
				status: http.StatusInternalServerError,
			},
			testRequest: testRequest{
				body: CreateRequest{
					Order: "4539682995824395",
					Sum:   1,
				},
				reward: rewardRepo.Reward{
					ID:        1,
					UserID:    1,
					Balance:   10,
					Withdrawn: 0,
				},
				getRewardErr: errors.New("user error"),
			},
		},
		{
			description: "Test #6. Operation create error",
			expect: expect{
				status: http.StatusInternalServerError,
			},
			testRequest: testRequest{
				body: CreateRequest{
					Order: "4539682995824395",
					Sum:   1,
				},
				reward: rewardRepo.Reward{
					ID:        1,
					UserID:    1,
					Balance:   10,
					Withdrawn: 0,
				},
				getRewardErr:       nil,
				operationCreateErr: errors.New("operation error"),
			},
		},
		{
			description: "Test #7. Update balance error",
			expect: expect{
				status: http.StatusInternalServerError,
			},
			testRequest: testRequest{
				body: CreateRequest{
					Order: "4539682995824395",
					Sum:   1,
				},
				reward: rewardRepo.Reward{
					ID:        1,
					UserID:    1,
					Balance:   10,
					Withdrawn: 0,
				},
				getRewardErr:       nil,
				operationCreateErr: nil,
				updateRewardErr:    errors.New("update error"),
			},
		},
		{
			description: "Test #8. Success",
			expect: expect{
				status: http.StatusOK,
			},
			testRequest: testRequest{
				body: CreateRequest{
					Order: "4539682995824395",
					Sum:   1,
				},
				reward: rewardRepo.Reward{
					ID:        1,
					UserID:    1,
					Balance:   10,
					Withdrawn: 0,
				},
				getRewardErr:       nil,
				operationCreateErr: nil,
				updateRewardErr:    nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			jsonBody, _ := json.Marshal(test.testRequest.body)
			request := httptest.NewRequest(http.MethodPost, "/api/user/balance/withdraw", bytes.NewReader(jsonBody))
			w := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockOperationService := operationMock.NewMockServiceInterface(ctrl)
			mockOperationService.EXPECT().
				Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(test.testRequest.operationCreateErr).AnyTimes()

			mockRewardService := rewardMock.NewMockServiceInterface(ctrl)
			mockRewardService.EXPECT().
				GetByUser(gomock.Any(), 1).
				Return(test.testRequest.reward, test.testRequest.getRewardErr).AnyTimes()
			mockRewardService.EXPECT().
				UpdateByUser(gomock.Any(), 1, gomock.Any(), gomock.Any()).
				Return(test.testRequest.updateRewardErr).AnyTimes()

			customHandler := NewHandler(mockOperationService, mockRewardService)
			ctx := context.WithValue(request.Context(), authorization.UserIDKey, 1)
			customHandler.Create(w, request.WithContext(ctx))
			result := w.Result()
			defer result.Body.Close()
			assert.Equal(t, test.expect.status, result.StatusCode)
		})
	}
}
