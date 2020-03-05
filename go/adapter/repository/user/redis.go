package user

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(host string, port int) *RedisRepository {
	return &RedisRepository{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: "",
			DB:       0,
		}),
	}
}

func (r *RedisRepository) Store(user *domain.User) error {
	id := user.Id()
	firstName := user.Name().FirstName()
	lastName := user.Name().LastName()

	// key: id
	// value: hash map
	value := map[string]interface{}{
		"firstName": firstName,
		"lastName":  lastName,
	}

	if err := r.client.HSet(id, value).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RedisRepository) FindById(id *domain.Id) (*domain.User, error) {

	result := r.client.HGetAll(id.Value())
	if result.Err() != nil {
		return nil, result.Err()
	}

	// no such id
	if len(result.Val()) < 1 {
		return nil, fmt.Errorf("Couldn't find specified user")
	}

	// check whether hash key exists in results
	for _, key := range []string{"firstName", "lastName"} {
		if _, ok := result.Val()[key]; !ok {
			return nil, fmt.Errorf("Couldn't find %s", key)
		}
	}

	user, err := domain.NewUser(
		domain.NewId(id.Value()),
		result.Val()["firstName"],
		result.Val()["lastName"],
	)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (r *RedisRepository) Update(user *domain.User) error {
	return r.Store(user)
}
