package main

func getConcatenation(nums []int) []int {
	s := []int{}
	result := append(s,nums...)
	result = append(result, nums...)
	return  result

}

