package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notepad/note"
	"example.com/notepad/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()
	tod, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(tod)
	if err != nil {
		return
	}
	err = outputData(userNote)
	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	err := saveData(data)
	if err != nil {
		return err
	}
	return nil
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}
func getTodoData() (content string) {
	content = getUserInput("Todo content:")
	return
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
