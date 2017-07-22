package utils

import (
	"os"
	"strconv"
	"bufio"
)

func WriteRandom2File(filePath string, arr []int) {
	file, err := os.Create(filePath)
	if err == nil {
		for _, v := range arr {
			file.WriteString(strconv.Itoa(v))
			file.WriteString("\n")
		}
	}
}

func ReadRandom(filePath string, size int) []int {
	file, _ := os.Open(filePath)
	read := bufio.NewReader(file)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		line, _, _ := read.ReadLine()
		arr[i], _ = strconv.Atoi(string(line))
	}
	return arr
}
