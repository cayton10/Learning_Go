package main

import (
	c "oop/objects"
)

// Entry point
func main() {
	comic := c.Comic{
		Publisher:  "Marvel",
		Writer:     "Stan Lee",
		Artist:     "Jack Kirby",
		Title:      "Spiderman",
		Year:       1986,
		PageNumber: 32,
		Grade:      8.7,
	}
	c.PrintComicInfo(comic)
}
