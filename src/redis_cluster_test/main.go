package main

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.100.103:30004", "192.168.100.103:30005", "192.168.100.103:30009"},
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
	i := 1
	for {
		fmt.Println("number" + strconv.Itoa(i))
		a, err := client.Set("number"+strconv.Itoa(i), i, 0).Result()
		//res, err := client.Get("b").Result()
		fmt.Println(a)
		//fmt.Println(res)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok: ", i)
		}
		i++
		time.Sleep(1 * time.Second)
	}
}

// func main() {
// 	client1 := redis.NewClient(&redis.Options{
// 		Addr:     "192.168.100.103:30004",
// 		Password: "",
// 		DB:       0,
// 	})

// 	pong, err := client1.Ping().Result()
// 	fmt.Println(pong, err)

// 	client2 := redis.NewClient(&redis.Options{
// 		Addr:     "192.168.100.103:30005",
// 		Password: "",
// 		DB:       0,
// 	})

// 	pong, err = client2.Ping().Result()
// 	fmt.Println(pong, err)

// 	client3 := redis.NewClient(&redis.Options{
// 		Addr:     "192.168.100.103:30009",
// 		Password: "",
// 		DB:       0,
// 	})

// 	pong, err = client3.Ping().Result()
// 	fmt.Println(pong, err)

// 	err = client1.Set("c", "d", 0).Err()
// 	if err != nil {
// 		fmt.Println(nil)
// 	}
// }
