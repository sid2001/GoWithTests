package main

// func Sum(numbers [5]int) int {
//	sum := 0
//	for i := 0; i < 5; i++ {
//		sum += numbers[i]
//	}
//	return sum
//}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, numbers := range numbers {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
