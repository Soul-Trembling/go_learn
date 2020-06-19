package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	val  int
	next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//
//}

func insert(head *ListNode, i int) bool {
	p := head
	for nil != p {
		p = p.next
	}
	s := &ListNode{val: i}
	s.next = p.next
	p.next = s
	return true
}

func linkListInit(nums []int) *ListNode {
	if !(len(nums) > 0) {
		return nil
	}
	var node ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		if node == (ListNode{}) {
			if ok := insert(&node, nums[i]); ok {
				fmt.Println("xxxxx", node, &node)
			} else {
				break
			}
		} else {
			tempNode := ListNode{val: nums[i], next: &node}
			node = tempNode
			fmt.Println("sssss", node, &tempNode)
		}
	}
	return &node
}

func linkListPrint(nums *ListNode) {
	next := nums.next
	//for nil != next {
	//	println(next.val)
	//	next = next.next
	//}
	fmt.Println(next.val)
	fmt.Println(next.next.val)
}

func main() {
	first := linkListInit([]int{2, 4, 3})
	second := linkListInit([]int{5, 6, 4})

	fmt.Println(*first)
	linkListPrint(first)

	fmt.Println(*second)
	//linkListPrint(second)
}
