```go
package main

import "fmt"

// generic akan bersifat lebih fleksibel
// sehingga kita bisa menggunakan tipe data apapun

type Numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func sumGenerics[K comparable, V Numbers](g map[K]V) V {
	var sum V
	for _, v := range g {
		sum += v
	}
	return sum
}

func main( ) {
	// This is a placeholder for the main function.
	// The actual code will be in the goroutine-test.md file.
	ints := map[string]int{"a": 1, "b": 2, "c": 3}
	floats := map[string]float64{"x": 1.1, "y": 2.2, "z": 3.3}
	// Example usage of sumGenerics
	fmt.Println("Sum of integers:", sumGenerics(ints))
	fmt.Printf("Generics of Sum : %v and %v\n", sumGenerics(ints), sumGenerics(floats))
}
```