package main

/*
The idea here is to iterate two different parts of the array simultaneously to get the answer faster.
There are primarily two ways of implementing the two-pointer technique
1. One pointer at each end
One pointer starts from beginning and other from the end and they proceed towards each other

*/
func main() {
	//println(pairExists([]int{1, 2, 3, 4, 8, 9, 10}, 7))
	//reverseArray([]int{1, 2, 3, 4, 8, 9, 10})
	//reverseArrayTwoPointers([]int{1, 2, 3, 4, 8, 9, 10})
	//println(NumOfUniqueValues([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//println(removeDuplicatesMoreThanTwo([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
	println(removeDuplicatesMoreThanTwo2([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))

}

/*In a sorted array, find if a pair exists with a given sum S*/
func pairExists(arr []int, target int) bool {
	forwardIndex := 0
	backwardIndex := len(arr) - 1
	for forwardIndex < backwardIndex {
		sumValue := arr[forwardIndex] + arr[backwardIndex]
		println(sumValue)
		if sumValue == target {
			return true
		} else if sumValue < target {
			forwardIndex += 1
		} else {
			backwardIndex -= 1
		}
	}
	return false
}

/*revers array o(n)*/
func reverseArray(arr []int) []int {
	var reversedArray = make([]int, len(arr))
	for index, item := range arr {
		currentIndex := len(arr) - index - 1
		reversedArray[currentIndex] = item
	}
	for _, item := range reversedArray {
		print(item)

	}
	return reversedArray
}

/*revers array o(n)*/
func reverseArrayTwoPointers(arr []int) []int {
	forwardIndex := 0
	backwardIndex := len(arr) - 1
	for forwardIndex < backwardIndex {
		tempValue := arr[forwardIndex]
		arr[forwardIndex] = arr[backwardIndex]
		arr[backwardIndex] = tempValue

		forwardIndex += 1
		backwardIndex -= 1
	}
	for _, item := range arr {
		print(item)

	}
	return arr
}

/*lazy pointer*/

//find number of UniqueElement in a sorted array(they array must be sorted to avoid recalculated)
// remove duplicates

func NumOfUniqueValues(arr []int) int {
	forwardIndex := 1
	lazyIndex := 0
	countOfUniqueness := 1
	insertIndex := 0
	for forwardIndex < len(arr) {
		if arr[lazyIndex] == arr[forwardIndex] {
			forwardIndex++
		} else {
			countOfUniqueness++
			insertIndex++
			arr[insertIndex] = arr[forwardIndex]
			lazyIndex = forwardIndex
		}
	}

	return countOfUniqueness
}
func removeDuplicatesMoreThanTwo(arr []int) int {
	forwardIndex := 1
	lazyIndex := 0
	countOfUniqueness := 1
	maxUnique := 0
	for forwardIndex < len(arr) {
		if arr[lazyIndex] == arr[forwardIndex] {
			maxUnique++
			if maxUnique < 2 {
				countOfUniqueness++
			}
			forwardIndex++
		} else {
			maxUnique = 0
			countOfUniqueness++
			lazyIndex = forwardIndex
			forwardIndex++
		}
	}
	return countOfUniqueness
}

// 0, 0, 1, 1, 1, 2, 2, 3, 4
// 0, 0, 1, 1, 2, 2, 3, 4, 4
func removeDuplicatesMoreThanTwo2(arr []int) int {
	forwardIndex := 2
	lazyIndex := 2
	for forwardIndex < len(arr) {
		if arr[forwardIndex] != arr[lazyIndex-2] {
			arr[lazyIndex] = arr[forwardIndex]
			lazyIndex++
			forwardIndex++
		} else {
			forwardIndex++
		}
	}
	return lazyIndex
}

//0,0,1,1,2,2,3,3,4
//0112233
//rotate the array

func Solution(A []int) int {

	for i := 1; i < 9; i++ {
		smallest := i
		for _, item := range A {
			if item != smallest && item < smallest && item > 0 {
				smallest = item
				return smallest
			}
		}
	}
	return 1
}
