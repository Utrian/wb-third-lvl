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
	reader := bufio.NewReader(os.Stdin)

	for {
		// Приглашение
		fmt.Print("P$ ")

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
				homeDir, err := os.UserHomeDir()
				if err != nil {
					fmt.Println(err)
				}
				args = append(args, homeDir)
			}
			// В обратном случае, просто меняем директорию
			// на указанную в аргументе
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
		// Блок fork
		case "fork":
			// Передаем имя текущего исполняемого файла и переданные ему
			// аргументы
			cmd := exec.Command(os.Args[0], append(os.Args[1:], "forked")...)
			// Созданный процесс наследует stdin, -out, -err
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			// Т.к. мы создаем отдельный процесс, то просто запускаем его
			// не дожидаясь завершения используя Start()
			err := cmd.Start()
			if err != nil {
				fmt.Println("Error:", err)
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
