package repository

import (
	"math"
	"testing"
)

func TestPow10Bounded(t *testing.T) {
	got, overflow := pow10Bounded(0)
	if overflow || got != 1 {
		t.Fatalf("pow10Bounded(0): got %d overflow %v", got, overflow)
	}

	got, overflow = pow10Bounded(1)
	if overflow || got != 10 {
		t.Fatalf("pow10Bounded(1): got %d overflow %v", got, overflow)
	}

	got, overflow = pow10Bounded(18)
	if overflow || got != 1_000_000_000_000_000_000 {
		t.Fatalf("pow10Bounded(18): got %d overflow %v", got, overflow)
	}

	got, overflow = pow10Bounded(19)
	if !overflow || got != math.MaxInt64 {
		t.Fatalf("pow10Bounded(19): got %d overflow %v", got, overflow)
	}
}

func TestCreditRangeForDigits(t *testing.T) {
	min, max := creditRangeForDigits(3)
	if min != 0 || max != 999 {
		t.Fatalf("digits=3: got %d-%d", min, max)
	}

	min, max = creditRangeForDigits(4)
	if min != 1000 || max != 9999 {
		t.Fatalf("digits=4: got %d-%d", min, max)
	}

	min, max = creditRangeForDigits(20)
	if min <= 0 || max != math.MaxInt64 {
		t.Fatalf("digits=20: got %d-%d", min, max)
	}
}
