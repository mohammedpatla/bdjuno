package utils

import (
	"github.com/rs/zerolog/log"
)

const ErrValidatorNotFound = "error while getting validator: rpc error: code = %s desc = rpc error: code = %s"

// WatchMethod allows to watch for a method that returns an error.
// It executes the given method in a goroutine, logging any error that might raise.
func WatchMethod(method func() error) {
	go func() {
		err := method()
		if err != nil {
			log.Error().Err(err).Send()
		}
	}()
}
