package main

import "testing"

func TestUnpackStrManyNums(t *testing.T) {
	_, err := UnpackStr("a12dsj2")
	if err == nil {
		t.Error("expected the error")
	}
}

func TestUnpackStrFirstNum(t *testing.T) {
	_, err := UnpackStr("1a3fd")
	if err == nil {
		t.Error("expected the error")
	}
}

func TestUnpackStrEmptyStr(t *testing.T) {
	r, err := UnpackStr("")
	if len(r) != 0 || err != nil {
		t.Error("expected empty string")
	}
}
