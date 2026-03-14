package main

import (
	"fmt"

	"github.com/openwatch-icu/the-bar/internal/chatroom"
)

func main() {
	fmt.Println("Starting client from cmd/client...")
	chatroom.StartClient()
}
