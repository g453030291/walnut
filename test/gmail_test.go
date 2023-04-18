package test

import (
	"testing"
	"walnut/app"
)

func TestGetProfile(t *testing.T) {
	profile := app.GetProfile()
	println(profile)
}
