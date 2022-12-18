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

func TestEquals(t *testing.T) {
	f1, _ := fraction.New(-19, 27)
	f2, _ := fraction.New(57, -81)
	f3, _ := fraction.New(-57, -81)
	if !f1.Equal(f2) {
		t.Fatal("expected both fractions (-19/27) to be equal, got not equal")
	}
	if f1.Equal(f3) {
		t.Fatal("expected fraction -19/27 not to be equal to 19/27, got equal")
	}

	f1, _ = fraction.New(0, 23)
	f2, _ = fraction.New(0, 2)
	if !f1.Equal(f2) {
		t.Fatal("expected both fractions (0/1) to be equal, got not equal")
	}
}

func TestAdd(t *testing.T) {
	f1, _ := fraction.New(6, 36)
	f2, _ := fraction.New(14, 18)
	compare(t, f1.Add(f2), 17, 18)

	f1, _ = fraction.New(26, 33)
	f2, _ = fraction.New(49, -27)
	compare(t, f1.Add(f2), -305, 297)

	f1, _ = fraction.New(49, 42)
	f2, _ = fraction.New(0, -29)
	compare(t, f1.Add(f2), 7, 6)
}

func TestSubtract(t *testing.T) {
	f1, _ := fraction.New(6, 36)
	f2, _ := fraction.New(14, 18)
	compare(t, f1.Subtract(f2), -11, 18)

	f1, _ = fraction.New(26, 33)
	f2, _ = fraction.New(-49, 27)
	compare(t, f1.Subtract(f2), 773, 297)

	f1, _ = fraction.New(49, 42)
	f2, _ = fraction.New(0, -29)
	compare(t, f1.Subtract(f2), 7, 6)

	f1, _ = fraction.New(-12, 22)
	f2, _ = fraction.New(47, -5)
	compare(t, f1.Subtract(f2), 487, 55)
}

func TestMultiply(t *testing.T) {
	f1, _ := fraction.New(49, 14)
	f2, _ := fraction.New(7, 15)
	compare(t, f1.Multiply(f2), 49, 30)

	f1, _ = fraction.New(26, 33)
	f2, _ = fraction.New(0, 27)
	compare(t, f1.Multiply(f2), 0, 1)

	f1, _ = fraction.New(48, 9)
	f2, _ = fraction.New(6, -16)
	compare(t, f1.Multiply(f2), -2, 1)
}
