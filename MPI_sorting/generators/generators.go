package generators

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(9999) - rand.Intn(9998)
	}
	return slice
}

func GenerateSlicePartlySorted(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		tmp := rand.Intn(9999)
		if i%2 == 0 {
			slice[i] = i
		} else {
			for tmp < i {
				tmp = rand.Intn(9999)
			}
			slice[i] = tmp
		}
	}
	return slice
}

func GenerateInverseSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	j := 0
	for i := size; i > 0; i-- {
		slice[j] = i
		j++
	}
	return slice
}

func GenerateInversePartlysortedSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	j := 0
	for i := size; i > 0; i-- {
		tmp := rand.Intn(9999)
		if j%2 == 0 {
			slice[j] = i
		} else {
			for tmp > i {
				tmp = rand.Intn(9999)
			}
			slice[j] = tmp
		}
		j++
	}
	return slice
}
