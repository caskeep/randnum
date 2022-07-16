package randnum

import "fmt"

// RundTimeWeightRand Return the weighted readom result in runtime data
// data is like weight1,choice1,weight2,choice2,etc...
func RunTimeWeightRand(rand int, data []uint32) (uint32, error) {
	var err error
	if len(data)%2 != 0 {
		err = fmt.Errorf("len(data) mod 2 != 0")
		data = data[0 : len(data)-1]
	}
	if rand < 0 {
		rand = -rand
	}
	sum := uint64(0)
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			sum += uint64(data[i])
		}
	}
	realRand := uint32(uint64(rand) % sum)
	cur := uint32(0)
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			cur += data[i]
			if i+1 >= len(data) {
				panic("weight rand internal err")
			}
		} else {
			if realRand < cur {
				return data[i], err
			}
		}
	}
	return 0, err
}

// TODO pre constructed weight rand with O(1) time complecity algorithm
type WeightRandPool struct {
}

func (p *WeightRandPool) Build(data []uint32) error {

	return nil
}

func (p *WeightRandPool) DoRand(randA, randB int) (chosen int, err error) {
	return 0, nil
}
