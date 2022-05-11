package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yanruogu/gin_demo/pkg/config"
)

type RedisClient struct {
	Client *redis.Client
	cfg    *config.Redis
}

func New(cfg *config.Redis) *RedisClient {
	return &RedisClient{
		cfg: cfg,
	}
}

func (r *RedisClient) Init() error {
	addr := fmt.Sprintf("%s:%s", r.cfg.Host, strconv.Itoa(r.cfg.Port))
	r.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: r.cfg.Password, // no password set
		DB:       r.cfg.Db,       // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.Client.Ping(ctx).Result()
	fmt.Println(res)
	return err
}

func (r *RedisClient) Close() {
	_ = r.Client.Close()
}
