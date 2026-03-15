package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text   string    `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("%s\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(todo)
	
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (Todo, error) {
		if text == "" {
		return Todo{}, errors.New("Invalid input")
	}
	return Todo{
		Text: text,
	}, nil
}

