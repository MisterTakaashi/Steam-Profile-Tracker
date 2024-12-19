package storage

import "time"

type StorageModel struct {
	Date            time.Time
	Game            string
	GameId          uint
	Playtime2Weeks  uint
	PlaytimeForever uint
}
