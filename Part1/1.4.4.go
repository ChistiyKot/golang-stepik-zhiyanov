/*
Напишите функцию filter(), которая фильтрует срез целых чисел с помощью функции-предиката и возвращает отфильтрованный срез.
https://stepik.org/lesson/526869/step/4?unit=519588
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	var result []int
	for _, val := range iterable {
		if predicate(val) == true {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	src := readInput()

	res := filter(func(x int) bool { return x%2 == 0 }, src)

	fmt.Println(res)
}

func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
