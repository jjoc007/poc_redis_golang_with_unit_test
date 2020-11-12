package config

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type redisDataBase struct {
	databaseConnection redis.Cmdable
	config             *Redis
	logger             zerolog.Logger
}

// NewRedisDBStorage creates and returns redis db connection instance
func NewRedisDBStorage(configDB *Redis, logger zerolog.Logger) (DataBase, error) {
	logger.Debug().Msgf("New instance Redis storage [%s]", configDB.RedisURL)

	dataBase := &redisDataBase{
		config: configDB,
		logger: logger,
	}
	err := dataBase.OpenConnection()
	if err != nil {
		return nil, err
	}
	return dataBase, nil
}

// OpenConnection start redis db connection
func (db *redisDataBase) OpenConnection() error {
	db.logger.Info().Msgf("Starting redisDB connection (%s)", db.config.RedisURL)

	rdb := redis.NewClient(&redis.Options{
		Addr:        db.config.RedisURL,
		Password:    db.config.RedisPassword,
	})

	pong, err := rdb.Ping().Result()
	if err != nil {
		return errors.Wrap(err, "Error on connection to redis")
	}
	db.logger.Info().Msgf("Ping Redis connection: %s", pong)

	db.databaseConnection = rdb
	db.logger.Info().Msg("RedisDB UP")
	return nil
}

// GetConnection get redisDB connection
func (db *redisDataBase) GetConnection() interface{} {
	return db.databaseConnection
}
