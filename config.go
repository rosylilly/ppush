package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	App     string `json:"app"`
	User    string `json:"user"`
	Message string
	Pids    []int
}

func NewConfig() *Config {
	c := &Config{
		Message: "Process executed",
	}

	homeDir := os.Getenv("HOME")
	bytes, err := ioutil.ReadFile(homeDir + "/.ppush.json")
	if err == nil {
		err := json.Unmarshal(bytes, c)

		if err != nil {
			log.Fatal(err)
		}
	}

	return c
}
