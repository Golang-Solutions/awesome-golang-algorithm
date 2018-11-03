# [169. Majority Element][title]

## Description

Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.
**Example 1:**

```
Input: [3,2,3]
Output: 3
```

**Example 2:**

```
Input: [2,2,1,1,1,2,2]
Output: 2
```

**Tags:** Math, String

## 题意
>在一个数组除寻找出现次数超过一半的数

## 题解

### 思路1
> 暴力2层循环

```go
func majorityElement(nums []int) int {
	major, count := nums[0], 0

	for _, v1 := range nums {
		tmpCount := 0
		for _, v2 := range nums {
			if v2 == v1 {
				tmpCount++
			}
		}
		if tmpCount > count {
			count = tmpCount
			major = v1
		}
	}

	return major
}
```

### 思路2
> Map 计数

```go
func majorityElement(nums []int) int {
	major, count := nums[0], 0
	majorMap := make(map[int]int)

	for _, v := range nums {
		majorMap[v]++
	}

	for i, v := range majorMap {
		if v > count {
			major = i
			count = v
		}
	}

	return major
}
```

### 思路3
> 先排序再循环判断每个数出现的次数

```go
func majorityElement4(nums []int) int {
	//	将原数组排序
	nums = InsertSort(nums)

	major, count := nums[0], 0

	for _, v := range nums {
		if count == 0 {
			count++
			major = v
		} else if major == v {
			count++
		} else {
			count--
		}
	}

	return major
}

// 插入排序(Insert Sort)
func InsertSort(arr []int) []int {
	var i, j, tmp int
	for i = 1; i < len(arr); i++ {
		tmp = arr[i]
		for j = i; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}
```

### 思路4
> 利用分治分方法，借鉴了LeetCode某个大佬的分治解法，设计的非常精妙。


```go
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)

	if n%2 == 1 {
		count := 1

		for i := 0; i < n-1; i++ {
			if nums[n-1] == nums[i] {
				count++
			}
		}
		if count > n/2 {
			return nums[n-1]
		}
	}
	numsTmp := []int{}

	for i := 0; i < n/2; i++ {
		if nums[2*i] == nums[2*i+1] {
			numsTmp = append(numsTmp, nums[2*i])
		}
	}

	return majorityElement(numsTmp)
}
```


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: hthttps://leetcode.com/problems/majority-element/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
