package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"test.com/note/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving file failed")
		return
	}

	fmt.Println("file saved successfully")

}

func getUserInput(prompt string) string {
	if prompt == "" {
		return ""
	}
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

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", "", err
	// }
	content := getUserInput("Note content: ")
	return title, content
}
