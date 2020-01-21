package main

import (
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.OnErrorCode(iris.StatusNotFound, handle)
	app.Get("/bubble", func(ctx iris.Context) {
		array := parseArray(ctx.URLParam("array"))
		sorted := sort(array)
		ctx.WriteString(arrayToString(sorted))
	})
	app.Run(iris.Addr(":80"))
}

func handle(ctx iris.Context) {
	ctx.WriteString("Not Found ;). " + ctx.Request().Host + ctx.Request().URL.Path)
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
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if array[j] > array[j+1] {
				buf := array[j]
				array[j] = array[j+1]
				array[j+1] = buf
			}
		}
	}

	return array
}
