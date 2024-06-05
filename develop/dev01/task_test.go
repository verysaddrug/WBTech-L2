package main

import (
	"testing"
)

func TestTime(t *testing.T) {
	err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		t.Error("Error main function")
	}
}

func TestTimeServer(t *testing.T) {
	err := GetTime("abcd")
	if err == nil {
		t.Error("Error: wrong server without catching error")
	}
}
