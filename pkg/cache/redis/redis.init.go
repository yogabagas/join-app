package redis

import "github/yogabagas/print-in/pkg/cache"

const (
	schemaDefault = "redis"
	schema        = "redis"
)

func init() {
	cache.Register(schema, NewCache)
}
