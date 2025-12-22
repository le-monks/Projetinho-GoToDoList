package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"GoToDoList/utils"
)

func main() {

	var myList utils.TodoList

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem-vindo ao toDoList - aqui você realiza suas tarefas")
	fmt.Println("Comandos: 'add' [task], 'list', 'done' [index], 'end'")

while:
	for {

		if scanner.Scan() {

			typedText := scanner.Text()
			parts := strings.SplitN(typedText, " ", 2) //split = cria uma slice de tamanho 2. slice[0] = comando; slice[1] = argumento

			command := parts[0]
			arg := ""

			if len(parts) > 1 {
				arg = parts[1]
			}

			switch command {

			case "command":
				fmt.Println("Comandos: 'add', 'list', 'done', 'end'")

			case "add":
				note := strings.Trim(arg, " \"'") //Aqui trim faz o tratamento e remove " ' ou espaços em branco

				if note == "" {
					fmt.Println("Tarefa vazia")
				} else {
					myList.Add(note)
				}

			case "list":
				fmt.Println("")
				fmt.Println("A seguir, sua lista de afazeres: ")
				myList.PrintList()

			case "done":

				index, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Println("Digite um número válido")
				} else {
					myList.Remove(index)
				}

			case "end":
				break while
			default:
				fmt.Println()
			}
		}
	}
}
