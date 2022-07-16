package randnum

import "testing"

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
