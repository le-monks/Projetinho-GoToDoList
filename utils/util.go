package utils

import "fmt"

type Task struct {
	Name string
	Data string
	Note string
	Done bool
}

type TodoList struct {
	Tasks []Task
}

func (list *TodoList) Add(note string) {
	newTask := Task{Note: note}
	list.Tasks = append(list.Tasks, newTask)
	fmt.Println("Item adicionado!")
}

func (list *TodoList) PrintList() {
	for i, p := range list.Tasks {
		fmt.Printf("%d.\t%s\n", i+1, p.Note)
	}
}

func (list *TodoList) Remove(index int) {

	i := index - 1
	taskNote := list.Tasks[i].Note
	list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)

	fmt.Printf("Tarefa '%s' concluída e removida da lista.\n", taskNote)
}
