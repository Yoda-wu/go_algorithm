package lc_daily

// 树状数组

func isIdealPermutation(nums []int) bool {
	n := len(nums)
	tr := make([]int, n+10)
	// lowbit函数
	lowbit := func(x int) int {
		return x & -x
	}
	// 修改树状数组值
	add := func(x int) {
		for i := x; i <= n; i += lowbit(x) {
			tr[i]++
		}

	}
	// 查询树状数组值——返回下标为x的前缀和
	query := func(x int) int {
		ans := 0
		for i := x; i > 0; i -= lowbit(x) {
			ans += tr[i]
		}
		return ans
	}

	add(nums[0] + 1)
	a := 0
	b := 0
	// 有一个局部倒置就肯定由一个全局倒置
	for i := 0; i < n; i++ {
		// 求全局倒置
		a += query(n) - query(nums[i]+1) // 对于每个nums[i] 其左边比它大的数字的个数，是以nums[i]为右端点的全局倒置数的数量
		if nums[i] == nums[i-1] {
			b += 1
		}
		add(nums[i] + 1)
	}

	return a == b

}
