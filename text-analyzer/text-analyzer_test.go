package main

import (
	"reflect"
	"testing"
)

func TestSplitWords(t *testing.T) {
	text := "Go, кто-то сказал: don't stop."

	got := splitWords(text)

	want := []string{
		"go",
		"кто-то",
		"сказал",
		"don't",
		"stop",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("splitWords(%q) = %#v, want %#v", text, got, want)
	}
}
