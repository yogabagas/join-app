package model

import (
	"time"
)

type JWK struct {
	ID        string
	Key       string
	ExpiredAt time.Time
}

type ReadUnexpiredKeyReq struct {
	Time time.Time
}

type ReadUnexpiredKeyResp struct {
	ID        string
	Key       string
	ExpiredAt time.Time
}
