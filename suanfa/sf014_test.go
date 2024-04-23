package suanfa

import (
	"fmt"
	"testing"
)

func TestSuanfa014(t *testing.T) {
	fmt.Println("这里是算法014")
}

// 两数之和
//func twoSum(nums []int, target int) (a []int) {
//	map1 := make(map[int]int, len(nums))
//	for k, v := range nums {
//		map1[v] = k
//	}
//
//	for k, v := range nums {
//		v1 := target - v
//		if k1, ok := map1[v1]; ok == true && k != k1 {
//			a = []int{k, k1}
//			return
//		}
//	}
//	return
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (l3 *ListNode) {
	return nil
}
