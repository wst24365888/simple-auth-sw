package main

import "github.com/ServiceWeaver/weaver"

type UserPayload struct {
	weaver.AutoMarshal
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int
	Username string
	Password string
}
