package main

import (
	"ArrayInterViewQuestion/example"
	"fmt"
)

func main() {
	data := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	sum := example.FindSubArrayMaxSum(data)

	fmt.Println("Example 1 ", sum)

	// data2 := []int{-2, -3, 4, -1, -2, 1, 5, -3}

	sum = example.KadaneAlorith(data)

	fmt.Println("Example 2 ", sum)

	sum, subarray := example.KadaneAlorithSubArray(data)

	fmt.Println("Sub Array is ", sum, subarray)

	fmt.Println(example.IsPalindromNumber(121))

	fmt.Println(example.IsPalindromNumber(121122))

	// sum of sub array

	example.FindSubarrayComplement([]int{1, 7, 1, 1, 2, 0}, 9)

	//Zigzag Array
	fmt.Println("ZigZag Array ")
	zigZagArray := []int{4, 3, 7, 8, 6, 2, 1}

	zigZagArray = example.MakeArrayZigZagArray(zigZagArray)

	fmt.Println("", zigZagArray)
}
