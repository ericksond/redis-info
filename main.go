package main

import (
	"flag"
	"fmt"
	"os"

	redis "gopkg.in/redis.v4"
)

func main() {

	var addr = flag.String("addr", "", "redis address")
	var password = flag.String("password", "", "redis password")
	flag.Parse()

	client := redis.NewClient(&redis.Options{
		Addr:     *addr,
		Password: *password,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(pong)
	os.Exit(0)
}
