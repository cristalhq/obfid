package obfid

import "fmt"

func ExampleGenerator() {
	prime := uint64(32452867)
	random := uint64(123456)
	bits := 30

	generator, err := NewGenerator(prime, random, bits)
	if err != nil {
		panic(err)
	}

	for num := uint64(0); num < 10; num++ {
		enc := generator.Encode(num)
		dec := generator.Decode(enc)

		fmt.Printf("%d => %d => %d\n", num, enc, dec)

		if num != dec {
			fmt.Printf("oops must be equal: %d != %d\n", num, dec)
		}
	}

	// Output:
	// 0 => 123456 => 0
	// 1 => 32428867 => 1
	// 2 => 64979014 => 2
	// 3 => 97284425 => 3
	// 4 => 129836620 => 4
	// 5 => 162142031 => 5
	// 6 => 194692178 => 6
	// 7 => 227259733 => 7
	// 8 => 259549784 => 8
	// 9 => 292117339 => 9
}
