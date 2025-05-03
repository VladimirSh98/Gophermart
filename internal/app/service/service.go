package service

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/config"
	"github.com/VladimirSh98/Gophermart.git/internal/app/database"
	"github.com/VladimirSh98/Gophermart.git/internal/app/handler"
	authHandler "github.com/VladimirSh98/Gophermart.git/internal/app/handler/auth"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	operationsRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"
	orderRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
	rewardRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
	userRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/user"
	operationService "github.com/VladimirSh98/Gophermart.git/internal/app/service/operation"
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

	config.Conf = config.Config{}
	err = config.Conf.Load()
	if err != nil {
		sugar.Errorf("Error loading config: %v", err)
		return err
	}

	var db database.DBConnectionStruct
	err = db.OpenConnection(&config.Conf)
	if err != nil {
		sugar.Errorf("Database connection failed: %v", err)
		return err
	}
	defer db.CloseConnection()

	err = db.UpgradeMigrations(&config.Conf)
	if err != nil {
		sugar.Errorf("Database migrations failed: %v", err)
		return err
	}

	customAuthHandler := initAuthHandler(&db)
	customHandker := initHandler(&db)
	router := chi.NewMux()
	router.Use(middleware.Compress)
	router.Use(middleware.Logger)

	addEndpoints(router, customAuthHandler, customHandker)

	return http.ListenAndServe(config.Conf.RunAddress, router)
}

func addEndpoints(router *chi.Mux, customAuthHandler *authHandler.Handler, handler *handler.Handler) {
	router.Route("/api/user", func(router chi.Router) {
		router.Post("/register", customAuthHandler.Register)
		router.Post("/login", customAuthHandler.Login)

		authorizationGroup := router.Group(nil)
		authorizationGroup.Use(authorization.Authorization(handler))
	})

}

func initAuthHandler(db *database.DBConnectionStruct) *authHandler.Handler {
	userRepo := userRepository.Repository{Conn: db.Conn}
	newUserRepo := userService.NewService(userRepo)
	newAuthHandler := authHandler.NewHandler(newUserRepo)
	return newAuthHandler
}

func initHandler(db *database.DBConnectionStruct) *handler.Handler {
	operationsRepo := operationsRepository.Repository{Conn: db.Conn}
	orderRepo := orderRepository.Repository{Conn: db.Conn}
	userRepo := userRepository.Repository{Conn: db.Conn}
	rewardRepo := rewardRepository.Repository{Conn: db.Conn}
	newOperationsRepo := operationService.NewService(operationsRepo)
	newOrderRepo := orderService.NewService(orderRepo)
	newRewardRepo := rewardService.NewService(rewardRepo)
	newUserRepo := userService.NewService(userRepo)
	newHandler := handler.NewHandler(newOperationsRepo, newOrderRepo, newRewardRepo, newUserRepo)
	return newHandler
}
