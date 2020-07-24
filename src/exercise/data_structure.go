package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func baseStructure() {
	// 基础数据结构使用
	fmt.Println("stack")
	// 创建栈
	stack := make([]int, 0)
	// push压入
	fmt.Printf("%v", stack)
	stack = append(stack, 10)
	fmt.Printf("%v", stack)
	// pop弹出
	v := stack[len(stack)-1]
	fmt.Printf("%v", v)
	stack = stack[:len(stack)-1]
	fmt.Printf("%v", stack)
	// 检查栈空
	// len(stack)==0
	fmt.Println("")
	fmt.Println("queue")
	// 创建队列
	queue := make([]int, 0)
	fmt.Printf("%v", queue)
	// enqueue入队
	queue = append(queue, 10)
	fmt.Printf("%v", queue)
	// dequeue出队
	q := queue[0]
	fmt.Printf("%v", q)
	queue = queue[1:]
	fmt.Printf("%v", queue)
	// 长度0为空
	// len(queue)==0
	/*
		注意点
		参数传递，只能修改，不能新增或者删除原始数据
		默认 s=s[0:len(s)]，取下限不取上限，数学表示为：[)
	*/

	fmt.Println("")
	fmt.Println("map")
	// 创建
	m := make(map[string]int)
	fmt.Println(m)
	// 设置kv
	m["hello"] = 1
	fmt.Println(m)
	// 删除k
	// delete(m, "hello")
	// fmt.Println(m)
	// 遍历
	for k, v := range m {
		println(k, v)
	}
	/* 	注意点
	// map 键需要可比较，不能为 slice、map、function
	// map 值都有默认值，可以直接操作默认值，如：m[age]++ 值由 0 变为 1
	比较两个 map 需要遍历，其中的 kv 是否相同，因为有默认值关系，所以需要检查 val 和 ok 两个值 */

	fmt.Println("sort")
	// int排序
	sort.Ints([]int{})
	// 字符串排序
	sort.Strings([]string{})
	// 自定义排序
	// s := make([]int, 0)
	// s = append(s, 1)
	s := []int{1, 2, 6, 7, 5}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	fmt.Println(s)

	fmt.Println("math")
	// int32 最大最小值
	fmt.Println(math.MaxInt32) // 实际值：1<<31-1
	fmt.Println(math.MinInt32) // 实际值：-1<<31
	// int64 最大最小值（int默认是int64）
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

	fmt.Println("copy")
	// 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
	a := []int{1, 2, 6, 7, 5}
	i := 1
	copy(a[i:], a[i+1:])
	fmt.Println(a)
	a = a[:len(a)-1]
	fmt.Println(a)
	// make创建长度，则通过索引赋值
	n, x := 10, 99
	b := make([]int, n)
	b[n-1] = x
	fmt.Println(b)
	// make长度为0，则通过append()赋值
	c := make([]int, 0)
	c = append(a, x)
	fmt.Println(c)

	fmt.Println("类型转换")
	// 类型转换
	// byte转数字
	ss := "12345"                          // s[0] 类型是byte
	num := int(ss[0] - '0')                // 1
	str := string(ss[0])                   // "1"
	bb := byte(num + '0')                  // '1'
	fmt.Printf("%d %s %c\n", num, str, bb) // 111

	// 字符串转数字
	tNum, _ := strconv.Atoi(str)
	fmt.Println(tNum)
	tStr := strconv.Itoa(num)
	fmt.Println(tStr)
}

func strStr(haystack string, needle string) int {
	/*
		给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串
		中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回 -1。

		需要注意点
			循环时，i 不需要到 len-1
			如果找到目标字符串，len(needle)==j
	*/
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i不需要到len-1
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1
}

