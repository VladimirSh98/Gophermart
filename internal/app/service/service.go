package service

import (
	accrualClient "github.com/VladimirSh98/Gophermart.git/internal/app/client/accrual"
	"github.com/VladimirSh98/Gophermart.git/internal/app/config"
	"github.com/VladimirSh98/Gophermart.git/internal/app/database"
	authHandler "github.com/VladimirSh98/Gophermart.git/internal/app/handler/auth"
	operationHanlder "github.com/VladimirSh98/Gophermart.git/internal/app/handler/operation"
	orderHandler "github.com/VladimirSh98/Gophermart.git/internal/app/handler/order"
	rewardHandler "github.com/VladimirSh98/Gophermart.git/internal/app/handler/reward"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	"github.com/VladimirSh98/Gophermart.git/internal/app/models"
	operationRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"
	orderRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
	rewardRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
	userRepository "github.com/VladimirSh98/Gophermart.git/internal/app/repository/user"
	accrualService "github.com/VladimirSh98/Gophermart.git/internal/app/service/accrual"
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

	config.Conf = models.Config{}
	err = config.Load(&config.Conf)
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
	customOrderHandler := initOrderHandler(&db)
	customOperationHandler := initOperationHandler(&db)
	customRewardHandler := initRewardHandler(&db)

	router := chi.NewMux()
	router.Use(middleware.Logger)
	router.Use(middleware.Compress)

	addEndpoints(router, customAuthHandler, customOrderHandler, customOperationHandler, customRewardHandler)

	return http.ListenAndServe(config.Conf.RunAddress, router)
}

func addEndpoints(
	router *chi.Mux,
	customAuthHandler *authHandler.Handler,
	customOrderHandler *orderHandler.Handler,
	customOperationHandler *operationHanlder.Handler,
	customRewardHandler *rewardHandler.Handler,
) {
	router.Route("/api/user", func(router chi.Router) {
		router.Post("/register", customAuthHandler.Register)
		router.Post("/login", customAuthHandler.Login)

		authorizationGroup := router.Group(nil)
		authorizationGroup.Use(authorization.Authorization(customAuthHandler))
		authorizationGroup.Get("/orders", customOrderHandler.GetByUser)
		authorizationGroup.Get("/balance", customRewardHandler.GetByUser)
		authorizationGroup.Get("/withdrawals", customOperationHandler.GetByUser)
		authorizationGroup.Post("/balance/withdraw", customOperationHandler.Create)
		authorizationGroup.Post("/orders", customOrderHandler.Create)
	})

}

func initAuthHandler(db *database.DBConnectionStruct) *authHandler.Handler {
	userRepo := userRepository.Repository{Conn: db.Conn}
	newUserService := userService.NewService(userRepo)
	rewardRepo := rewardRepository.Repository{Conn: db.Conn}
	newRewardService := rewardService.NewService(rewardRepo)
	newAuthHandler := authHandler.NewHandler(newUserService, newRewardService)
	return newAuthHandler
}

func initOrderHandler(db *database.DBConnectionStruct) *orderHandler.Handler {
	userRepo := orderRepository.Repository{Conn: db.Conn}
	newUserRepo := orderService.NewService(userRepo)
	newAccrualClient := accrualClient.NewHTTPClient()
	newAccrualService := accrualService.NewService(newAccrualClient)
	rewardRepo := rewardRepository.Repository{Conn: db.Conn}
	newRewardService := rewardService.NewService(rewardRepo)
	newAuthHandler := orderHandler.NewHandler(newUserRepo, newAccrualService, newRewardService)
	return newAuthHandler
}

func initRewardHandler(db *database.DBConnectionStruct) *rewardHandler.Handler {
	rewardRepo := rewardRepository.Repository{Conn: db.Conn}
	newRewardService := rewardService.NewService(rewardRepo)
	newRewardHandler := rewardHandler.NewHandler(newRewardService)
	return newRewardHandler
}

func initOperationHandler(db *database.DBConnectionStruct) *operationHanlder.Handler {
	operationRepo := operationRepository.Repository{Conn: db.Conn}
	newOperationService := operationService.NewService(operationRepo)
	rewardRepo := rewardRepository.Repository{Conn: db.Conn}
	newRewardService := rewardService.NewService(rewardRepo)
	newOperationHandler := operationHanlder.NewHandler(newOperationService, newRewardService)
	return newOperationHandler
}
