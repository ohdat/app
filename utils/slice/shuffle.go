package utils

import (
	"math/rand"
	"time"
)

func Shuffle[T int | int64 | int32 | string](nums []T) []T {
	var randomly = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(nums) - 1; i > 0; i-- {
		j := randomly.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
