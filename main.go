package main

import (
	"fmt"
)

type Difficulty struct {
	Value1 int
	Value2 int
}

func main() {
	m := map[string]Difficulty{
		"easy":        {0, 100},
		"normal":      {0, 1000},
		"dont normal": {0, 10000},
		"unreal":      {-10000, 10000},
	}

	fmt.Println("Приветствуем вас в увлекательной игре — Угадай Число!")

	fmt.Printf("Уровни сложности:\n"+
		"1) Легкий - от %d до %d\n"+
		"2) Нормальный - от %d до %d\n"+
		"3) Ненормальный - от %d до %d\n"+
		"4) Нереальный - от %d до %d\n\n",
		m["easy"].Value1, m["easy"].Value2,
		m["normal"].Value1, m["normal"].Value2,
		m["dont normal"].Value1, m["dont normal"].Value2,
		m["unreal"].Value1, m["unreal"].Value2,
	)

	fmt.Println("Введите номер сложности:")
	var diff = difficult()
	go game(diff)
}

func difficult() (scandiff uint) {
	for {
		_, err := fmt.Scan(&scandiff)
		if err != nil || scandiff <= 0 || scandiff >= 5 {
			fmt.Println("Введите номер сложности (от 1 до 4):")
			continue
		}
		break
	}
	return
}

func game(diff uint) {
	//
}
