package service

import (
	"context"
	"fmt"
	"github.com/JungleMC/web-api/internal/config"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/hellofresh/health-go/v4"
	"log"
	"net/http"
	"time"
)

var Instance *WebApiService

type WebApiService struct {
	rdb *redis.Client
}

func Start() {
	config.Get = &config.Config{}
	err := env.Parse(config.Get)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Get.RedisHost, config.Get.RedisPort),
		Password: config.Get.RedisPassword,
		DB:       config.Get.RedisDatabase,
	})
	defer rdb.Close()

	Instance = &WebApiService{
		rdb: rdb,
	}

	Instance.Bootstrap()
}

func (s *WebApiService) Bootstrap() {
	r := mux.NewRouter()
	r.Handle("/status", s.healthCheck().Handler())
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%v:%v", config.Get.ApiHost, config.Get.ApiPort),
		ReadTimeout:  time.Second * time.Duration(config.Get.ApiReadTimeout),
		WriteTimeout: time.Second * time.Duration(config.Get.ApiWriteTimeout),
	}

	log.Fatal(srv.ListenAndServe())
}

func (s *WebApiService) healthCheck() *health.Health {
	h, _ := health.New(health.WithChecks(
		health.Config{
			Name:      "redis",
			Timeout:   time.Second,
			SkipOnErr: false,
			Check:     s.redisHealthCheck,
		},
	))
	return h
}

func (s *WebApiService) redisHealthCheck(ctx context.Context) error {
	pong, err := s.rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redis ping failed: %w", err)
	}

	if pong != "PONG" {
		return fmt.Errorf("unexpected response for redis ping: %q", pong)
	}
	return nil
}
