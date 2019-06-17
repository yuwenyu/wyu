package sys

import (
	"time"
)

type Configs struct {
	Variable string    `xorm:"not null pk VARCHAR(128)"`
	Value    string    `xorm:"VARCHAR(128)"`
	SetTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	SetBy    string    `xorm:"VARCHAR(128)"`
}
