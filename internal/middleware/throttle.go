package middleware

import (
	"log"

	"github.com/throttled/throttled/v2"
	"github.com/throttled/throttled/v2/store/memstore"
)

func NewThrottle() throttled.HTTPRateLimiterCtx {
	store, err := memstore.NewCtx(65536)
	if err != nil {
		log.Fatal(err)
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(20),
		MaxBurst: 5,
	}
	rateLimiter, err := throttled.NewGCRARateLimiterCtx(store, quota)
	if err != nil {
		log.Fatal(err)
	}

	httpRateLimiter := throttled.HTTPRateLimiterCtx{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}

	return httpRateLimiter
}
