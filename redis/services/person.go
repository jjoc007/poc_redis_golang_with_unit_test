package services

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/jjoc007/poc_redis_golang_with_unit_test/redis/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type PersonService interface {
	Set(person model.Person) error
	Get(string) (*model.Person,error)
}

// New creates and returns a new lock service instance
func New(redisDatabase redis.Cmdable,
	logger *zerolog.Logger,
) PersonService {
	return &personService{
		redisDatabase: redisDatabase,
		logger: logger,
	}
}

type personService struct {
	redisDatabase redis.Cmdable
	logger        *zerolog.Logger
}

func (s *personService) Set(person model.Person) error {
	out, err := json.Marshal(person)
	if err != nil {
		return errors.Wrapf(err, "Error when it set document on redis. Key: [%s] ", person.UID)
	}
	err = s.redisDatabase.Set(person.UID, out, 0).Err()
	if err != nil {
		return errors.Wrapf(err, "Error when it set document on redis. Key: [%s] ", person.UID)
	}
	return nil
}

func (s *personService) Get(uid string) (person *model.Person, err error) {
	person = &model.Person{}
	val, err := s.redisDatabase.Get(uid).Result()
	if err == redis.Nil {
		s.logger.Debug().Msgf("key does not exist [%s]", uid)
		return
	} else if err != nil {
		return
	}
	err = json.Unmarshal([]byte(val), person)
	if err != nil {
		return
	}
	return
}