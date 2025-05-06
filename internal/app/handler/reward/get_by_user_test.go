package reward

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	rewardRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
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
		reward rewardRepo.Reward
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
				reward: rewardRepo.Reward{},
				err:    errors.New("database error"),
			},
		},
		{
			description: "Test #2. Success",
			expect: expect{
				status: http.StatusOK,
				response: GetByUserResponse{
					Current:   10,
					Withdrawn: 5.8,
				},
			},
			testRequest: testRequest{
				reward: rewardRepo.Reward{
					ID:        1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					UserID:    1,
					Balance:   10,
					Withdrawn: 5.8,
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
			mockRewardService := rewardMock.NewMockServiceInterface(ctrl)
			mockRewardService.EXPECT().
				GetByUser(gomock.Any()).
				Return(test.testRequest.reward, test.testRequest.err).AnyTimes()
			customHandler := NewHandler(mockRewardService)
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
				var realResponse GetByUserResponse
				err = json.Unmarshal(body, &realResponse)
				assert.NoError(t, err)
				assert.Equal(t, test.expect.response, realResponse)
			}
		})
	}
}
