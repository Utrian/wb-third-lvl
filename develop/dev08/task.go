package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Необходимо реализовать свою собственную UNIX-шелл-утилиту с
// поддержкой ряда простейших команд: cd, pwd, echo, kill, ps

func main() {
	for {
		// Приглашение
		fmt.Print("P$ ")

		// Ридер для чтения команд
		reader := bufio.NewReader(os.Stdin)

		// Получаем введенную команду
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		switch cmdName := args[0]; cmdName {
		// Блок выхода
		case "exit":
			return
		// Блок cd
		case "cd":
			// Если аргументы отсутствуют, то меняем
			// текущую директорию на пользовательскую
			if len(args) < 2 {
				wordDirectory, err := os.UserHomeDir()
				if err != nil {
					fmt.Println(err)
				}
				args = append(args, wordDirectory)
			}
			// В обратном случае, просто меняем директорию
			// на указанную в аргументе
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}

		// Данный блок покрывает большинство утилит:
		// pwd, echo, kill, ps
		default:
			cmd := exec.Command(cmdName, args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}
