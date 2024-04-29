package suanfa

import (
	"fmt"
	"sort"
	"testing"
)

func Test_suanfa002(t *testing.T) {
	fmt.Println("这里是算法 2024年4月25日22:13:55  练习 ")
	slice2 := []int{-5, -3, -6, -2, 0, 1, 3, 3, 7, 9}
	//slice2 := []int{-3, -2, -1, 0, 1, 2, 3}
	fmt.Println(countPairs(slice2, 0))
	//slice2 := []int{0, 0, 0, 0}
	//fmt.Println(slice1[2:8])
	//fmt.Println(threeSum(slice2))
}

func countPairs(nums []int, target int) (ans int) {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return
}

// 三数之和
func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i, val := range nums[:n-2] {
		if i > 0 && nums[i-1] == val {
			continue
		}
		if val+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if val+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			if val+nums[j]+nums[k] < 0 {
				j++
			} else if val+nums[j]+nums[k] > 0 {
				k--
			} else {
				ans = append(ans, []int{val, nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				} // 跳过重复数字
				j++
				for k > j && nums[k] == nums[k-1] {
					k--
				}
				k--
			}

		}
	}
	return
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	res := make([]int, 0)
	//var res2 []int
	for left < right {
		temp := numbers[left] + numbers[right]
		if temp == target {
			return append(res, left, right)
		}
		if temp > target {
			right--
		}
		if temp < target {
			left--
		}
	}
	return nil
}

func threeSum2(nums []int) (ans [][]int) {
	sort.Ints(nums)

	n := len(nums)
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] { // 跳过重复数字
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 { // 优化一
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 { // 优化二
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 跳过重复数字
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // 跳过重复数字
			}
		}
	}
	return
}
