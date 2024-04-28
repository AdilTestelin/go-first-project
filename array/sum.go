package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sumOfNumbers []int

	for _, arr := range numbersToSum {
		sumOfNumbers = append(sumOfNumbers, Sum(arr))
	}

	return sumOfNumbers
}

func SumAllTails(numbersToSum ...[]int) []int {
	var result []int

	for _, arr := range numbersToSum {
		arrayLength := len(arr)
		if arrayLength == 0 {
			result = append(result, 0)
		} else {
			tail := arr[1:]
			result = append(result, Sum(tail))
		}
	}

	return result
}
