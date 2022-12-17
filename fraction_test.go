package fraction_test

import (
	"testing"

	"github.com/nethruster/go-fraction"
)

func fatalIfErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected nil error got %v instead", err)
	}
}

func compare(t *testing.T, f fraction.Fraction, numerator, denominator int64) {
	t.Helper()
	if f.Numerator() != numerator {
		t.Fatalf("expected numerator value to be %v, got %v", numerator, f.Numerator())
	}
	if f.Denominator() != denominator {
		t.Fatalf("expected denominator value to be %v, got %v", denominator, f.Denominator())
	}
}

func TestNew(t *testing.T) {
	_, err := fraction.New(-3, 5)
	fatalIfErr(t, err)
	_, err = fraction.New(int32(1), uint16(2))
	fatalIfErr(t, err)
	_, err = fraction.New(0, 2)
	fatalIfErr(t, err)

	_, err = fraction.New(1, 0)
	if err != fraction.ErrZeroDenominator {
		t.Fatalf("expected ErrZeroDenominator, got %v", err)
	}
	_, err = fraction.New(0, 0)
	if err != fraction.ErrZeroDenominator {
		t.Fatalf("expected ErrZeroDenominator, got %v", err)
	}
}

func TestNewSimplify(t *testing.T) {
	f, err := fraction.New(402, 21)
	fatalIfErr(t, err)
	compare(t, f, 134, 7)

	f, err = fraction.New(-10, 20)
	fatalIfErr(t, err)
	compare(t, f, -1, 2)

	f, err = fraction.New(6, -9)
	fatalIfErr(t, err)
	compare(t, f, -2, 3)

	f, err = fraction.New(-44, -11)
	fatalIfErr(t, err)
	compare(t, f, 4, 1)

	f, err = fraction.New(0, 9)
	fatalIfErr(t, err)
	compare(t, f, 0, 1)

	f, err = fraction.New(0, -6)
	fatalIfErr(t, err)
	compare(t, f, 0, 1)
}

// func TestEquals(t *testing.T) {
// 	f1, _ := fraction.New(134, 7)
// 	f3, _ := fraction.New(54, 13)

// 	if !f1.Equals(f1) {
// 		t.Fatal("f1 shold be equal to itself.")
// 	}

// }

// func TestAdd(t *testing.T) {
// 	f1, _ := fraction.New(134, 7)
// 	f2, _ := fraction.New(54, 13)}

// 	result := f1.Add(f2)

// }
