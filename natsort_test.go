package natsort

import (
	"bytes"
	_ "embed"
	"math/rand"
	"sort"
	"testing"
)

func TestLess(t *testing.T) {
	testCases := []struct {
		a, b string
	}{
		{"", "a"},
		{"a", "b"},
		{"a", "aa"},
		{"a0", "a1"},
		{"a0", "a00"},
		{"a00", "a01"},
		{"a01", "a2"},
		{"a01x", "a2x"},
		// Only the last number matters.
		{"a0b00", "a00b1"},
		{"a0b00", "a00b01"},
		{"a00b0", "a0b00"},
		{"a00b00", "a0b01"},
		{"a00b00", "a0b1"},
	}
	for _, tc := range testCases {
		if !Less(tc.a, tc.b) {
			t.Errorf("Less(%q, %q) returned false", tc.a, tc.b)
		}
		if Less(tc.a, tc.b) == Less(tc.b, tc.a) {
			t.Errorf("Bad result! %q vs %q", tc.a, tc.b)
		}
	}
}

func TestLessEqual(t *testing.T) {
	testCases := []struct {
		a, b string
	}{
		{"a", "a"},
		{"aa", "aa"},
		{"a01", "a01"},
		{"a01x", "a01x"},
		{"a00b01", "a0b01"},
	}
	for _, tc := range testCases {
		if Less(tc.a, tc.b) != Less(tc.b, tc.a) {
			t.Errorf("Bad result! %q vs %q", tc.a, tc.b)
		}
	}
}

func TestSort(t *testing.T) {
	t.Run("data.txt", func(t *testing.T) {
		data := copyFrom(testDataGolden)
		Sort(data)
		checkResult(t, data, testDataGolden)
	})

	t.Run("doc.txt", func(t *testing.T) {
		data := copyFrom(testDocGolden)
		Sort(data)
		checkResult(t, data, testDocGolden)
	})
}

//go:noinline
func lessLex(a, b string) bool {
	return a < b
}

func BenchmarkLessLex(b *testing.B) {
	b.ReportAllocs()

	b.Run("short", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if lessLex("a01a2", "a01a01") {
				b.Fatal("unexpected result")
			}
		}
	})

	b.Run("long", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if lessLex("a01a01a01a01a01a01a01a01a01a2", "a01a01a01a01a01a01a01a01a01a01") {
				b.Fatal("unexpected result")
			}
		}
	})
}

func BenchmarkLess(b *testing.B) {
	b.ReportAllocs()

	b.Run("short", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if Less("a01a2", "a01a01") {
				b.Fatal("unexpected result")
			}
		}
	})

	b.Run("long", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if Less("a01a01a01a01a01a01a01a01a01a2", "a01a01a01a01a01a01a01a01a01a01") {
				b.Fatal("unexpected result")
			}
		}
	})
}

func BenchmarkSort(b *testing.B) {
	data := bench(b, testDataGolden, func(list []string) {
		Sort(list)
	})
	checkResult(b, data, testDataGolden)
}

func BenchmarkSortString(b *testing.B) {
	data := bench(b, testDataGolden, func(list []string) {
		sort.Sort(Slice[string](list))
	})
	checkResult(b, data, testDataGolden)
}

func BenchmarkSortStdlib(b *testing.B) {
	bench(b, testDataGolden, func(list []string) {
		sort.Strings(list)
	})
	// no need to check stdlib impl
}

func bench(b *testing.B, input []string, fn func(list []string)) []string {
	b.Helper()
	b.StopTimer()

	data := make([]string, len(input))
	r := rand.New(rand.NewSource(69420))
	copy(data, input)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})
		fn(data)
	}
	return data
}

func checkResult(tb testing.TB, have, want []string) {
	tb.Helper()

	if len(have) != len(want) {
		tb.Fatalf("have %d want %d", len(have), len(want))
	}
	for i := range have {
		if have[i] != want[i] {
			tb.Errorf("at %d have %q want %q", i, have[i], want[i])
		}
	}
}

func copyFrom(s []string) []string {
	res := make([]string, len(s))
	copy(res, s)
	return res
}

var (
	//go:embed testdata/data.txt
	dataFile []byte
	//go:embed testdata/doc.txt
	docFile []byte
)

var (
	testDataGolden = readFile(dataFile)
	testDocGolden  = readFile(docFile)
)

func readFile(raw []byte) []string {
	lines := bytes.Split(raw, []byte{'\n'})
	res := make([]string, len(lines))

	for i, line := range lines {
		res[i] = string(line)
	}
	return res
}
