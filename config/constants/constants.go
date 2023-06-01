package constants

import "time"

type Timeouts time.Duration

const (
	DB_TIMEOUT Timeouts = 10
)
