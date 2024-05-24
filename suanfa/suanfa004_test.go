package suanfa

import (
	"fmt"
	"testing"
)

func TestSuanfa(t *testing.T) {
	nums := []int{2, 3, 3, 5, 2, 12, 4, 5, 6, 3}
	and := 17
	x := twoSum(nums, and)
	fmt.Println(x)
}
func twoSum1(nums []int, target int) (a []int) {
	prevNums := map[int]int{}
	for i, num := range nums {
		targetNum := target - num
		targetNumIndex, ok := prevNums[targetNum]
		if ok {
			return []int{targetNumIndex, i}
		} else {
			prevNums[num] = i
		}
	}
	return []int{}
}
