package main

import (
	"bufio"
	"fmt"
	"lukasrar/go-essentials/exercises/structs/practice-project/note"
	"lukasrar/go-essentials/exercises/structs/practice-project/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	displayer
	saver
}

func main() {
	title, content := getNoteData()

	createdNote, errNote := note.New(title, content)
	todoText := getUserInput("Todo text")

	todo, errTodo := todo.New(todoText)

	if errNote != nil || errTodo != nil {
		fmt.Println(errNote)
		return
	}

	displayData(createdNote)
	displayData(todo)

}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func displayData(data outputtable) {
	data.Display()
	saveData(data)
}

func saveData(data saver) error {

	err := data.Save()

	if err != nil {
		fmt.Println("Error while saving")
		return err
	}

	fmt.Println("Note succesfully saved")
	return nil
}

func getUserInput(prompt string) string {

	fmt.Print(prompt)

	r := bufio.NewReader(os.Stdin)
	txt, err := r.ReadString('\n')

	if err != nil {
		return ""
	}

	txt = strings.TrimSuffix(txt, "\n")
	txt = strings.TrimSuffix(txt, "\r")

	return txt

}
