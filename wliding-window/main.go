package main

import (
	"math"
)

/*
Sliding window based problems are easy to identify if we practice a few.
Problems such as finding the longest substring or shortest substring with some constraints
are mostly based on sliding window.

*/
func main() {

	//println(minSubArrayLen([]int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}, 15))
	//println(maxSubArrayBrute([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3))
	//println(minSubArrayWindow([]int{0, 1, 2, 3, 4, 5}, 4))
	//println(numberOfSubstrings("aaabbbccc"))
	println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}

/*
Given an array of positive integers nums and a positive integer target,
return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr]
of which the sum is greater than or equal to target.
If there is no such subarray, return 0 instead.
*/
func minSubArrayLenBrute(arr []int, target int) int {
	finalLength := 0
	for index, _ := range arr {

		tempLength := 0
		currentIndex := index
		sumResult := 0
		for target > sumResult && currentIndex < len(arr) {
			sumResult += arr[currentIndex]
			currentIndex++
			tempLength++
		}
		if sumResult >= target {
			if finalLength == 0 {
				finalLength = tempLength
			}
			finalLength = int(math.Min(float64(finalLength), float64(tempLength)))
		}
	}

	return finalLength
}

//3
/*1 2 3 5*/
func minSubArrayWindow(arr []int, target int) int {
	minLength := math.MaxInt
	sumValues := 0
	//Start index
	startIndex := 0
	for index, _ := range arr {
		sumValues += arr[index]
		for sumValues >= target {
			minLength = int(math.Min(float64(minLength), float64(index-startIndex+1)))
			sumValues -= arr[startIndex]
			startIndex++
		}

	}
	if minLength == math.MaxInt {
		minLength = 0
	}
	return minLength
}

/**
Find max sum of subarray with fixed length
ex: [4,2,1,7,8,1,2,8,1,0] k=3
solve: [1,7,8] \ 16
**/
func maxSubArrayBrute(arr []int, size int) int {
	maxSum := 0
	for index, _ := range arr {
		localSum := 0
		//condition
		sizeLength := 0
		currentIndex := index
		for currentIndex < len(arr) && sizeLength < size {
			localSum += arr[currentIndex]
			sizeLength++
			currentIndex++
		}
		maxSum = int(math.Max(float64(maxSum), float64(localSum)))
	}
	return maxSum
}
func maxSubArrayWindow(arr []int, size int) int {
	maxSum := 0
	localSum := 0
	for index, _ := range arr {
		localSum += arr[index]
		if index >= size-1 {
			maxSum = int(math.Max(float64(localSum), float64(maxSum)))
			localSum -= arr[index-(size-1)]
		}
	}
	return maxSum
}

/**
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array
if you can flip at most k 0's.
Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
link:https://leetcode.com/problems/max-consecutive-ones-iii/
**/
func maxNumberOfOnes(arr []int, numbOfFlipTrials int) int {
	maxConsecutiveOnes := math.MinInt8
	starter := 0
	counterLimit := 0
	for index, _ := range arr {
		if arr[index] == 0 {
			counterLimit++
		}
		for counterLimit > numbOfFlipTrials && starter < len(arr) {
			if arr[starter] == 0 {
				counterLimit--
			}
			starter++
		}
		// 111 00 0 1
		maxConsecutiveOnes = int(math.Max(float64(maxConsecutiveOnes), float64(index-starter+1)))
	}
	return maxConsecutiveOnes
}

/**
Given a string s consisting only of characters a, b and c.
Return the number of substrings containing at least one occurrence of all these characters a, b and c.
Input: s = "abcabc"
Output: 10
Explanation: The substrings containing at least one occurrence of the characters a, b and c
 are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again).
**/
func numberOfSubstrings(s string) int {
	stringLength := len(s)
	numOfSubStrings := 0
	indexStarter := 0
	localCounts := 0
	charMap := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
	}
	for _, item := range s {
		charMap[string(item)]++
		for indexStarter < stringLength && charMap["a"] > 0 && charMap["b"] > 0 && charMap["c"] > 0 {
			localCounts++
			currentItem := s[indexStarter]
			charMap[string(currentItem)]--
			indexStarter++
		}
		numOfSubStrings += localCounts
	}
	return numOfSubStrings
}

/*
	link:https://leetcode.com/problems/fruit-into-baskets/
*/
func totalFruit(fruits []int) int {
	totalTreesPicks := math.MinInt
	starterIndex := 0
	treeTypeMap := make(map[int]int)
	for index, item := range fruits {
		treeTypeMap[item]++
		for len(treeTypeMap) > 2 && starterIndex < len(fruits) {
			treeTypeMap[fruits[starterIndex]]--
			if treeTypeMap[fruits[starterIndex]] == 0 {
				delete(treeTypeMap, fruits[starterIndex])
			}
			starterIndex++
		}

		totalTreesPicks = int(math.Max(float64(totalTreesPicks), float64(index-starterIndex+1)))
	}

	return totalTreesPicks
}
