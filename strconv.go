package tape

import "strconv"

// WriteBool appends the string representation of v (either "true" or "false")
// to the buffer, growing the buffer as needed.
// The return value n is the number of bytes written; err is always nil.
// If the buffer becomes too large, WriteBool will panic with [ErrTooLarge].
func (b *Buffer) WriteBool(v bool) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(5)
	if !ok {
		m = b.grow(5)
	}
	b.buf = strconv.AppendBool(b.buf[:m], v)
	return len(b.buf) - m, nil
}

// WriteInt appends the string representation of the integer i in the given
// base to the buffer, growing the buffer as needed.
// The return value n is the number of bytes written; err is always nil.
// If the buffer becomes too large, WriteInt will panic with [ErrTooLarge].
func (b *Buffer) WriteInt(i int64, base int) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(65)
	if !ok {
		m = b.grow(65)
	}
	b.buf = strconv.AppendInt(b.buf[:m], i, base)
	return len(b.buf) - m, nil
}

// WriteUint appends the string representation of the unsigned integer i in the
// given base to the buffer, growing the buffer as needed.
// The return value n is the number of bytes written; err is always nil.
// If the buffer becomes too large, WriteUint will panic with [ErrTooLarge].
func (b *Buffer) WriteUint(i uint64, base int) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(64)
	if !ok {
		m = b.grow(64)
	}
	b.buf = strconv.AppendUint(b.buf[:m], i, base)
	return len(b.buf) - m, nil
}

// WriteFloat appends the string representation of the floating-point number f
// to the buffer, growing the buffer as needed.
// The format fmt and precision prec are as defined in [strconv.FormatFloat].
// The return value n is the number of bytes written; err is always nil.
// If the buffer becomes too large, WriteFloat will panic with [ErrTooLarge].
func (b *Buffer) WriteFloat(f float64, fmt byte, prec, bitSize int) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(128)
	if !ok {
		m = b.grow(128)
	}
	b.buf = strconv.AppendFloat(b.buf[:m], f, fmt, prec, bitSize)
	return len(b.buf) - m, nil
}
