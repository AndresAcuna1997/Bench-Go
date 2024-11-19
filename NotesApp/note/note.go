package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (n Note) Display() {
	fmt.Println("========")
	fmt.Println("Title: ", n.Title)
	fmt.Println("Description: ", n.Description)
	fmt.Println("========")
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, description string) (Note, error) {

	if title == "" || description == "" {
		return Note{}, errors.New("title or description are required")
	}

	return Note{
		title,
		description,
	}, nil
}
