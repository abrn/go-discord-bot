package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func testFunc() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		OnConnect(AppContext, cn*redis.Conn),
	})
	Database = rdb
	log.Printf("Connected to Redis cache successfully")
}

func hookOnConnect(ctx context.Context, cn *redis.Conn) error {

}
