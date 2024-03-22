package suanfa

import (
	"fmt"
	"testing"
)

type NumArray struct {
	slice    []int
	indexNum []int
}

func Constructor(nums []int) NumArray {
	indexNum1 := make([]int, len(nums)+1)
	for i, num := range nums {
		indexNum1[i+1] = indexNum1[i] + num
	}
	return NumArray{
		slice:    nums,
		indexNum: indexNum1,
	}
}
func (this *NumArray) SumRange(left int, right int) int {
	return this.indexNum[right+1] - this.indexNum[left]
}
func TestSuanfa101(t *testing.T) {
	fmt.Println("前缀树求数组任意区间和")
	s1 := []int{10, 20, 30, 40, 50}
	s2 := s1[2:4]
	fmt.Println(s2)
}
