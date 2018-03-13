package main

import "github.com/kinokoruumu/go-cache/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
