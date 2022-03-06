package main

import "math"

func main() {
	println(Solution("aaaa", []int{3, 4, 5, 6}))
}
func Solution(S string, C []int) int {
	length := len(S)
	gSum := 0
	gMax := 0
	ans := 0
	for i := 1; i < length; i++ {
		if i > 0 && S[i] != S[i-1] {
			ans += gSum - gMax
			gSum = 0
			gMax = 0
		}
		gSum += C[i]
		gMax = int(math.Max(float64(gMax), float64(C[i])))
	}
	ans += gSum - gMax
	// write your code in Go 1.4
	return ans
}
func Solution2(S string, C []int) int {
	length := len(S)
	gSum := 0
	gMax := 0
	ans := 0
	for i := 1; i < length; i++ {
		if i > 0 && S[i] != S[i-1] {
			ans += gSum - gMax
			gSum = 0
			gMax = 0
		}
		gSum += C[i]
		gMax = int(math.Max(float64(gMax), float64(C[i])))
	}
	ans += gSum - gMax
	// write your code in Go 1.4
	return ans
}
int result = 0, accu = 0, max_cost = 0
for (int i = 0; i < s.length(); ++i) {
if (i && s[i] != s[i - 1]) {
result += accu - max_cost;
accu = max_cost = 0;
}
accu += cost[i];
max_cost = max(max_cost, cost[i]);
}
result += accu - max_cost
return result
