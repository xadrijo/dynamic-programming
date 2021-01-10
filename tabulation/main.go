package main

import "fmt"

func main() {

	fmt.Println("Fibonacci Tabulation:")
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(50))

	fmt.Println("gridTraveler:")
	fmt.Println(gridTraveler(1, 1)) // 1
	fmt.Println(gridTraveler(2, 3)) // 3
	fmt.Println(gridTraveler(3, 2)) // 3
	fmt.Println(gridTraveler(3, 3)) // 3
	fmt.Println(gridTraveler(18, 18)) // 2333606220

	fmt.Println("canSum:")
	fmt.Println(canSum(7, []int{2, 3})) // true
	fmt.Println(canSum(7, []int{5, 4, 3, 7})) // true
	fmt.Println(canSum(7, []int{2, 4})) // false
	fmt.Println(canSum(8, []int{2, 3, 5})) // true
	fmt.Println(canSum(300, []int{7, 14})) // false

	fmt.Println("howSum:")
	fmt.Println(howSum(7, []int{2, 3})) // [3, 2, 2]
	fmt.Println(howSum(7, []int{5, 3, 4, 7})) // [4, 3]
	fmt.Println(howSum(7, []int{2, 4})) // null
	fmt.Println(howSum(8, []int{2, 3, 5})) // [2, 2, 2, 2]
	fmt.Println(howSum(300, []int{7, 14})) // null

	fmt.Println("bestSum:")
	fmt.Println(bestSum(7, []int{5, 3, 4, 7})) // [7]
	fmt.Println(bestSum(8, []int{2, 3, 5})) // [3, 5]
	fmt.Println(bestSum(8, []int{1, 4, 5})) // [4, 4]
	fmt.Println(bestSum(100, []int{1, 2, 5, 25})) // [25, 25, 25, 25]

	fmt.Println("canConstruct:")
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"})) // true
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"})) // false
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) // true
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // false

	fmt.Println("countConstruct:")
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"})) // 1
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"})) // 1
	fmt.Println(countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"})) // 0
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) // 4
	fmt.Println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // 0

}

func countConstruct(target string, wordBank []string) int {
	table := make([]int, len(target) + 1)
	table[0] = 1

	for i := 0; i <= len(target); i++ {
		for _, word := range wordBank {
			if (i + len(word)) < len(table) {
				current := target[i:(i + len(word))]
				// if the word matches the characters starting at position i
				if current == word {
					table[i+len(word)] += table[i]
				}
			}
		}
	}
	return table[len(target)]
}

func canConstruct(target string, wordBank []string) bool {
	table := make([]bool, len(target) + 1)
	table[0] = true

	for i := 0; i <= len(target); i++ {
		if table[i] == true {
			for _, word := range wordBank {
				if (i + len(word)) < len(table) {
					current := target[i:(i + len(word))]
					// if the word matches the characters starting at position i
					if current == word {
						table[i+len(word)] = true
					}
				}
			}
		}
	}
	return table[len(target)]
}

func bestSum(targetSum int, nums []int) []int {
	table := make([][]int, targetSum + 1)
	table[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {
			for _, n := range nums {
				combination := append(table[i], n)
				if (i + n) <= targetSum {
					// if this current combination is shorter that what is already stored
					if table[i + n] == nil || len(table[i + n]) > len(combination) {
						table[i + n] = combination
					}
				}
			}
		}
	}
	return table[targetSum]
}

func howSum(targetSum int, nums []int) []int {
	table := make([][]int, targetSum + 1)
	table[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {
			for _, n := range nums {
				if (i + n) <= targetSum {
					table[i + n] = append(table[i], n)
				}
			}
		}
	}
	return table[targetSum]
}

func canSum(targetSum int, nums []int) bool {
	table := make([]bool, targetSum + 1)
	table[0] = true

	for  i := 0; i <= targetSum; i++{
		if table[i] == true {
			for _, v := range nums {
				if (i + v) <= targetSum {
					table[i + v] = true
				}
			}
		}
	}
	return table[targetSum]
}

func gridTraveler(m int, n int) int {
	table := make([][]int, m + 1)
	for i := range table {
		table[i] = make([]int, n + 1)
	}

	table[1][1] = 1
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			current := table[i][j]
			if (j + 1) <= n {
				table[i][j + 1] += current
			}
			if (i + 1) <= m {
				table[i + 1][j] += current
			}
		}
	}

	return table[m][n]
}

func fib(n int) int {
	table := make([]int, n + 1)
	table[1] = 1

	for i := 0; i <= n - 1; i++ {
		table[i + 1] += table[i]
		if (i + 2) <= n {
			table[i + 2] += table[i]
		}
	}
	return table[n]
}