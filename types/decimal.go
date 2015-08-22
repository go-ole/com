package types

// Decimal is the COM Decimal type.
type Decimal struct {
	reserved uint16

	// Where the decimal point is, between 0 and 28.
	//
	// So if the number is 12.345, then Scale is 3. If the number is 1.2345,
	// then number is 4.
	Scale int8

	// Whether number is negative.
	// Positive number if 0. This means that the number is negative if it has
	// any value besides 0.
	Sign int8

	// Hi right side number.
	Hi uint32

	// Low is left side number.
	Low uint64
}

// DecimalNegative whether Decimal type is negative.
const DecimalNegative int8 = int8(0x80)
