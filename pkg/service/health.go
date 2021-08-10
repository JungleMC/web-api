package service

import (
	"context"
	"fmt"
	"github.com/hellofresh/health-go/v4"
	"time"
)

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
