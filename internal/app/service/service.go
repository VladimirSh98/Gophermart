package service

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/config"
	"github.com/VladimirSh98/Gophermart.git/internal/app/database"
	"github.com/VladimirSh98/Gophermart.git/internal/app/handler"
	operationsRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operations"
	orderRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
	rewardRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
	userRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/user"
	operationsService "github.com/VladimirSh98/Gophermart.git/internal/app/service/operations"
	orderService "github.com/VladimirSh98/Gophermart.git/internal/app/service/order"
	rewardService "github.com/VladimirSh98/Gophermart.git/internal/app/service/reward"
	userService "github.com/VladimirSh98/Gophermart.git/internal/app/service/user"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

func Run() error {
	var err error
	sugar := zap.S()

	conf := config.Config{}
	err = conf.Load()
	if err != nil {
		sugar.Errorf("Error loading config: %v", err) // fatalf ???
	}

	var db database.DBConnectionStruct
	err = db.OpenConnection(&conf)
	if err != nil {
		sugar.Errorf("Database connection failed: %v", err)
	}
	defer db.CloseConnection()

	err = db.UpgradeMigrations(&conf)
	if err != nil {
		sugar.Errorf("Database migrations failed: %v", err)
	}

	customHandler := initHandler(&db)
	router := chi.NewMux()
	addEndpoints(router, customHandler)

	return http.ListenAndServe(conf.RunAddress, router)
}

func addEndpoints(router *chi.Mux, handler *handler.Handler) {

}

func initHandler(db *database.DBConnectionStruct) *handler.Handler {
	operationsRepo := operationsRepository.Repository{Conn: db.Conn}
	orderRepo := orderRepository.Repository{Conn: db.Conn}
	userRepo := userRepository.Repository{Conn: db.Conn}
	rewardRepo := rewardRepository.Repository{Conn: db.Conn}
	newOperationsRepo := operationsService.NewService(operationsRepo)
	newOrderRepo := orderService.NewService(orderRepo)
	newUserRepo := userService.NewService(userRepo)
	newRewardRepo := rewardService.NewService(rewardRepo)
	newHandler := handler.NewHandler(newOperationsRepo, newOrderRepo, newUserRepo, newRewardRepo)
	return newHandler
}
