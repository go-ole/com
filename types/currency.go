package types

import "fmt"

// Currency is COM Currency type.
//
// This is not a float. The decimal point is between the left 15 digits and the
// right 4 digits. In base 10. Should only handle maths with other Currency
// objects.
//
// Does not currently support conversions from or to float64.
type Currency struct {
	Value int64
}

// ToString implements string interface and returns string form of currency.
func (c Currency) ToString() string {
	const placement = int64(10000)
	// Remove everything except last 4 digits.
	left := c.Value % placement
	// Shift number right 4 digits.
	right := c.Value / placement
	return fmt.Sprintf("%d.%d", right, left)
}

// Add this currency object to another Currency object.
//
// this + right
func (c Currency) Add(right Currency) Currency {
	return Currency{Value: (c.Value + right.Value)}
}

// Subtract this currency object to another Currency object.
//
// this - right
func (c Currency) Subtract(right Currency) Currency {
	return Currency{Value: (c.Value - right.Value)}
}

// Multiply this currency object to another Currency object.
//
// this * right
func (c Currency) Multiply(right Currency) Currency {
	return Currency{Value: (c.Value * right.Value)}
}

// Divide this currency object to another Currency object.
//
// this / right
func (c Currency) Divide(right Currency) Currency {
	return Currency{Value: (c.Value / right.Value)}
}
