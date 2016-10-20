package main

import "github.com/bep/kittn/auth"

func main() {
	api := auth.Authorize("meowmeowmeow")

	_ = api.GetKitten(2)
}
