package testrepo

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
	time.Sleep(1 * time.Second)
}

func TestB(t *testing.T) {
	time.Sleep(1 * time.Second)
}

func TestC(t *testing.T) {
	time.Sleep(1 * time.Second)
	t.Fatal("failed")
}

func TestD(t *testing.T) {
	time.Sleep(2 * time.Second)
}

func TestE(t *testing.T) {
	time.Sleep(2 * time.Second)
}
