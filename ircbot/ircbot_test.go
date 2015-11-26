package ircbot

import (
	"testing"
)

func TestCreation(t *testing.T) {
	bot := New("johnny")
	if bot == nil {
		t.Fatal("bot == nil, want !nil")
	}
}
