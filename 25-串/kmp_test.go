package kmp

import (
	"fmt"
	"strings"
	"testing"
)

func TestBruteForce(t *testing.T) {
	text := "Hello World"
	pattern := "or"
	//fmt.Println(indexOf(text, pattern))
	//fmt.Println(indexOf1(text, pattern))
	fmt.Println(indexOf2(text, pattern))
	fmt.Println(strings.Index(text, pattern))
}

func TestKmp(t *testing.T) {
	text := "Hello World"
	pattern := "or"
	fmt.Println(kmp(text, pattern))
	fmt.Println(strings.Index(text, pattern))
}
