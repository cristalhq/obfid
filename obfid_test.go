package obfid

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	generator, err := NewGenerator(32452867, 123, 30)
	failIfErr(t, err)

	for num := uint64(0); num < 30; num++ {
		enc := generator.Encode(num)
		dec := generator.Decode(enc)

		t.Logf("%d => %d => %d", num, enc, dec)

		mustEqual(t, num, dec)
	}
}

func TestGenerator2(t *testing.T) {
	generator, err := NewGenerator(32452867, 123, 30)
	failIfErr(t, err)

	for num := uint64(0); num < 30; num++ {
		dec := generator.Decode(num)
		enc := generator.Encode(dec)

		t.Logf("%d => %d => %d", num, dec, enc)

		mustEqual(t, num, enc)
	}
}

func TestGeneratorSmall(t *testing.T) {
	generator, err := NewGenerator(32452867, 123, 5)
	failIfErr(t, err)

	res := map[uint64]struct{}{}

	for num := uint64(0); num < 100; num++ {
		enc := generator.Encode(num)
		dec := generator.Decode(enc)

		t.Logf("%d => %d => %d", num, enc, dec)

		mustEqual(t, num%32, dec)
		res[dec] = struct{}{}
	}

	for i := uint64(0); i < 32; i++ {
		delete(res, i)
	}

	mustEqual(t, len(res), 0)
}

func Test_inverse(t *testing.T) {
	inv, err := inverse(32452867, 1<<31-1)
	failIfErr(t, err)
	mustEqual(t, inv, uint64(23970219))
}

func BenchmarkEncode(b *testing.B) {
	generator, err := NewGenerator(32452867, 123, 30)
	failIfErr(b, err)

	var count uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count += generator.Encode(uint64(i))
	}
	sink(count)
}

func BenchmarkDecode(b *testing.B) {
	generator, err := NewGenerator(32452867, 123, 30)
	failIfErr(b, err)

	var count uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count += generator.Decode(uint64(i))
	}
	sink(count)
}

func Benchmark_inverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inv, err := inverse(32452867, 1<<31-1)
		failIfErr(b, err)
		mustEqual(b, inv, uint64(23970219))
	}
}

func failIfErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func mustEqual(t testing.TB, have, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", have, want)
	}
}

func sink(v uint64) {
	if rand.Float32() > 1 {
		panic(v)
	}
}
