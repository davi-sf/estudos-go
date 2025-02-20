package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {

	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed")
		return err
	}
	fmt.Println("Saving the data succeeded")
	return nil

}

func getUserInput(prompt string) string {

	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	return strings.TrimSpace(text)

}

func getNoteData() (string, string) {

	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content

}

func getTodoData() string {

	content := getUserInput("Todo text: ")

	return content

}

func main() {

	title, content := getNoteData()

	todoText := getTodoData()

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)

	if err != nil {
		return
	}

	outputData(userNote)

}
