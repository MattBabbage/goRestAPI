package main

import (
	"testing"

	"github.com/MattBabbage/goRestAPI/internal/weather"
)

func TestHello(t *testing.T) {
	got := weather.GetWeather()
	want := "Sunny!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
