package natsort_test

import (
	"fmt"
	"sort"

	"github.com/cristalhq/natsort"
)

func ExampleSort() {
	files := []string{"img12.png", "img10.png", "img2.png", "img1.png"}

	fmt.Println("Lexicographically:")

	sort.Strings(files)
	for _, f := range files {
		fmt.Println(f)
	}

	fmt.Println("\nNaturally:")

	natsort.Sort(files)
	for _, f := range files {
		fmt.Println(f)
	}

	// Output:
	// Lexicographically:
	// img1.png
	// img10.png
	// img12.png
	// img2.png
	//
	// Naturally:
	// img1.png
	// img2.png
	// img10.png
	// img12.png
}

func ExampleIsSorted() {
	files := []string{"img12.png", "img10.png", "img2.png", "img1.png"}

	if natsort.IsSorted(files) {
		fmt.Println("impossible!")
	}

	natsort.Sort(files)

	if !natsort.IsSorted(files) {
		fmt.Println("impossible!")
	}

	// Output:
}
