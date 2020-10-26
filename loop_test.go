package main

import "testing"

func TestLoop(t *testing.T) {
	got := loop(0.0001)
	if got != 249995787104.951508 {
		t.Errorf("got: %f, want: %f", got, 249995787104.951508)
	}
}