package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/concourse/time-resource/models"
)

func main() {
	now := time.Now()

	var request models.CheckRequest

	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		println("error decoding payload: " + err.Error())
		os.Exit(1)
	}

	duration, err := time.ParseDuration(request.Source.Interval)
	if err != nil {
		println("invalid interval: " + request.Source.Interval + "; " + err.Error())
		os.Exit(1)
	}

	versions := []models.Version{}

	if now.Sub(request.Version.Time) > duration {
		versions = append(versions, models.Version{Time: now})
	}

	json.NewEncoder(os.Stdout).Encode(versions)
}