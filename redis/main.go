package main

import (
	"fmt"
	"gopkg.in/redis.v5"
	//"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err :=client.Get("hello").Result()
	fmt.Println("------", val,"-------")

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//time , err := time.ParseDuration("10s")
	//fmt.Println(time)
	//err = client.Set("hello","world",time).Err()
	//if err != nil {
	//	panic(err)
	//}


}
