package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface { // if interface has 1 method, name of interface should be method name plus 'er'
	Save() error
}

// type displayer interface {
// 	Display()
// }

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
		return err
	}
	return nil
}

func printSomething(value interface{}) { // any value allowed
	switch value.(type) {
	case float64:
		fmt.Println("Float: ", value)
	case int:
		fmt.Println("Integer: ", value)
	case outputtable:
		typedValue, ok := value.(outputtable)
		if ok {
			fmt.Println("outputtable")
			typedValue.Display()
		}

	}
}

func main() {
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	printSomething(userNote)
	err = outputData(userNote)
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeeded")
	err = outputData(todo)
	if err != nil {
		fmt.Println("Saving the todo failed")
		return
	}
	fmt.Println("Saving the todo succeeded")

}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
