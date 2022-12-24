# 📚 go-fraction
[![Go Reference](https://pkg.go.dev/badge/github.com/nethruster/go-fraction.svg)](https://pkg.go.dev/github.com/nethruster/go-fraction)

go-fractions is a Go library for working with fractions. It provides a
`Fraction` type for representing fractions, as well as a set of functions for
performing operations on fractions such as addition, subtraction,
multiplication, and division.

## 🔨 Usage
To use go-fractions, import the package in your Go code:

```go
import "github.com/nethruster/go-fraction"
```

You can create a new fraction using the `New` function

```go
f1, err := fraction.New(1,2) // 1/2, nil
f2, err := fraction.New(2,3) // 2/3, nil
_, err := fraction.New(1,0)  // ErrZeroDenominator
```

You can perform operations on fractions using the provided functions:

```go
f3 := f1.Add(f2)         // 7/6
f4 := f1.Subtract(f2)    // -1/6
f5 := f1.Multiply(f2)    // 1/3
f6, err := f1.Divide(f2) // 3/4, nil
```

You can also convert a fraction to a float, or vice versa, using `Float64` and `FromFloat64` functions:

```go
floatValue := f1.Float64() // 0.5
f7, err := fraction.FromFloat64(0.5) // 1/2, nil
```

### 🤔 Rationale
The Fraction type in go-fractions aims to provide a lightweight, primitive-like representation of fractions. As a
result, it has limitations in terms of precision and can overflow when performing certain operations. If you need a
type that can represent all rational numbers without regard for memory and CPU consumption, consider using the
`big.Rat` type from the standard library.

Please note that these limitations should not be an issue for most use cases, and go-fractions provides a convenient
and efficient way to work with fractions in Go. However, it is important to be aware of these limitations and choose
the appropriate type based on your specific needs.

## 📜 Documentation
For more detailed documentation and a full list of functions, see the
[reference page](https://pkg.go.dev/github.com/nethruster/go-fraction).

## 🤝 Contributing
We welcome contributions to go-fractions. If you have an idea for a new
feature or a bug fix, please open an issue on the
[issues page](https://github.com/nethruster/go-fraction/issues).

🎉 Thank you for considering contributing to go-fractions!
