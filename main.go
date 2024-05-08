package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Difficulty struct {
	Value1 int
	Value2 int
}

func main() {
	m := map[uint]Difficulty{
		1: {0, 100},
		2: {0, 1000},
		3: {0, 10000},
		4: {-10000, 10000},
	}

	fmt.Println("Приветствую вас в увлекательной игре — Угадай Число!")

	fmt.Printf("Уровни сложности: \n"+
		"1) Легкий — от %d до %d\n"+
		"2) Нормальный — от %d до %d\n"+
		"3) Ненормальный — от %d до %d\n"+
		"4) Нереальный — от %d до %d\n\n",
		m[1].Value1, m[1].Value2,
		m[2].Value1, m[2].Value2,
		m[3].Value1, m[3].Value2,
		m[4].Value1, m[4].Value2,
	)

	var diff = difficult(m)
	game(diff, m)
}

func difficult(diffmap map[uint]Difficulty) (scandiff uint) {
	diffcount := uint(len(diffmap))
	fmt.Println("Введите номер сложности: ")
	for {
		_, err := fmt.Scan(&scandiff)
		if err != nil || scandiff <= 0 || scandiff > diffcount {
			fmt.Printf("Введите номер сложности (от 1 до %d): ", diffcount)
			continue
		}
		break
	}
	return
}

func game(diff uint, diffmap map[uint]Difficulty) {
	var guess int
	minNum, maxNum, randomNumber := generateNumber(diff, diffmap)

	gameInfo(minNum, maxNum)

gameLoop:
	for {
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Print("Вводите только числа! ")
			gameInfo(minNum, maxNum)
			continue
		}
		switch {
		case guess < randomNumber:
			fmt.Println("Загаданное число больше")
		case guess > randomNumber:
			fmt.Println("Загаданное число меньше")
		case guess == randomNumber:
			fmt.Println("Поздравляю! Вы угадали!")
			break gameLoop
		}
	}

	for {
		fmt.Println("\nХотите сыграть еще раз? (y/n)")
		var input string
		fmt.Scanln(&input)
		if input != "y" && input != "Y" && input != "yes" && input != "YES" {
			break
		}

		diff = difficult(diffmap)

		minNum, maxNum, randomNumber = generateNumber(diff, diffmap)
		gameInfo(minNum, maxNum)
		goto gameLoop
	}
}

func generateNumber(diff uint, diffmap map[uint]Difficulty) (minNum, maxNum, randomNumber int) {
	minNum = diffmap[diff].Value1
	maxNum = diffmap[diff].Value2

	rand.Seed(time.Now().UnixNano())
	randomNumber = rand.Intn(maxNum-minNum) + minNum
	return
}

func gameInfo(minNum, maxNum int) {
	fmt.Printf("Угадайте число от %d до %d: \n", minNum, maxNum)
}
