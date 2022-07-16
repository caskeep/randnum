package randnum

import "testing"

func TestRunTimeWeightRandSucc(t *testing.T) {
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

func TestRunTimeWeightNegitaveRand(t *testing.T) {
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
