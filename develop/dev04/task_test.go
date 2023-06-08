package main

import (
	"fmt"
	"testing"
)

var words = []string{
	"одуван",
	"пятак", "пятка", "тяпка",
	"листок", "слиток", "столик",
}

var expt = map[string][]string{
	"листок": {"листок", "слиток", "столик"},
	"пятак":  {"пятак", "пятка", "тяпка"},
	"одуван": {"одуван"},
}

func TestAnagramSize(t *testing.T) {
	rslt := *anagram(&words)

	if len(rslt) != len(expt) {
		fmt.Println(len(rslt), len(expt))
		t.Error("the number of keys is different")
	}

	for key, words := range expt {
		if len(*rslt[key]) != len(words) {
			t.Error("the number of words is different")
		}
	}
}

func TestAnagramValues(t *testing.T) {
	rslt := *anagram(&words)

	for exptKey, exptWords := range expt {
		if _, ok := rslt[exptKey]; !ok {
			t.Error("key not found")
		}

		rsltWords := *rslt[exptKey]

		for i, exptWord := range exptWords {
			rsltWord := rsltWords[i]

			if rsltWord != exptWord {
				t.Error("word not found")
			}
		}
	}
}
