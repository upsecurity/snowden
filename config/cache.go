package config

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"os"
	"snowden/client/model"
	"time"
)

type redisCache struct {
	host string
	db   int
	exp  time.Duration
}

func (cache redisCache) DeleteVulnerability(id string) error {
	c := cache.getClient()
	numDeleted, err := c.HDel("vulnerabilities", id).Result()
	if numDeleted == 0 {
		return errors.New("vulnerability to delete not found")
	}
	if err != nil {
		return err
	}
	return nil
}

func (cache redisCache) SaveVulnerability(vulnerability model.Cve) error {
	c := cache.getClient()
	vulnerability.ID = uuid.New().String()
	json, err := json.Marshal(vulnerability)
	if err != nil {
		return err
	}
	c.HSet("vulnerabilities", vulnerability.ID, json)

	return nil
}

func (cache redisCache) GetVulnerability(id string) (model.Cve, error) {
	c := cache.getClient()
	vulnerability, err := c.HGet("vulnerabilities", id).Result()
	if err != nil {
		return model.Cve{}, errors.New("vulnerability not found")
	}
	var vuln model.Cve
	err = json.Unmarshal([]byte(vulnerability), &vuln)
	if err != nil {
		return model.Cve{}, err
	}
	return vuln, nil
}

func NewRedisCache(host string, db int, exp time.Duration) model.VulnerabilityService {
	return &redisCache{
		host: host,
		db:   db,
		exp:  exp,
	}
}

func (cache redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       cache.db,
	})
}
