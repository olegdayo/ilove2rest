package repository

import (
	"context"

	"github.com/offluck/ilove2rest/internal/entities/user"
)

type RedisClient struct{}

var _ Client = NewRedisClient()

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (*RedisClient) GetUsers(ctx context.Context) ([]user.UserDB, error) {
	panic("unimplemented")
}

func (*RedisClient) GetUser(ctx context.Context, username string) (user.UserDB, error) {
	panic("unimplemented")
}

func (*RedisClient) AddUser(ctx context.Context, userRequest user.UserDB) (user.UserDB, error) {
	panic("unimplemented")
}

func (*RedisClient) UpdateUser(ctx context.Context, username string, newUserDB user.UserDB) (user.UserDB, error) {
	panic("unimplemented")
}

func (*RedisClient) DeleteUser(ctx context.Context, username string) error {
	panic("unimplemented")
}