func subsets(nums []int) [][]int {
	/*
		给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

		思路：这是一个典型的应用回溯法的题目，简单来说就是穷尽所有可能性
		通过不停的选择，撤销选择，来穷尽所有可能性，最后将满足条件的结果返回

		面试注意点
			我们大多数时候，刷算法题可能都是为了准备面试，所以面试的时候需要注意一些点
			快速定位到题目的知识点，找到知识点的通用模板，可能需要根据题目特殊情况做特殊处理。
			先去朝一个解决问题的方向！先抛出可行解，而不是最优解！先解决，再优化！
			代码的风格要统一，熟悉各类语言的代码规范。
			命名尽量简洁明了，尽量不用数字命名如：i1、node1、a1、b2
			常见错误总结
			访问下标时，不能访问越界
			空值 nil 问题 run time error
	*/
	// 保存最终结果
	result := make([][]int, 0)
	// 保存中间结果
	list := make([]int, 0)
	subsetsBacktrack(nums, 0, list, &result)
	return result
}

func subsetsBacktrack(nums []int, pos int, list []int, result *[][]int) {
	/*
		nums 给定的集合
		pos 下次添加到集合中的元素位置索引
		list 临时结果集合(每次需要复制保存)
		result 最终结果
	*/
	// 把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	copy(ans, list)
	if len(ans) > 0 {
		// 子列表中没有元素 无意义
		*result = append(*result, ans)
	}

	// 选择、处理结果、再撤销选择
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		subsetsBacktrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

// TreeNode 二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	二叉树遍历
	前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
	中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
	后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点
*/
func preorderTraversalRecursive(root *TreeNode) {
	/*
		前序递归
	*/
	if root == nil {
		return
	}
	// 先访问根再访问左右
	fmt.Println(root.Val)
	preorderTraversalRecursive(root.Left)
	preorderTraversalRecursive(root.Right)
}

func preorderTraversalUnRecursive(root *TreeNode) []int {
	// 前序非递归遍历
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 前序遍历，所以先保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

func inorderTraversalUnRecursive(root *TreeNode) []int {
	/*
		中序非递归遍历
		思路：通过stack 保存已经访问的元素，用于原路返回
	*/
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 一直向左
		}
		// 弹出
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = val.Right
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	/*
		后序非递归
		通过lastVisit标识右子节点是否已经弹出
		核心就是：根节点必须在右节点弹出之后，再弹出
	*/
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

func traversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	// V1：深度遍历，结果指针作为参数传入到函数内部
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// func (node *TreeNode) treeNodeNew(index int, data []int) *TreeNode {
// 	// 二叉树节点创建
// 	if node == nil {
// 		return nil
// 	}
// 	bin := &TreeNode{data[index], nil, nil}
// 	// 设置完全二叉树左节点  其特征是深度 *2+1为左节点  +2为右节点
// 	if index < len(data) && 2*index+1 < len(data) {
// 		fmt.Println()
// 		bin.left = treeNodeNew(index*2+1, data)
// 	}
// 	if i < len(data) && 2*index+2 < len(data) {
// 		bin.right = treeNodeNew(index*2+2, data)
// 	}
// 	return bin
// }

func treeNodeTraverse(t *TreeNode) {
	if t == nil {
		return
	}
	treeNodeTraverse(t.Left)
	fmt.Println(t.Val, " ")
	treeNodeTraverse(t.Right)
}

func treeNodeCreate(n int) *TreeNode {
	var t *TreeNode
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = treeNodeInsert(t, temp)
	}
	return t
}

func treeNodeInsert(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return &TreeNode{v, nil, nil}
	}
	if v == t.Val {
		return t
	}
	if v < t.Val {
		t.Left = treeNodeInsert(t.Left, v)
		return t
	}
	t.Right = treeNodeInsert(t.Right, v)
	return t
}

func main() {
	baseStructure()

	strStr("yiuiihugu123456xneiiss", "123456")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(subsets(nums))

	fmt.Println("treeNode")
	tree := treeNodeCreate(10)
	// node = node.treeNodeNew(0, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
	treeNodeTraverse(tree)
	fmt.Println(tree.Val)
}
