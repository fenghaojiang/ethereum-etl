package common

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

func Retry(maxRetries uint64, timeout time.Duration, fn func() error) error {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = timeout
	backOffFunc := backoff.WithMaxRetries(b, maxRetries)
	return backoff.Retry(fn, backOffFunc)
}
