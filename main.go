package main

import "rust-cmd-webrcon/action"

func main() {
	if err := action.ParseArgs(); err != nil {
		panic(err)
	}

	if err := action.SendCommand(); err != nil {
		panic(err)
	}
}
