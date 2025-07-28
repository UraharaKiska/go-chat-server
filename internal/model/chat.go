package model

import (
	"time"
)

type MessageInfo struct {
	From string
	Text string
	Timestamp time.Time
}

type ChatInfo struct {
	Name string
}

type Chat struct {
	Info ChatInfo
	Users []string
}




