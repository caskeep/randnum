package randnum

import "fmt"

// RunTimeWeightRand Return the weighted random result in runtime data
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

type singleSelect struct {
	edge, prev, next uint32
}

type WeightRandPool struct {
	x, y uint32
	data []singleSelect
}

func (p *WeightRandPool) Build(data []uint32) error {
	if len(data)%2 != 0 {
		return fmt.Errorf("len=%d err", len(data))
	}
	p.x = uint32(len(data) / 2)
	sum := uint32(0)
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			sum += data[i]
		}
	}
	p.y = sum
	tmp := make([]uint32, len(data))
	copy(tmp, data)
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

func (p *WeightRandPool) DoRand(randA, randB int) (chosen uint32, err error) {
	randA = randA % int(p.x)
	randB = randB % int(p.y)
	if randB < int(p.data[randA].edge) {
		return p.data[randA].prev, nil
	} else {
		return p.data[randA].next, nil
	}
}
