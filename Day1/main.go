package main

import (
	"100_days_of_go/Day1/log"
)

func main() {
	log.Init("PRODUCTION")
	log.Info("varutha info")
	log.Debug("varutha debug")
	log.Warn("varutha warn")
	log.Error("varutha error")
	log.Fatal("varutha fatal")
	log.Panic("varutha panic")
}
