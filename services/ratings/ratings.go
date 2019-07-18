package main

import (
	"log"
)

func rateMovie(title string) int {
	log.Printf("rating %v!", title)
	if len(title)%2 == 0 {
		return 50
	}
	return 100
}
