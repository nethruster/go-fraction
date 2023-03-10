package fraction

import (
	"errors"
	"math"
)

// Fraction represents a fraction. It is an immutable type.
//
// It is always a valid fraction (never x/0) and it is always simplified.
type Fraction struct {
	numerator   int64
	denominator int64
}

// integer is a generic interface that represents all the integer types of Go.
type integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

var (
	// ErrDivideByZero is returned when trying to divide by a fraction with a value of 0.
	ErrDivideByZero = errors.New("denominator cannot be zero")
	// ErrInvalid is returned when trying to get a fraction from a NaN float.
	ErrInvalid = errors.New("invalid conversion")
	// ErrOutOfRange is returned when trying to get a fraction from a float that is out of the range that this library
	// can represent.
	ErrOutOfRange = errors.New("the number is out of range for this library")
	// ErrZeroDenominator is returned when trying to create a new fraction with 0 as a denominator.
	ErrZeroDenominator = errors.New("denominator cannot be zero")

	// zeroValue is the simplified representation of a fraction with value 0.
	zeroValue = Fraction{
		numerator:   0,
		denominator: 1,
	}
)

// New creates a new fraction with the given numerator and denominator.
//
// It always simplifies the fraction. It returns ErrZeroDenominator if the value of the denominator is 0.
func New[T, K integer](numerator T, denominator K) (Fraction, error) {
	if denominator == 0 {
		return zeroValue, ErrZeroDenominator
	}
	if numerator == 0 {
		return zeroValue, nil
	}

	n := int64(numerator)
	d := int64(denominator)
	if d < 0 {
		d *= -1
		n *= -1
	}
	gcf := gcd(abs(n), d)

	return Fraction{
		numerator:   n / gcf,
		denominator: d / gcf,
	}, nil
}

// FromFloat64 tries to create an exact fraction from the Float64 provided. Keep in mind that the range of numbers that
// floats can represent are bigger than what this fraction type that uses int64 internally can; in that case,
// ErrOutOfRange will be returned. Also keep in mind that floats usually represent approximations to a number; this
// function will try to approximate it as much as possible, but some precision may be lost.
//
// If a NaN is provided, ErrInvalid will be returned.
func FromFloat64(f float64) (Fraction, error) {
	if math.IsNaN(f) {
		return zeroValue, ErrInvalid
	}
	if f < -9.223372036854775e+18 || f > 9.223372036854775e+18 {
		return zeroValue, ErrOutOfRange
	}
	if f > -2.168404344971009e-19 && f < 2.168404344971009e-19 {
		return zeroValue, nil
	}

	// Decompose float64
	bits := math.Float64bits(f)
	isNegative := bits&(1<<63) != 0
	exp := int64((bits>>52)&(1<<11-1)) - 1023
	mantissa := (bits & (1<<52 - 1)) | 1<<52 // Since we discarded tiny values, it'll never be denormalized.

	// Amount of times to shift the mantissa to the right to compensate for the exponent
	shift := 52 - exp

	// Reduce shift and mantissa as far as we can
	for mantissa&1 == 0 && shift > 0 {
		mantissa >>= 1
		shift--
	}

	// Choose whether to shift the numerator or denominator
	var shiftN, shiftD int64 = 0, 0
	if shift > 0 {
		shiftD = shift
	} else {
		shiftN = shift
	}

	// Shift that require larger shifts that what an int64 can hold, or larger than the mantissa itself, will be
	// approximated splitting it between the numerator and denominator.
	if shiftD > 62 {
		shiftD = 62
		shiftN = shift - 62
	} else if shiftN > 52 {
		shiftN = 52
		shiftD = shift - 52
	}

	numerator, denominator := int64(mantissa), int64(1)
	denominator <<= shiftD
	if shiftN < 0 {
		numerator <<= -shiftN
	} else {
		numerator >>= shiftN
	}

	if isNegative {
		numerator *= -1
	}
	return New(numerator, denominator)
}

// Add adds both fractions and returns the result.
func (f1 Fraction) Add(f2 Fraction) Fraction {
	m := lcm(f1.denominator, f2.denominator)
	return Fraction{
		numerator:   f1.numerator*(m/f1.denominator) + f2.numerator*(m/f2.denominator),
		denominator: m,
	}
}

// Divide divides both fractions and returns the result.
//
// It returns ErrDivideByZero if it tries to divide by a fraction with value 0.
func (f1 Fraction) Divide(f2 Fraction) (Fraction, error) {
	f, err := New(f1.numerator*f2.denominator, f1.denominator*f2.numerator)
	if err != nil {
		err = ErrDivideByZero
	}
	return f, err
}

// Equal compares the value of both fractions, returning true if they are equals, and false otherwise.
func (f1 Fraction) Equal(f2 Fraction) bool {
	return f1.numerator == f2.numerator && f1.denominator == f2.denominator
}

// Multiply multiplies both fractions and returns the result.
func (f1 Fraction) Multiply(f2 Fraction) Fraction {
	f, _ := New(f1.numerator*f2.numerator, f1.denominator*f2.denominator)
	return f
}

// Subtract subtracts both fractions and returns the result.
func (f1 Fraction) Subtract(f2 Fraction) Fraction {
	f2.numerator *= -1
	return f1.Add(f2)
}

// Float64 returns the value of the fraction as a float64.
func (f1 Fraction) Float64() float64 {
	return float64(f1.numerator) / float64(f1.denominator)
}

// Denominator returns the fraction denominator.
func (f1 Fraction) Denominator() int64 {
	return f1.denominator
}

// Numerator returns the fraction numerator.
func (f1 Fraction) Numerator() int64 {
	return f1.numerator
}

// abs returns the absolute value of an integer.
func abs[T integer](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

// gcd returns the greatest common divisor of the two numbers. It assumes that both numbers are positive integers.
func gcd(n1, n2 int64) int64 {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	return n1
}

// lcm returns the least common multiple of the two numbers. It assumes that both numbers are positive integers.
func lcm(n1, n2 int64) int64 {
	// Put the largest number in n2 because it's divided first, avoiding overflows in some cases
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	return n1 * (n2 / gcd(n1, n2))
}
