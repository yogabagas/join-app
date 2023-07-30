package redis

import "github/yogabagas/join-app/pkg/cache"

const (
	schemaDefault = "redis"
	schema        = "redis"
)

func init() {
	cache.Register(schema, NewCache)
}
