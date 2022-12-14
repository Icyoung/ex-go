package errorx

import (
	"errors"
)

var (
	ErrTokenMiss   = errors.New("token miss")
	ErrTokenExpire = errors.New("token expire")
)
