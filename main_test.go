package testrepo

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
	time.Sleep(10 * time.Millisecond)
}

func TestB(t *testing.T) {
	time.Sleep(10 * time.Millisecond)
}

func TestC(t *testing.T) {
	time.Sleep(10 * time.Millisecond)
	t.Error("error: cannot open a new connection")
	t.Fatal("failed")
}

func TestD(t *testing.T) {
	t.Parallel()
	time.Sleep(200 * time.Millisecond)
}

func TestDD(t *testing.T) {
	t.Parallel()
	time.Sleep(500 * time.Millisecond)
}
func TestE(t *testing.T) {
	t.Run("F", func(t *testing.T) {
		t.Error("error: cannot open a new connection")
	})
	time.Sleep(10 * time.Millisecond)
}
