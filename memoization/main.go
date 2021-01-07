package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Fibonacci:")
	fmt.Println(fib(6, map[int]int{})) // 8
	fmt.Println(fib(50, map[int]int{})) // 12586269025

	fmt.Println("Grid Traveler:")
	fmt.Println(gridTraveler(2, 3, map[string]int{})) // 3
	fmt.Println(gridTraveler(18, 18, map[string]int{})) // 2333606220

	fmt.Println("Can Sum:")
	fmt.Println(canSum(7, []int{2, 3}, map[int]bool{})) // true
	fmt.Println(canSum(300, []int{7, 14}, map[int]bool{})) // false

	fmt.Println("How Sum:")
	fmt.Println(howSum(7, []int{2, 3}, map[int][]int{})) // [3 2 2]
	fmt.Println(howSum(7, []int{5, 3, 4, 7}, map[int][]int{})) // [4 3]
	fmt.Println(howSum(7, []int{2, 4}, map[int][]int{}))
	fmt.Println(howSum(8, []int{2, 3, 5}, map[int][]int{}))
	fmt.Println(howSum(300, []int{7, 14}, map[int][]int{}))

	fmt.Println("Best Sum:")
	fmt.Println(bestSum(7, []int{5, 3, 4, 7}, map[int][]int{})) // [7]
	fmt.Println(bestSum(8, []int{2, 3, 5}, map[int][]int{})) // [5, 3]
	fmt.Println(bestSum(8, []int{1, 4, 5}, map[int][]int{})) // [4, 4]
	fmt.Println(bestSum(100, []int{1, 2, 5, 25}, map[int][]int{})) // [25, 25, 25, 25]

	fmt.Println("Can Construct:")
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string]bool{})) // true
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]bool{})) // false
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, map[string]bool{})) // true
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, map[string]bool{})) // false

	fmt.Println("Count Construct:")
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, map[string]int{})) // 2
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string]int{})) // 1
	fmt.Println(countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]int{})) // 0
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, map[string]int{})) // 4
	fmt.Println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, map[string]int{})) // 0

	fmt.Println("All Construct:")
	fmt.Println(allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, map[string][][]string{})) // [[purp le] [p ur p le]]
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}, map[string][][]string{})) //[[ab cd ef] [ab c def] [abc def] [abcd ef]]
	fmt.Println(allConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string][][]string{})) // []
}

func allConstruct(target string, wordBank []string, memo map[string][][]string) [][]string {
	r, found := memo[target]
	if found {
		return r
	}
	if target == "" {
		return [][]string{{}}
	}

	var nilSlice [][]string
	if len(target) == 0 {
		return nilSlice
	}

	var result [][]string
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			suffixWays := allConstruct(suffix, wordBank, memo)

			if suffixWays != nil {
				var targetWays [][]string
				for _, way := range suffixWays {
					if way != nil {
						way = insertString(way, word, 0)
					}
					targetWays = append(targetWays, way)
				}
				result = append(result, targetWays...)
			}
		}
	}
	memo[target] = result
	return result
}

func insertString(array []string, value string, index int) []string {
	return append(array[:index], append([]string{value}, array[index:]...)...)
}


func countConstruct(target string, wordBank []string, memo map[string]int) int {
	r, found := memo[target]
	if found {
		return r
	}
	if len(target) == 0 {
		return 1
	}

	totalCount := 0
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			numWaysForRest := countConstruct(target[len(word):], wordBank, memo)
			totalCount += numWaysForRest
		}
	}

	memo[target] = totalCount
	return totalCount
}

func canConstruct(target string, wordBank []string, memo map[string]bool) bool {
	r, found := memo[target]
	if found {
		return r
	}

	if len(target) == 0 {
		return true
	}

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			if canConstruct(suffix, wordBank, memo) {
				memo[target] = true
				return true
			}
		}
	}
	memo[target] = false
	return false
}

func bestSum(targetSum int, nums []int, memo map[int][]int) []int {
	r, found := memo[targetSum]
	if found {
		return r
	}

	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	var shortestCombination []int

	for _, v := range nums {
		remainder := targetSum - v
		remainderCombination := bestSum(remainder, nums, memo)
		if remainderCombination != nil {
			combination := append(remainderCombination, v)
			// if the combination is shorter that the current "shortest", update it
			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
	}
	memo[targetSum] = shortestCombination
	return shortestCombination
}

func howSum(targetSum int, nums []int, memo map[int][]int) []int {
	v, found := memo[targetSum]
	if found {
		return v
	}

	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, v := range nums {
		remainder := targetSum - v
		remainderResult := howSum(remainder, nums, memo)
		if remainderResult != nil {
			memo[targetSum] = append(remainderResult, v)
			return memo[targetSum]
		}
	}

	memo[targetSum] = nil
	return nil
}

func canSum(targetSum int, nums []int, memo map[int]bool) bool {
	value, found := memo[targetSum]
	if found {
		return value
	}
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, v := range nums {
		remainder := targetSum - v
		if canSum(remainder, nums, memo) == true {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}

func gridTraveler(m int, n int, memo map[string]int) int {
	key := fmt.Sprintf("%d%s%d", m, ",", n)
	_, found := memo[key]
	if found {
		return memo[key]
	}

	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	memo[key] = gridTraveler(m - 1, n, memo) + gridTraveler(m, n - 1, memo)
	return memo[key]
}

func fib(n int, memo map[int]int) int {
	_, found := memo[n]
	if found {
		return memo[n]
	}

	if n <= 2 {
		return 1
	}

	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}
