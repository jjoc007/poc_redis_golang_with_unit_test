package main

import (
	"github.com/go-redis/redis"
	"github.com/jjoc007/poc_redis_golang_with_unit_test/redis/config"
	"github.com/jjoc007/poc_redis_golang_with_unit_test/redis/log"
	"github.com/jjoc007/poc_redis_golang_with_unit_test/redis/model"
	"github.com/jjoc007/poc_redis_golang_with_unit_test/redis/services"
	"time"
)

func main()  {
	c := &config.Redis{
		RedisURL:      "127.0.0.1:6379",
		RedisPassword: "",
		Timeout:       10 * time.Second,
		PoolSize:      10,
		DialTimeout:   10 * time.Second,
	}
	logger := log.NewConsole(true)

	config,err := config.NewRedisDBStorage(c, logger)
	if err != nil {
		panic(err)
	}

	personService := services.New(config.GetConnection().(*redis.Client), &logger)

	p1 :=model.Person{
		UID:      "123456789",
		Name:     "Juan",
		LastName: "Orjuela",
	}


	personService.Set(p1)

	p2,err := personService.Get(p1.UID)
	if err != nil {
		panic(err)
	}

	println(p2.UID)
	println(p2.Name)
	println(p2.LastName)
}
