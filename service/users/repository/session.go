package repository

import (
	"context"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, userUUID string) error
	DeleteSession(ctx context.Context, userUUID string) error
}
