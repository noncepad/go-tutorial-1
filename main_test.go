package main_test

import (
	"log"
	"testing"
)

func TestBasic(t *testing.T) {
	var a int
	//a = 1
	b := 2
	c := a + b
	if c != 3 {
		t.Fatalf("c is not 3; c=%d", c)
	} else {
		log.Print("success!")
	}
}
