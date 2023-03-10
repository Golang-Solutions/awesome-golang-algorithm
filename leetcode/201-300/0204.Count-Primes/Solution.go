package Solution

func countPrimes_1(n int) int {
	if n <= 1 {
		return 0
	}
	var ans int
	for i := 2; i < n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			ans++
		}
	}
	return ans
}

// 整合简化版
func countPrimes_2(n int) int {
	var ans int
	prime := make([]bool, n)
	for i := 2; i < n; i++ {
		if prime[i] == false {
			ans++
			for j := 2; j*i < n; j++ {
				prime[j*i] = true
			}
		}
	}
	return ans
}

func countPrimes_3(num int) int {
	isPrime := make([]bool, num)
	//	初始化数组 true表示都是质数
	for i := 0; i < num; i++ {
		isPrime[i] = true
	}
	//	循环检查质数 i<sqrt(n)可以转换为i*i < n
	//	因为提劲的进行标记所以时间复杂可以到O(log N)
	for i := 2; i*i < num; i++ {
		if isPrime[i] == false {
			continue
		}
		for j := i * i; j < num; j = j + i {
			isPrime[j] = false
		}
	}
	ans := 0
	for i := 2; i < num; i++ {
		if isPrime[i] {
			ans++
		}
	}
	return ans
}
