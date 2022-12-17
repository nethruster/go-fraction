package fraction

import "errors"

type Fraction struct {
	numerator   int64
	denominator int64
}

type integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

var (
	ErrZeroDenominator = errors.New("denominator cannot be zero")

	zeroValue = Fraction{
		numerator:   0,
		denominator: 1,
	}
)

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

	gcf := n
	if gcf < 0 {
		gcf *= -1
	}
	for tmp := d; tmp != 0; {
		gcf, tmp = tmp, gcf%tmp
	}

	return Fraction{
		numerator:   n / gcf,
		denominator: d / gcf,
	}, nil
}

func (f Fraction) Numerator() int64 {
	return f.numerator
}

func (f Fraction) Denominator() int64 {
	return f.denominator
}
