package randnum

import (
	"math/rand"
	"testing"
)

const (
	randSource = int64(1232)
)

func TestRunTimeWeightRandSuccess(t *testing.T) {
	var err error
	var chosen uint32
	//  0~6   => 1
	//  7~11  => 10
	// 12~24  => 7
	data := []uint32{7, 1, 5, 10, 13, 7}
	r1 := 0
	r2 := 6
	r3 := 7
	r4 := 11
	r5 := 12
	r6 := 24
	r7 := 25
	r8 := 49

	chosen, err = RunTimeWeightRand(r1, data)
	if err != nil {
		t.Error(err)
	}
	if chosen != 1 {
		t.Error(r1)
	}
	chosen, err = RunTimeWeightRand(r2, data)
	if err != nil {
		t.Error(err)
	}
	if chosen != 1 {
		t.Error(r2)
	}
	chosen, err = RunTimeWeightRand(r3, data)
	if err != nil {
		t.Error(err)
	}
	if chosen != 10 {
		t.Error(r3)
	}
	chosen, err = RunTimeWeightRand(r4, data)
	if err != nil {
		t.Error(err)
	}
	if chosen != 10 {
		t.Error(r4)
	}
	chosen, err = RunTimeWeightRand(r5, data)
	if err != nil {
		t.Error(err)
	}
	if chosen != 7 {
		t.Error(r5)
	}
	chosen, err = RunTimeWeightRand(r6, data)
	if err != nil {
		t.Error(err)
	}
	if chosen != 7 {
		t.Error(r6)
	}
	chosen, err = RunTimeWeightRand(r7, data)
	if err != nil {
		t.Error(err)
	}
	if chosen != 1 {
		t.Error(r7)
	}
	chosen, err = RunTimeWeightRand(r8, data)
	if err != nil {
		t.Error(err)
	}
	if chosen != 7 {
		t.Error(r8)
	}
}

func TestRunTimeWeightRandUnmatch(t *testing.T) {
	var err error
	var chosen uint32
	//  0~6   => 1
	//  7~11  => 10
	// 12~24  => 7
	data := []uint32{7, 1, 5, 10, 13, 7, 8}
	r1 := 11
	r2 := 25
	chosen, err = RunTimeWeightRand(r1, data)
	if err == nil {
		t.Error(err)
	}
	if chosen != 10 {
		t.Error(r1)
	}
	chosen, err = RunTimeWeightRand(r2, data)
	if err == nil {
		t.Error(err)
	}
	if chosen != 1 {
		t.Error(r2, chosen)
	}
}

func TestRunTimeWeightNegativeRand(t *testing.T) {
	var err error
	var chosen uint32
	//  0~6   => 1
	//  7~11  => 10
	// 12~24  => 7
	data := []uint32{7, 1, 5, 10, 13, 7}
	r1 := -11
	chosen, err = RunTimeWeightRand(r1, data)
	if err != nil {
		t.Error(err)
	}
	if chosen != 10 {
		t.Error(r1)
	}
}

func helperSingleSelectCompare(r, l singleSelect) bool {
	if r.edge != l.edge {
		return false
	}
	if r.prev != l.prev {
		return false
	}
	if r.next != l.next {
		return false
	}
	return true
}

func TestWeightedRandPoolSuccess(t *testing.T) {
	var err error
	var chosen uint32
	data := []uint32{7, 1, 5, 10, 13, 7}

	pool := WeightRandPool{}
	err = pool.Build(data)
	if err != nil {
		t.Error(err)
	}

	if pool.x != 3 {
		t.Fatal("x not 3")
	}
	if pool.y != 25 {
		t.Fatal("y not 25")
	}
	if len(pool.data) != 3 {
		t.Fatal("len(pool.data] not 3")
	}
	if !helperSingleSelectCompare(pool.data[0], singleSelect{
		edge: 15,
		prev: 10,
		next: 7,
	}) {
		t.Fatal("data[0] err")
	}
	if !helperSingleSelectCompare(pool.data[1], singleSelect{
		edge: 21,
		prev: 1,
		next: 7,
	}) {
		t.Fatal("data[0] err")
	}
	if !helperSingleSelectCompare(pool.data[2], singleSelect{
		edge: 25,
		prev: 7,
		next: 7,
	}) {
		t.Fatal("data[0] err")
	}

	chosen, err = pool.DoRand(0, 0)
	if err != nil {
		t.Error(err)
	}
	if chosen != 10 {
		t.Error(0, 0)
	}

	chosen, err = pool.DoRand(0, 14)
	if err != nil {
		t.Error(err)
	}
	if chosen != 10 {
		t.Error(0, 0)
	}

	chosen, err = pool.DoRand(0, 15)
	if err != nil {
		t.Error(err)
	}
	if chosen != 7 {
		t.Error(0, 0)
	}

	chosen, err = pool.DoRand(0, 24)
	if err != nil {
		t.Error(err)
	}
	if chosen != 7 {
		t.Error(0, 0)
	}

	chosen, err = pool.DoRand(1, 0)
	if err != nil {
		t.Error(err)
	}
	if chosen != 1 {
		t.Error(0, 0)
	}

	chosen, err = pool.DoRand(1, 20)
	if err != nil {
		t.Error(err)
	}
	if chosen != 1 {
		t.Error(0, 0)
	}

	chosen, err = pool.DoRand(1, 21)
	if err != nil {
		t.Error(err)
	}
	if chosen != 7 {
		t.Error(0, 0)
	}

	chosen, err = pool.DoRand(1, 24)
	if err != nil {
		t.Error(err)
	}
	if chosen != 7 {
		t.Error(0, 0)
	}

	chosen, err = pool.DoRand(2, 0)
	if err != nil {
		t.Error(err)
	}
	if chosen != 7 {
		t.Error(0, 0)
	}

	chosen, err = pool.DoRand(2, 24)
	if err != nil {
		t.Error(err)
	}
	if chosen != 7 {
		t.Error(0, 0)
	}
}

