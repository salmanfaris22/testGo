package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	TodoItems []TodoItems
}
type TodoItems struct {
	title string
	desc  string
}

func (t *Todo) AddTodo(title string, desc string) {
	tempTodo := TodoItems{
		title: title,
		desc:  desc,
	}
	t.TodoItems = append(t.TodoItems, tempTodo)
	fmt.Println("Bloge Added 🦚")
}
func (t *Todo) ListTodo() {
	for i, v := range t.TodoItems {
		fmt.Printf("%d: title : %s desc:%s \n", i+1, v.title, v.desc)
	}
}
func main() {
	var todo Todo
	fmt.Println("Welcome To Todo Apps")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("_____________________")
		fmt.Println("Bloge Option:")
		fmt.Println("1 : ADD Todo 🫡")
		fmt.Println("2 : List Blog 🫠")
		fmt.Println("6: Exit")
		fmt.Println("_____________________")
		fmt.Println("🫵🏻 Enter Your Option : ")
		scanner.Scan()
		input := scanner.Text()
		newinput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalide Number ")
			continue
		}
		switch newinput {
		case 1:
			fmt.Println("Enter Your Todo title:")
			scanner.Scan() //title
			todoTitle := scanner.Text()
			fmt.Println("Enter Your Todo Desc:")
			scanner.Scan() //Desc
			todoDesc := scanner.Text()
			todo.AddTodo(todoTitle, todoDesc)
		case 2:
			fmt.Println("View Todo")
			todo.ListTodo()
		case 6:
			fmt.Println("Exit Todo")
			return
		}
	}
}
