package main

import (
	"reflect"
	"testing"
)

func TestAn(t *testing.T) {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "привет", "ба1", "1аб"}
	got := An(input)
	want := map[string][]string{
		"ба1":    []string{"1аб", "ба1"},
		"листок": []string{"листок", "слиток", "столик"},
		"пятак":  []string{"пятак", "пятка", "тяпка"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
