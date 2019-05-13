package main

import "time"

func main() {
	time.Sleep(5 * time.Second) //Wait for postgres initialize
	BeginService()
}
