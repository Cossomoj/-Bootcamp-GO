package main

import "fmt"

func main() {
	//Найти первую, вторую и последнюю  цифры задонного трехзначного числа.
	var input int
	fmt.Println("Введи число трехзначное")

	for {
		fmt.Scanf("%d\n", &input)
		if input >= 100 && input < 1000 {
			first := input / 100
			second := input / 10 % 10
			three := input % 10
			fmt.Println("Первое:", first, "\nВторое:", second, "\nТретье:", three)
			return
		} else {
			fmt.Println("Ты уверен, что знаешь что такое трехзначное число? Пробуй еще!!!")
		}
	}
}
