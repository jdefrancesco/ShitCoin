package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestShitBlock(t *testing.T) {

	// Create a genesis first block....
	fb := ShitBlock{1, time.Now().String(), 4, "0000", "0"}
	sb, _ := NewShitBlock(fb, 4)
	fmt.Println("New Shit Block has been created! Congrats!")
	spew.Sdump(sb)
}
