package natsort

import (
	"math/bits"
)

func Sort[T ~string](x []T) {
	n := len(x)
	pdqsort(x, 0, n, bits.Len(uint(n)))
}

func IsSorted[T ~string](s []T) bool {
	for i := len(s) - 1; i > 0; i-- {
		if Less(s[i], s[i-1]) {
			return false
		}
	}
	return true
}

type Slice[T ~string] []T

func (p Slice[T]) Len() int           { return len(p) }
func (p Slice[T]) Less(i, j int) bool { return Less(p[i], p[j]) }
func (p Slice[T]) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Less[T ~string](a, b T) bool {
	for {
		a, b = skipPrefix(a, b)
		switch {
		case a == "":
			return b != ""
		case b == "":
			return false
		}

		a1, a2 := firstNonDigit(a)
		if a1 == "" {
			return a < b
		}
		b1, b2 := firstNonDigit(b)
		if b1 == "" {
			return a < b
		}

		an, aok := parseUint(a1)
		if !aok {
			return a < b
		}
		bn, bok := parseUint(b1)
		if !bok {
			return a < b
		}

		switch {
		case an != bn:
			return an < bn
		case len(a1) == len(a) || len(b1) == len(b):
			return a < b
		default:
			a, b = a2, b2
		}
	}
}

func skipPrefix[T ~string](a, b T) (T, T) {
	s := min(len(a), len(b))
	if s == 0 {
		return a, b
	}

	_, _ = a[s-1], b[s-1]

	if isDigit2(a[0], b[0]) {
		return a, b
	}

	for i := 1; i < s; i++ {
		ca, cb := a[i], b[i]
		if isDigit2(ca, cb) {
			return a[i:], b[i:]
		}
	}
	return a[s:], b[s:]
}

func firstNonDigit[T ~string](s T) (T, T) {
	switch {
	case s == "":
		return "", ""

	case !isDigit(s[0]):
		return "", s

	default:
		for i := 1; i < len(s); i++ {
			if !isDigit(s[i]) {
				return s[:i], s[i:]
			}
		}
		return s, ""
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const (
	maxUint64 = 1<<64 - 1
	cutoff    = maxUint64/10 + 1
)

// lightweight version of strconv.ParseUint.
func parseUint[T ~string](s T) (uint64, bool) {
	var n uint64
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !isDigit(c) || n >= cutoff {
			return 0, false
		}

		d := c - '0'
		n *= uint64(10)
		n1 := n + uint64(d)

		if n1 < n || n1 > maxUint64 {
			return maxUint64, false
		}
		n = n1
	}
	return n, true
}

func isDigit(r byte) bool {
	return r^'0' < 10
}

func isDigit2(a, b byte) bool {
	return (a^'0' < 10) || (b^'0' < 10) || a != b
}
