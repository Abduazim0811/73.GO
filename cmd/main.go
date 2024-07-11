package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to redis: ", err)
	}
	fmt.Println("Connected to redis: ", pong)

	err = client.Set(ctx, "key", "nimadur", 0).Err()
    if err != nil {
        log.Fatalf("Failed to set key: %v", err)
    }

    val, err := client.Get(ctx, "key").Result()
    if err != nil {
        log.Fatalf("Failed to get key: %v", err)
    }
    fmt.Println("key:", val)

    err = client.Set(ctx, "number", 2, 0).Err()
    if err != nil {
        log.Fatalf("Failed to set number: %v", err)
    }

    newnumber, err := client.Incr(ctx, "number").Result()
    if err != nil {
        log.Fatalf("Failed to increment number: %v", err)
    }
    fmt.Println("number:", newnumber)


	user := map[string]interface{}{
        "id":    "777",
        "name":  "Abduazim Yusufov",
        "email": "abduazim@gmail.com",
    }
    err = client.HSet(ctx, "user:777", user).Err()
    if err != nil {
        log.Fatalf("Failed to set user hash: %v", err)
    }

    err = client.HSet(ctx, "user:777", "name", "Javohir").Err()
    if err != nil {
        log.Fatalf("Failed to update user hash: %v", err)
    }

    userFields, err := client.HGetAll(ctx, "user:777").Result()
    if err != nil {
        log.Fatalf("Failed to get user hash: %v", err)
    }
    fmt.Println("User:", userFields)


	err = client.RPush(ctx,"list", "bir", "ikki", "uch").Err()
	if err != nil {
		log.Fatal("Failed to push elements to list:", err)
	}

	vals, err := client.LRange(ctx, "list", 0, -1).Result()
	if err != nil {
		log.Fatalf("Failed to get elements from list: %v", err)
	}

	fmt.Println("Lists:", vals)

	err = client.LRem(ctx, "list", 1, "ikki").Err()
	if err != nil {
        log.Fatalf("Failed to remove element from list: %v", err)
    }
	vals, err = client.LRange(ctx, "list", 0, -1).Result()
	if err != nil {
		log.Fatalf("Failed to get elements from list: %v", err)
	}

	fmt.Println("Lists:", vals)

	pubsub := client.Subscribe(ctx, "mychannel")
    defer pubsub.Close()


    err = client.Publish(ctx, "mychannel", "hello").Err()
    if err != nil {
        log.Fatalf("Failed to publish message: %v", err)
    }

    msg, err := pubsub.ReceiveMessage(ctx)
    if err != nil {
        log.Fatalf("Failed to receive message: %v", err)
    }
    fmt.Println("Received message:", msg.Payload)

}
