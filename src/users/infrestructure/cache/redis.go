package database

import (
	"context"
	"encoding/json"
	"fmt"
	entities "github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type UserRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewUserRepository(client *redis.Client) *UserRepository {
	return &UserRepository{
		client: client,
		ctx:    context.Background(),
	}
}

const (
	USERS_KEY         = "all_users"
	USERS_PAGE_PREFIX = "users_page_"
	LAST_UPDATE_KEY   = "users_last_update"
	CACHE_TTL         = 24 * time.Hour
)

func (r *UserRepository) SetUsers(users []entities.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("failed to marshal users: %v", err)
	}

	err = r.client.Set(r.ctx, USERS_KEY, data, CACHE_TTL).Err()
	if err != nil {
		return fmt.Errorf("failed to set users in redis: %v", err)
	}

	return nil
}

func (r *UserRepository) GetUsers() ([]entities.User, error) {
	data, err := r.client.Get(r.ctx, USERS_KEY).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get users from redis: %v", err)
	}

	var users []entities.User
	err = json.Unmarshal([]byte(data), &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal users: %v", err)
	}

	return users, nil
}

func (r *UserRepository) SetUsersPaginated(page int, users []entities.User) error {
	key := fmt.Sprintf("%s%d", USERS_PAGE_PREFIX, page)

	data, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("failed to marshal users: %v", err)
	}

	err = r.client.Set(r.ctx, key, data, CACHE_TTL).Err()
	if err != nil {
		return fmt.Errorf("failed to set paginated users in redis: %v", err)
	}

	return nil
}

func (r *UserRepository) GetUsersPaginated(page int) ([]entities.User, error) {
	key := fmt.Sprintf("%s%d", USERS_PAGE_PREFIX, page)

	data, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get paginated users from redis: %v", err)
	}

	var users []entities.User
	err = json.Unmarshal([]byte(data), &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal users: %v", err)
	}

	fmt.Println("se solicito desde redis ")

	return users, nil
}

func (r *UserRepository) DeleteUsersCache() error {
	err := r.client.Del(r.ctx, USERS_KEY).Err()
	if err != nil {
		return fmt.Errorf("failed to delete users cache: %v", err)
	}

	keys, err := r.client.Keys(r.ctx, USERS_PAGE_PREFIX+"*").Result()
	if err != nil {
		return fmt.Errorf("failed to get paginated keys: %v", err)
	}

	if len(keys) > 0 {
		err = r.client.Del(r.ctx, keys...).Err()
		if err != nil {
			return fmt.Errorf("failed to delete paginated cache: %v", err)
		}
	}

	return nil
}

func (r *UserRepository) SetLastUpdateTimestamp(timestamp int64) error {
	err := r.client.Set(r.ctx, LAST_UPDATE_KEY, timestamp, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set last update timestamp: %v", err)
	}
	return nil
}

func (r *UserRepository) GetLastUpdateTimestamp() (int64, error) {
	result, err := r.client.Get(r.ctx, LAST_UPDATE_KEY).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get last update timestamp: %v", err)
	}

	timestamp, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse timestamp: %v", err)
	}

	return timestamp, nil
}
