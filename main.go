package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type task struct {
	name string
	data string
	note string
	done bool
}

type TodoList struct {
	tasks []task
}

func main() {

	var myList TodoList

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem-vindo ao toDoList - aqui você realiza suas tarefas")
	fmt.Println("Comandos: 'add', 'list', 'done', 'end'")
While:
	for {
		if scanner.Scan() {
			typedText := scanner.Text()

			switch typedText {
			case "add":
				fmt.Println("Insira um item na lista: ")
				scanner.Scan()
				newTask := task{note: scanner.Text()}
				myList.add(newTask)

			case "list":
				fmt.Println("")
				fmt.Println("A seguir, sua lista de afazeres: ")
				myList.printList()

			case "done":
				if len(myList.tasks) == 0 {
					fmt.Println("Lista vazia!")
					break
				}
				fmt.Println("Digite o número da tarefa que deseja remover:")
				myList.printList()

				scanner.Scan()
				indexRemove := scanner.Text()

				index, _ := strconv.Atoi(indexRemove)
				err := myList.Remove(index)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
			case "end":
				break While
			default:
				fmt.Println("Token inválido!")
			}
		}
	}
}

func (list *TodoList) add(t task) {
	list.tasks = append(list.tasks, t)
	fmt.Println("Item adicionado!")
}

func (list *TodoList) printList() {
	for i, p := range list.tasks {
		fmt.Printf("%d.\t%s\n", i+1, p.note)
	}
}

func (list *TodoList) Remove(index int) error {
	i := index - 1

	if i < 0 || i >= len(list.tasks) {
		return fmt.Errorf("Índice inválido!")
	}
	taskNote := list.tasks[i].note
	list.tasks = append(list.tasks[:i], list.tasks[i+1:]...)

	fmt.Printf("Tarefa \"%s\" concluída e removida da lista.\n", taskNote)
	return nil
}
