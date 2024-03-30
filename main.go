package main

import (
	"encoding/json"
	"os"

	"github.com/f-taxes/binance_conversion/ctl"
	"github.com/f-taxes/binance_conversion/global"
	"github.com/kataras/golog"
)

func init() {
	manifestContent, err := os.ReadFile("./manifest.json")

	if err != nil {
		golog.Fatalf("Failed to read manifest file: %v", err)
		os.Exit(1)
	}

	err = json.Unmarshal(manifestContent, &global.Plugin)

	if err != nil {
		golog.Fatalf("Failed to parse manifest: %v", err)
		os.Exit(1)
	}
}

func main() {
	ctl.Start(global.Plugin.Ctl.Address)
}
