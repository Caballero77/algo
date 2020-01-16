package main

import (
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Get("/heap", func(ctx iris.Context) {
		array := parseArray(ctx.URLParam("array"))
		sorted := sort(array)
		ctx.WriteString(arrayToString(sorted))
	})

	app.Run(iris.Addr(":80"))
}

func parseArray(array string) []int {
	numbers := strings.Split(array, ",")
	length := len(numbers)
	result := make([]int, length)
	for i := 0; i < length; i++ {
		value, _ := strconv.Atoi(numbers[i])
		result[i] = value
	}
	return result
}

func arrayToString(array []int) string {
	str := strconv.Itoa(array[0])
	length := len(array)
	for i := 1; i < length; i++ {
		str += "," + strconv.Itoa(array[i])
	}
	return str
}

func swap(array []int, x int, y int) []int {
	buf := array[x]
	array[x] = array[y]
	array[y] = buf

	return array
}

func heapify(array []int, index int, n int) []int {
	if index >= n/2 {
		return array
	}

	i := index + 1

	if !(array[index] >= array[i*2-1] && (n == i*2 || array[index] >= array[i*2])) {
		if n == i*2 || array[i*2] < array[i*2-1] {
			array = swap(array, index, i*2-1)
		} else {
			array = swap(array, index, i*2)
		}
	}

	array = heapify(array, i*2-1, n)
	array = heapify(array, i*2, n)
	return array
}

func sort(array []int) []int {
	n := len(array)

	// Build max heap
	for i := n / 2; i > 0; i-- {
		array = heapify(array, i-1, n)
	}

	for n > 0 {
		array = swap(array, 0, n-1)
		n--
		array = heapify(array, 0, n)
	}

	return array
}
