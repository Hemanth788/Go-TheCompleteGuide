package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go.com/struct-project/note"
	"go.com/struct-project/todo"
)

type ISaver interface {
	Save() error
}

type IOutputtable interface {
	ISaver
	Display()
}

// need not connect this interface explicitly to any struct type in Go
func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	printSomething("Which")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getNoteData() (string, string) {
	title := getUserInput("Title: ")

	content := getUserInput("Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt + " ")
	var value string

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // no double quotes, rune, single character

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	value = text

	return value
}

func saveData(data ISaver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func outputData(data IOutputtable) error {
	data.Display()
	return saveData(data)
}

func printSomething(value any) { // (value interface{})
	// intVal, ok := value.(int)
	// stringVal, ok := value.(string)
	// float64Val, ok := value.(float64)
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	case string:
		fmt.Println("String: ", value)
	}
}
