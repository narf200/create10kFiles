package main

import (
	"fmt"
	"os"
	"strconv"
)

var amountOfFiles = 10000
var startCount = 1
var largePortion = 9999

func main() {
	// цикл, который отработает 10к раз
	for count := startCount; count < amountOfFiles; count++ {

		// создание имени для файлов
		fileName := "file" + strconv.Itoa(count) + ".txt"
		// создание файла
		file, err := os.Create(fileName)

		// обработка ошибки
		if err != nil {
			fmt.Printf("Ошибка при создании файла %s: %v\n", fileName, err)
			continue // Продолжить цикл, если произошла ошибка
		}

		if count < largePortion {
			fmt.Fprintln(file, "Хуй в правой руке")
		} else {
			fmt.Fprintln(file, "Хуй в левой руке")
		}

		defer file.Close()
		fmt.Println(count)
	}
}
