package main

import (
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Get("/selection", func(ctx iris.Context) {
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

func sort(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		index := i
		for j := i; j < length; j++ {
			if array[index] > array[j] {
				index = j
			}
		}
		buf := array[index]
		array[index] = array[i]
		array[i] = buf
	}
	return array
}
