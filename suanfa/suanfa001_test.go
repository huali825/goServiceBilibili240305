package suanfa

import (
	"fmt"
	"testing"
)

func Test_suanfa001(t *testing.T) {
	fmt.Println("这里是算法 2024年4月24日18:55:15 练习 ")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode算法题2  两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，
// 并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// addTwoNumbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	total, next1 := 0, 0
	result := &ListNode{}
	cur := result
	for l1 != nil && l2 != nil {
		total = l1.Val + l2.Val + next1
		cur.Next = &ListNode{Val: total % 10}
		next1 = total / 10
		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
	}
	for l2 != nil {
		total = l2.Val + next1
		cur.Next = &ListNode{Val: total % 10}
		next1 = total / 10
		l2 = l2.Next
		cur = cur.Next
	}
	for l1 != nil {
		total = l1.Val + next1
		cur.Next = &ListNode{Val: total % 10}
		next1 = total / 10
		l1 = l1.Next
		cur = cur.Next
	}
	if next1 != 0 {
		cur.Next = &ListNode{Val: next1}
	}
	return result.Next
}

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			//  如果 栈里面没东西了 || 栈顶的 符号 != 从 map 中取出来的 对应的符号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// leetcode算法题21  合并两个有序链表
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			cur.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}

	return result.Next
}
