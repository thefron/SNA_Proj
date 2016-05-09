package main

import (
	"time"
)

type User struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Event struct {
	Type      string    `json:"type"`
	Actor     User      `json:"actor"`
	Repo      Repo      `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
	Org       *User     `json:"org"`
}
