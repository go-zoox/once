package once

import (
	"math/rand"
	"testing"
)

func TestDo(t *testing.T) {
	v := 1
	name := "do"
	for i := 0; i < 100; i++ {
		Do(name, func() {
			v += 1
		})
	}

	if v != 2 {
		t.Errorf("expect %v, got %v", 2, v)
	}
}

func TestGet(t *testing.T) {
	v := tGet()
	for i := 0; i < 100; i++ {
		vt := tGet()
		if v != vt {
			t.Fatalf("expect %v, got %v", v, vt)
		}
	}

	if runsIn != 1 {
		t.Errorf("expect %v, got %v", 1, runsIn)
	}

	if runsOut != 101 {
		t.Errorf("expect %v, got %v", 100, runsOut)
	}
}

func randomInt(max, min int) int {
	return min + rand.Intn(max-min)
}

func tNew() int {
	return randomInt(100, 1)
}

var runsOut = 0
var runsIn = 0

func tGet() int {
	runsOut += 1
	return Get("tget", func() int {
		runsIn += 1
		return tNew()
	})
}
