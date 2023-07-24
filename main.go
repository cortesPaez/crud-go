package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/cortesPaez/crud-go/task"
)

func main() {
	// read or create the tasks.json file
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []task.Task

	// get file info
	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	// check the size
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	// check the user input and print options to choose
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Cual es su tarea?")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		task.SaveTasks(file, tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar un id para eliminar")
			return
		}
		id, error := strconv.Atoi(os.Args[2])
		if error != nil {
			fmt.Println("el id debe ser un numero y no un caracter")
		}
		tasks = task.DeleteTask(tasks, id)
		task.SaveTasks(file, tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar un id para marcar la tarea como finalizada")
			return
		}
		id, error := strconv.Atoi(os.Args[2])
		if error != nil {
			fmt.Println("el id debe ser un numero y no un caracter")
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveTasks(file, tasks)
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Uso: crud-go [list|add|complete|delete]")
	return
}
