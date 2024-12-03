package main

import (
	"go_tutorials/cmd/functions/anonymous"
	"go_tutorials/cmd/functions/callback"
)

func main() {
	anonymous.RunAnonymousFunc()
	callback.Execute()
}
