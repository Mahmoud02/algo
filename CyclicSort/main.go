package main

import (
	"math"
)

/*
The Cyclic Sort pattern iterates over the array one number at a time,
and if the current number you are iterating is not at the correct index,
you swap it with the number at its correct index.
How do I identify this pattern?

They will be problems involving a sorted array with numbers in a given range
If the problem asks you to find the missing/duplicate/smallest number in an sorted/rotated array
Problems featuring cyclic sort pattern:
1-Find the Missing Number (easy)
2-Find the Smallest Missing Positive Number (medium)
*/
func main() {
	//println(missingNumber([]int{0, 1}))
	//println(missingNumberCyclicSort([]int{0, 1}))
	//missingNumberCyclicSort([]int{3, 0, 1})
	//findDuplicatesCyclicSort([]int{4, 3, 2, 7, 8, 2, 3, 1})
	//missingNumbersCyclicSort([]int{1, 1})
	//mismatchCyclicSort([]int{1, 1})
	println(findDuplicateNumberCyclicSort([]int{3, 1, 3, 4, 2}))

}

/*
Given an array nums containing n distinct numbers in the range [0, n],
return the only number in the range that is missing from the array.
link:https://leetcode.com/problems/missing-number/
*/
func missingNumber(nums []int) int {
	arrayLength := len(nums)
	missingNumber := math.MinInt
	for i := 0; i <= arrayLength; i++ {
		isMissing := true
		for _, item := range nums {
			if i == item {
				isMissing = false
			}
		}
		if isMissing {
			return i
		}
	}
	return missingNumber
}

//should no duplicates
func missingNumberCyclicSort(nums []int) int {

	currentIndex := 0
	arraySize := len(nums)
	for currentIndex < arraySize {
		if nums[currentIndex] < arraySize && currentIndex != nums[currentIndex] {
			currentValue := nums[currentIndex]
			tempValue := nums[currentValue]
			nums[currentIndex] = tempValue
			nums[currentValue] = currentValue
		} else {
			currentIndex++
		}
	}

	for index, item := range nums {
		if index != item {
			return index
		}
	}
	return arraySize
}

/*
 Find All Duplicates in an Array
Given an integer array nums of length n where all the integers of nums are in the range [1, n]
and each integer appears once or twice, return an array of all the integers that appears twice
link:https://leetcode.com/problems/find-all-duplicates-in-an-array/
*/
func findDuplicatesCyclicSort(nums []int) []int {
	currentIndex := 0
	arraySize := len(nums)
	for currentIndex < arraySize {
		// get right index for the item
		rightIndexForItem := nums[currentIndex] - 1
		if nums[currentIndex] != nums[rightIndexForItem] {
			//1 1 2
			tempValue := nums[currentIndex]
			nums[currentIndex] = nums[rightIndexForItem]
			nums[rightIndexForItem] = tempValue
		} else {
			currentIndex++
		}
	}
	var duplicateValues []int = make([]int, 0)
	for index, item := range nums {
		if index+1 != item {
			duplicateValues = append(duplicateValues, item)
		}
	}

	return duplicateValues
}

/*
 Find All Numbers Disappeared in an Array
Given an array nums of n integers where nums[i] is in the range [1, n],
return an array of all the integers in the range [1, n] that do not appear in nums.
link:https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
*/
func missingNumbersCyclicSort(nums []int) []int {

	currentIndex := 0
	arraySize := len(nums)
	for currentIndex < arraySize {
		// get right index for the item
		rightIndexForItem := nums[currentIndex] - 1
		if nums[currentIndex] != nums[rightIndexForItem] {
			//1 1 2
			tempValue := nums[currentIndex]
			nums[currentIndex] = nums[rightIndexForItem]
			nums[rightIndexForItem] = tempValue
		} else {
			currentIndex++
		}
	}
	var duplicateValues []int = make([]int, 0)
	for index, item := range nums {
		if index+1 != item {
			duplicateValues = append(duplicateValues, index+1)
		}
	}
	return duplicateValues
}

/*
645. Set Mismatch
You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately,
due to some error, one of the numbers in s got duplicated to another number in the set,
which results in repetition of one number and loss of another number.

link:https://leetcode.com/problems/set-mismatch/
*/
func mismatchCyclicSort(nums []int) []int {

	currentIndex := 0
	arraySize := len(nums)
	for currentIndex < arraySize {
		// get right index for the item
		rightIndexForItem := nums[currentIndex] - 1
		if nums[currentIndex] != nums[rightIndexForItem] {
			//1 1 2
			tempValue := nums[currentIndex]
			nums[currentIndex] = nums[rightIndexForItem]
			nums[rightIndexForItem] = tempValue
		} else {
			currentIndex++
		}
	}
	var missValues []int = make([]int, 0)
	for index, item := range nums {
		if index+1 != item {
			missValues = append(missValues, index+1)
			missValues = append(missValues, item)
		}
	}
	for _, item := range missValues {
		println(item)
	}
	return missValues
}

/*
 Find the Duplicate Number
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.
link:https://leetcode.com/problems/find-the-duplicate-number/
*/

func findDuplicateNumberCyclicSort(nums []int) int {

	currentIndex := 0
	arraySize := len(nums)
	for currentIndex < arraySize {
		// get right index for the item
		rightIndexForItem := nums[currentIndex] - 1
		if nums[currentIndex] != nums[rightIndexForItem] {
			//1 1 2
			tempValue := nums[currentIndex]
			nums[currentIndex] = nums[rightIndexForItem]
			nums[rightIndexForItem] = tempValue
		} else {
			currentIndex++
		}
	}
	for index, item := range nums {
		if index+1 != item {
			return item
		}
	}

	return 0
}
