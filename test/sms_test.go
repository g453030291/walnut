package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSms(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(900000) + 100000)
	}
}
