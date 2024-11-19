package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type ToDo struct {
	Text string `json:"text"`
}

func (t ToDo) Display() {
	fmt.Println("========")
	fmt.Println(t.Text)
	fmt.Println("========")
}

func (t ToDo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (ToDo, error) {

	if text == "" {
		return ToDo{}, errors.New("text is required")
	}

	return ToDo{
		text,
	}, nil
}
