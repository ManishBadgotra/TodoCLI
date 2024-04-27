package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	Title string
	Body  string
}

func main() {

	var (
		data     Note
		dataList []Note
	)

	fmt.Println("Todo App")
continuedForLoop:
	for {
		fmt.Print("Enter a command: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		switch {
		case input == "add\n":
			dataList = data.Add(dataList, reader)
		case input == "edit\n":
			fmt.Print("Enter the title of the note to edit: ")
			editTitle, _ := reader.ReadString('\n')
			dataList = data.Edit(dataList, editTitle, reader)
		case input == "delete\n":
			fmt.Print("Enter the title of the note to delete: ")
			delTitle, _ := reader.ReadString('\n')
			dataList = data.Delete(dataList, delTitle, reader)
		case input == "show\n":
			data.Show(dataList)
		case input == "exit\n":
			break continuedForLoop
		}
	}
	fmt.Println("Program Exited!")
}

func (n Note) Add(dataList []Note, reader *bufio.Reader) []Note {

	fmt.Print("Enter a title: ")
	title, _ := reader.ReadString('\n')
	n.Title = title

	fmt.Print("Enter a Description: ")
	body, _ := reader.ReadString('\n')
	n.Body = body

	dataList = append(dataList, n)
	return dataList

}

func (n *Note) Show(dataList []Note) {
	for i, v := range dataList {
		fmt.Print("Note ", i+1, ": ", v.Title, "\n\t\t", v.Body, "\n")
	}
}

func (n *Note) Edit(dataList []Note, editTitle string, reader *bufio.Reader) []Note {
	for i, v := range dataList {
		if strings.ToLower(v.Title) == editTitle {

			fmt.Print("Enter new title: ")
			title, _ := reader.ReadString('\n')
			n.Title = title

			fmt.Print("Enter new Description: ")
			body, _ := reader.ReadString('\n')
			n.Body = body

			dataList[i] = *n
			return dataList
		}
	}
	return dataList
}

func (n *Note) Delete(dataList []Note, delTitle string, reader *bufio.Reader) []Note {
	for i, v := range dataList {
		if strings.ToLower(v.Title) == delTitle {
			dataList = append(dataList[:i], dataList[i+1:]...)
			return dataList
		}
	}
	return dataList
}
