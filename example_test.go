package obfid

import "fmt"

func ExampleGenerator() {
	generator, err := NewGenerator(32452867, 123, 30)
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
	// 0 => 123 => 0
	// 1 => 32452984 => 1
	// 2 => 64905853 => 2
	// 3 => 97358706 => 3
	// 4 => 129811575 => 4
	// 5 => 162264436 => 5
	// 6 => 194717289 => 6
	// 7 => 227170158 => 7
	// 8 => 259623011 => 8
	// 9 => 292075872 => 9
}
