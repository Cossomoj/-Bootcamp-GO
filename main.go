package main

import "fmt"

func main() {
	var input int
	fmt.Println("Введи число трехзначное")

	for {
		fmt.Scanf("%d\n", &input)
		if input >= 100 && input < 1000 {
			first := input / 100
			three := input % 10
			fmt.Println(first, three)
			return
		} else {
			fmt.Println("Ты уверен, что знаешь что такое трехзначное число? Пробуй еще!!!")
		}
	}
}
