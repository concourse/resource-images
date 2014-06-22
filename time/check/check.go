package main

import (
	"encoding/json"
	"os"
	"time"
)

type Source struct {
	Interval string `json:"interval"`
}

type Version struct {
	Time time.Time `json:"time"`
}

type Request struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

func main() {
	now := time.Now()

	var request Request

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

	versions := []Version{}

	if now.Sub(request.Version.Time) > duration {
		versions = append(versions, Version{Time: now})
	}

	json.NewEncoder(os.Stdout).Encode(versions)
}
