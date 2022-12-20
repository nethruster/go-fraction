package fraction_test

import (
	"math"
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

func approxFloat(t *testing.T, fr fraction.Fraction, expected, precision float64) {
	t.Helper()
	fl := fr.Float64()
	if fl < expected-precision || fl > expected+precision {
		t.Fatalf("expected fraction around %v with %v of error, got %v", expected, precision, fl)
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

func TestDivide(t *testing.T) {
	f1, _ := fraction.New(49, 14)
	f2, _ := fraction.New(7, 15)
	result, err := f1.Divide(f2)
	fatalIfErr(t, err)
	compare(t, result, 15, 2)

	f1, _ = fraction.New(26, 33)
	f2, _ = fraction.New(0, 27)
	if _, err = f1.Divide(f2); err != fraction.ErrDivideByZero {
		t.Fatalf("expected ErrDivideByZero, got %v", err)
	}

	f1, _ = fraction.New(48, 9)
	f2, _ = fraction.New(6, -16)
	result, err = f1.Divide(f2)
	fatalIfErr(t, err)
	compare(t, result, -128, 9)
}

func TestFloat64(t *testing.T) {
	f, _ := fraction.New(49, 14)
	if f.Float64() != 3.5 {
		t.Fatalf("expected 3.5, got %v", f.Float64())
	}
	f, _ = fraction.New(0, -27)
	if f.Float64() != 0 {
		t.Fatalf("expected 0, got %v", f.Float64())
	}
	f, _ = fraction.New(8, -64)
	if f.Float64() != -0.125 {
		t.Fatalf("expected -0.125, got %v", f.Float64())
	}
}

func TestFromFloat64(t *testing.T) {
	f, err := fraction.FromFloat64(0)
	fatalIfErr(t, err)
	compare(t, f, 0, 1)

	f, err = fraction.FromFloat64(-0)
	fatalIfErr(t, err)
	compare(t, f, 0, 1)

	f, err = fraction.FromFloat64(1)
	fatalIfErr(t, err)
	compare(t, f, 1, 1)

	f, err = fraction.FromFloat64(-1)
	fatalIfErr(t, err)
	compare(t, f, -1, 1)

	f, err = fraction.FromFloat64(1.25)
	fatalIfErr(t, err)
	compare(t, f, 5, 4)

	f, err = fraction.FromFloat64(-1.25)
	fatalIfErr(t, err)
	compare(t, f, -5, 4)

	f, err = fraction.FromFloat64(4.5e10)
	fatalIfErr(t, err)
	compare(t, f, 45000000000, 1)

	f, err = fraction.FromFloat64(-4.5e10)
	fatalIfErr(t, err)
	compare(t, f, -45000000000, 1)

	// 4.5e-10 cannot be represented in a float64, the closest representation is
	// 2^(-32) * 1.1110111011000111101111010101000100101011010101110010 (base 2), which is
	// 4.4999999999999999700744318526239758082585495913008344359695911407470703125 * 10^(-10). The fractions in this
	// library cannot represent real numbers with arbitrary precision, so it will approximate the result.
	f, err = fraction.FromFloat64(4.5e-10)
	fatalIfErr(t, err)
	approxFloat(t, f, 4.5e-10, 1e-19)

	f, err = fraction.FromFloat64(-4.5e-10)
	fatalIfErr(t, err)
	approxFloat(t, f, -4.5e-10, 1e-19)

	// Max number that float64 can represent that fits in an int64.
	// Confusingly, printing this float returns 9223372036854775000, but this is an approximation, because if we do the
	// correct conversion based on the binary data following the IEEE 754 standard, we can see that the number that the
	// float holds it's 2^62 * 1.1111111111111111111111111111111111111111111111111111 (base 2), which is exactly
	// 9223372036854774784.
	f, err = fraction.FromFloat64(9223372036854774784)
	fatalIfErr(t, err)
	compare(t, f, 9223372036854774784, 1)

	f, err = fraction.FromFloat64(-9223372036854774784)
	fatalIfErr(t, err)
	compare(t, f, -9223372036854774784, 1)

	f, err = fraction.FromFloat64(math.Pow(2, -62))
	fatalIfErr(t, err)
	compare(t, f, 1, 1<<62)

	f, err = fraction.FromFloat64(math.Pow(2, -62) * (-1))
	fatalIfErr(t, err)
	compare(t, f, -1, 1<<62)

	f, err = fraction.FromFloat64(math.Pow(2, -63))
	fatalIfErr(t, err)
	compare(t, f, 0, 1)

	f, err = fraction.FromFloat64(math.Pow(2, -63) * (-1))
	fatalIfErr(t, err)
	compare(t, f, 0, 1)

	if _, err = fraction.FromFloat64(9223372036854776000); err != fraction.ErrOutOfRange {
		t.Fatalf("expected ErrOutOfRange, got %v", err)
	}
	if _, err = fraction.FromFloat64(-9223372036854776000); err != fraction.ErrOutOfRange {
		t.Fatalf("expected ErrOutOfRange, got %v", err)
	}
	if _, err = fraction.FromFloat64(math.Inf(1)); err != fraction.ErrOutOfRange {
		t.Fatalf("expected ErrOutOfRange, got %v", err)
	}
	if _, err = fraction.FromFloat64(math.Inf(-1)); err != fraction.ErrOutOfRange {
		t.Fatalf("expected ErrOutOfRange, got %v", err)
	}
	if _, err = fraction.FromFloat64(math.NaN()); err != fraction.ErrInvalid {
		t.Fatalf("expected ErrInvalid, got %v", err)
	}
}