func TestWeightedRandPoolLenZero(t *testing.T) {
	var err error
	data := []uint32{}

	pool := WeightRandPool{}
	err = pool.Build(data)
	if err == nil {
		t.Error(err)
	}
}

func helperBenchmarkRunTime(b *testing.B, r *rand.Rand, data []uint32) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = RunTimeWeightRand(r.Int(), data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func helperBenchmarkStaticTime(b *testing.B, r *rand.Rand, data []uint32) {
	var err error
	pool := WeightRandPool{}
	err = pool.Build(data)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		_, err = pool.DoRand(r.Int(), r.Int())
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRunTimeWeightedRandomLen3(b *testing.B) {
	data := []uint32{7, 1, 5, 10, 13, 7}
	r := rand.New(rand.NewSource(randSource))
	helperBenchmarkRunTime(b, r, data)
}

func BenchmarkStaticTimeWeightedRandomLen3(b *testing.B) {
	data := []uint32{7, 1, 5, 10, 13, 7}
	r := rand.New(rand.NewSource(randSource))
	helperBenchmarkStaticTime(b, r, data)
}

func BenchmarkRunTimeWeightedRandomLen6(b *testing.B) {
	data := []uint32{7, 1, 5, 2, 13, 3, 78, 4, 5, 6, 13, 7}
	r := rand.New(rand.NewSource(randSource))
	helperBenchmarkRunTime(b, r, data)
}

func BenchmarkStaticTimeWeightedRandomLen6(b *testing.B) {
	data := []uint32{7, 1, 5, 2, 13, 3, 78, 4, 5, 6, 13, 7}
	r := rand.New(rand.NewSource(randSource))
	helperBenchmarkStaticTime(b, r, data)
}

func BenchmarkRunTimeWeightedRandomLen12(b *testing.B) {
	data := []uint32{7, 1, 5, 2, 13, 3, 78, 4, 5, 6, 13, 7,
		45, 8, 78, 9, 89, 10, 32, 11, 78, 12}
	r := rand.New(rand.NewSource(randSource))
	helperBenchmarkRunTime(b, r, data)
}

func BenchmarkStaticTimeWeightedRandomLen12(b *testing.B) {
	data := []uint32{7, 1, 5, 2, 13, 3, 78, 4, 5, 6, 13, 7,
		45, 8, 78, 9, 89, 10, 32, 11, 78, 12}
	r := rand.New(rand.NewSource(randSource))
	helperBenchmarkStaticTime(b, r, data)
}

func BenchmarkRunTimeWeightedRandomLen24(b *testing.B) {
	data := []uint32{7, 1, 5, 2, 13, 3, 78, 4, 5, 6, 13, 7,
		45, 8, 78, 9, 89, 10, 32, 11, 78, 12,
		45, 13, 78, 14, 89, 15, 32, 16, 78, 17,
		35, 18, 56, 19, 67, 20, 78, 21, 53, 22, 56, 23, 56, 24}
	r := rand.New(rand.NewSource(randSource))
	helperBenchmarkRunTime(b, r, data)
}

func BenchmarkStaticTimeWeightedRandomLen24(b *testing.B) {
	data := []uint32{7, 1, 5, 2, 13, 3, 78, 4, 5, 6, 13, 7,
		45, 8, 78, 9, 89, 10, 32, 11, 78, 12,
		45, 13, 78, 14, 89, 15, 32, 16, 78, 17,
		35, 18, 56, 19, 67, 20, 78, 21, 53, 22, 56, 23, 56, 24}
	r := rand.New(rand.NewSource(randSource))
	helperBenchmarkStaticTime(b, r, data)
}

// Benchmark results
// goos: linux
// goarch: amd64
// pkg: github.com/caskeep/randnum
// cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
// BenchmarkRunTimeWeightedRandomLen3-8            49886893                21.53 ns/op
// BenchmarkStaticTimeWeightedRandomLen3-8         72011685                14.29 ns/op
// BenchmarkRunTimeWeightedRandomLen6-8            47187850                25.12 ns/op
// BenchmarkStaticTimeWeightedRandomLen6-8         71503300                14.40 ns/op
// BenchmarkRunTimeWeightedRandomLen12-8           28984764                41.09 ns/op
// BenchmarkStaticTimeWeightedRandomLen12-8        83873653                14.01 ns/op
// BenchmarkRunTimeWeightedRandomLen24-8           19592414                60.89 ns/op
// BenchmarkStaticTimeWeightedRandomLen24-8        71411912                14.29 ns/op
// PASS
// ok      github.com/caskeep/randnum       12.856s
