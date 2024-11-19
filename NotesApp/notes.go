package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save() error
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	title, description := getNoteData()
	text := getUserInput("Todo: ")
	note, err := note.New(title, description)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note)

	todo, err := todo.New(text)

	outputData(todo)

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving the note")
		return err
	}

	fmt.Println("File saved!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Title Note: ")
	description := getUserInput("Title description: ")

	return title, description
}

func getUserInput(propmt string) string {
	fmt.Print(propmt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
