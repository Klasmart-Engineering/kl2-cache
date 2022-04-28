package kl2cache

import (
	"errors"
	"time"
)

var (
	ErrBadUsageOfKl2Cache = errors.New("bad usage of kl2cache see log for detail")

	MaxCacheExpire = time.Minute * 10
)
