package tape

import "unsafe"

// UnsafeString returns the contents of the unread portion of the buffer
// as a string without allocating a copy. The returned string is only valid
// until the next buffer modification. If the [Buffer] is a nil pointer,
// it returns "<nil>".
func (b *Buffer) UnsafeString() string {
	if b == nil {
		return "<nil>"
	}
	if b.Len() == 0 {
		return ""
	}
	return unsafe.String(&b.buf[b.off], b.Len())
}
