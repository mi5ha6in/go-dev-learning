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
