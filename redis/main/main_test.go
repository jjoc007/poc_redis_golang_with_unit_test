package main_test

import (
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/jjoc007/poc_redis_golang_with_unit_test/redis/log"
	"github.com/jjoc007/poc_redis_golang_with_unit_test/redis/model"
	"github.com/jjoc007/poc_redis_golang_with_unit_test/redis/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	client *redis.Client
	personService services.PersonService
)

var (
	key = "123456789"
	val = []byte("{\"uid\":\"123456789\",\"name\":\"Juan\",\"last_name\":\"Orjuela\"}")
	logger = log.NewConsole(true)
)

func Init() {

	mr, err := miniredis.Run()
	if err != nil {
		logger.Err(err).Msgf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
}

func TestSet(t *testing.T) {
	Init()
	exp := time.Duration(0)

	mock := redismock.NewNiceMock(client)
	mock.On("Set", key, val, exp).Return(redis.NewStatusResult("", nil))

	p1 :=model.Person{
		UID:      "123456789",
		Name:     "Juan",
		LastName: "Orjuela",
	}

	personService = services.New(mock, &logger)
	err := personService.Set(p1)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	Init()
	mock := redismock.NewNiceMock(client)
	mock.On("Get", key).Return(redis.NewStringResult(string(val), nil))

	personService = services.New(mock, &logger)
	p2, err := personService.Get(key)
	assert.NoError(t, err)

	assert.Equal(t, "123456789", p2.UID)
	assert.Equal(t, "Juan", p2.Name)
	assert.Equal(t, "Orjuela", p2.LastName)
}