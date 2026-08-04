package main

import (
	"reflect"
	"testing"
)

func TestSplitWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []string
	}{
		{
			name: "пунктуация",
			text: "Go, Go! Go.",
			want: []string{"go", "go", "go"},
		},
		{
			name: "дефис внутри слова",
			text: "Кто-то пришёл",
			want: []string{"кто-то", "пришёл"},
		},
		{
			name: "апостроф внутри слова",
			text: "Don't stop",
			want: []string{"don't", "stop"},
		},
		{
			name: "дефис по краям",
			text: "-word test-",
			want: []string{"word", "test"},
		},
		{
			name: "двойной дефис",
			text: "rock--roll",
			want: []string{"rock", "roll"},
		},
		{
			name: "цифры",
			text: "Go123 version-2",
			want: []string{"go123", "version-2"},
		},
		{
			name: "пустая строка",
			text: "",
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitWords(tt.text)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(
					"splitWords(%q) = %#v, want %#v",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]int
	}{
		{
			name: "повторяющиеся слова",
			text: "Go, go! GO.",
			want: map[string]int{
				"go": 3,
			},
		},
		{
			name: "русский текст",
			text: "Кто-то пришёл, кто-то ушёл.",
			want: map[string]int{
				"кто-то": 2,
				"пришёл": 1,
				"ушёл":   1,
			},
		},
		{
			name: "разные слова",
			text: "Go is simple",
			want: map[string]int{
				"go":     1,
				"is":     1,
				"simple": 1,
			},
		},
		{
			name: "пустой текст",
			text: "",
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countWords(tt.text)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(
					"countWords(%q) = %#v, want %#v",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestMostFrequent(t *testing.T) {
	tests := []struct {
		name   string
		counts map[string]int
		want   string
	}{
		{
			name: "одно наиболее частое слово",
			counts: map[string]int{
				"go":     3,
				"is":     2,
				"simple": 1,
			},
			want: "go",
		},
		{
			name: "одно слово",
			counts: map[string]int{
				"rust": 1,
			},
			want: "rust",
		},
		{
			name:   "пустая карта",
			counts: map[string]int{},
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mostFrequent(tt.counts)

			if got != tt.want {
				t.Fatalf(
					"mostFrequent(%#v) = %q, want %q",
					tt.counts,
					got,
					tt.want,
				)
			}
		})
	}
}
