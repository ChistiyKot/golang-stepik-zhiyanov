/*
Напишите функцию shuffle(), которая тасует элементы среза в случайном порядке. Функция должна отрабатывать in-place, то есть менять содержимое переданного среза, а не создавать новый срез. Чтобы перетасовать элементы, используйте функцию
https://stepik.org/lesson/526869/step/9?unit=519588
*/
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func shuffle(nums []int) {
	rand.Seed(42)
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

func main() {
	nums := readInput()
	shuffle(nums)
	fmt.Println(nums)
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
