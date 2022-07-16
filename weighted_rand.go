package randnum

import "fmt"

// RunTimeWeightRand Return the weighted random result in runtime data
// data is like weight1,choice1,weight2,choice2,etc...
func RunTimeWeightRand(rand int, input []uint32) (uint32, error) {
	var err error
	if len(input)%2 != 0 {
		err = fmt.Errorf("len(data) mod 2 != 0")
		input = input[0 : len(input)-1]
	}
	if rand < 0 {
		rand = -rand
	}
	sum := uint64(0)
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			sum += uint64(input[i])
		}
	}
	realRand := uint32(uint64(rand) % sum)
	cur := uint32(0)
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			cur += input[i]
			if i+1 >= len(input) {
				panic("weight rand internal err")
			}
		} else {
			if realRand < cur {
				return input[i], err
			}
		}
	}
	return 0, err
}

type singleSelect struct {
	edge, prev, next uint32
}

// WeightRandPool is pre-built weighted random pool
type WeightRandPool struct {
	x, y uint32
	data []singleSelect
}

// Build builds pre-calculated data structure,
// see: https://www.keithschwarz.com/darts-dice-coins/
func (p *WeightRandPool) Build(input []uint32) error {
	if len(input) == 0 {
		return fmt.Errorf("len==0")
	}
	if len(input)%2 != 0 {
		return fmt.Errorf("len=%d err", len(input))
	}
	p.x = uint32(len(input) / 2)
	sum := uint32(0)
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			sum += input[i]
		}
	}
	p.y = sum
	tmp := make([]uint32, len(input))
	copy(tmp, input)
	for i := 0; i < len(tmp); i++ {
		if i%2 == 0 {
			tmp[i] = tmp[i] * p.x
		}
	}
	min := -1
	minVal := sum
	max := -1
	maxVal := uint32(0)
	for {
		min = -1
		minVal = sum
		max = -1
		maxVal = uint32(0)
		for i := 0; i < len(tmp); i++ {
			if i%2 == 0 {
				cur := tmp[i]
				if cur == 0 {
					continue
				}
				if cur < minVal {
					min = i
					minVal = cur
				}
				if cur > maxVal {
					max = i
					maxVal = cur
				}
			}
		}
		if maxVal == 0 {
			break
		}
		if min == -1 {
			min = max
		}
		p.data = append(p.data, singleSelect{
			edge: tmp[min],
			prev: tmp[min+1],
			next: tmp[max+1],
		})
		tmp[max] -= sum - tmp[min]
		tmp[min] = 0
	}
	if len(p.data) != int(p.x) {
		panic("internal len err")
	}
	return nil
}

// DoRand returns chosen select value
func (p *WeightRandPool) DoRand(randA, randB int) (chosen uint32, err error) {
	randA = randA % int(p.x)
	randB = randB % int(p.y)
	if randB < int(p.data[randA].edge) {
		return p.data[randA].prev, nil
	} else {
		return p.data[randA].next, nil
	}
}
