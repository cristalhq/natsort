package natsort

import (
	"bytes"
	_ "embed"
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	data := bench(b, smallList, func(list []string) {
		Sort(list)
	})
	checkResult(b, data)
}

func BenchmarkSortString(b *testing.B) {
	data := bench(b, smallList, func(list []string) {
		sort.Sort(Slice(list))
	})
	checkResult(b, data)
}

func BenchmarkSortStdlib(b *testing.B) {
	bench(b, smallList, func(list []string) {
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

func checkResult(b *testing.B, list []string) {
	b.Helper()

	ok := sort.SliceIsSorted(list, func(i, j int) bool {
		return Less(list[i], list[j])
	})
	if !ok {
		b.Errorf("not sorted %+v", list)
	}
}

//go:embed testdata/small.txt
var smallFile []byte

var smallList []string

func init() {
	smallList = readFile(smallFile)
}

func readFile(raw []byte) []string {
	lines := bytes.Split(raw, []byte{'\n'})
	res := make([]string, len(lines))

	for i, line := range lines {
		res[i] = string(line)
	}
	return res
}
