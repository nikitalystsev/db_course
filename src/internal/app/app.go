package app

import (
	repositories "SmartShopper-postgres"
	"SmartShopper-postgres/impl"
	implServices "SmartShopper-services/impl"
	"SmartShopper-services/intfRepo"
	"SmartShopper-services/pkg/auth"
	"SmartShopper-services/pkg/hash"
	"SmartShopper-ui/handlers"
	"SmartShopper/internal/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

func Run(configDir string) {
	cfg, err := config.Init(configDir)
	if err != nil {
		return
	}

	var (
		userRepo intfRepo.IUserRepo
	)

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.DBName,
		cfg.Postgres.Password, cfg.Postgres.SSLMode)

	fmt.Printf("dsn: %s\n", dsn)

	db, err := repositories.NewClient(dsn)
	if err != nil {
		fmt.Printf("error client`s connect: %v", err)
		return
	}

	userRepo = impl.NewUserRepo(db, client)

	tokenManager, err := auth.NewTokenManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	hasher := hash.NewPasswordHasher(cfg.Auth.PasswordSalt)

	userService := implServices.NewUserService(
		userRepo,
		tokenManager,
		hasher,
		cfg.Auth.JWT.AccessTokenTTL,
		cfg.Auth.JWT.RefreshTokenTTL,
	)

	handler := handlers.NewHandler(
		userService,
		tokenManager,
		cfg.Auth.JWT.AccessTokenTTL,
		cfg.Auth.JWT.RefreshTokenTTL,
	)

	router := handler.InitRoutes()

	fmt.Println("Server was successfully started!")

	err = router.Run(":" + cfg.Port)
	if err != nil {
		fmt.Printf("error run server: %v", err)
		return
	}
}
