package domain

import "time"

type DaemonState struct {
	Mode      string
	UpdatedAt time.Time
}
