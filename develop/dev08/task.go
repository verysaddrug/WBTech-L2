package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const intro = `Welcome to the dev08 - a bash emulator
Usage: <command> <args>
requires bash or Git-bash/WSL to work
Ctrl+C to exit`

func main() {
	fmt.Println(intro)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		commands := strings.Split(input, "|")
		if len(commands) > 1 {
			handlePipes(commands)
		} else {
			handleCommand(input)
		}
	}
}

// handleCommand обрабатывает отдельные команды
func handleCommand(input string) {
	args := strings.Fields(input)
	cmd := args[0]
	args = args[1:]

	switch cmd {
	case "cd":
		if len(args) == 0 {
			fmt.Println("cd: missing argument")
		} else {
			err := os.Chdir(args[0])
			if err != nil {
				fmt.Println(err)
			}
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(dir)
		}
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "kill":
		executeExternalCommand("kill", args)
	case "ps":
		executeExternalCommand("ps", args)
	default:
		executeExternalCommand(cmd, args)
	}
}

// executeExternalCommand запускает внешнюю команду
func executeExternalCommand(cmd string, args []string) {
	command := exec.Command(cmd, args...)
	command.Stdin, command.Stdout, command.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// handlePipes обрабатывает команды, соединенные через пайпы
func handlePipes(commands []string) {
	var cmds []*exec.Cmd
	for _, cmdStr := range commands {
		cmdArgs := strings.Fields(cmdStr)
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		cmds = append(cmds, cmd)
	}

	// Настройка пайпов между командами
	for i := 0; i < len(cmds)-1; i++ {
		stdout, err := cmds[i].StdoutPipe()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmds[i+1].Stdin = stdout
	}

	// Запуск команд
	for _, cmd := range cmds {
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Ожидание завершения команд
	for _, cmd := range cmds {
		err := cmd.Wait()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
