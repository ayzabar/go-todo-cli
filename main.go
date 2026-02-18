package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID     int    `json:"id"`
	Value  string `json:"value"`
	IsDone bool   `json:"is_done"`
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	taskList := []Task{}

	taskList = append(taskList, Task{ID: 1, Value: "Become a Systems Engineer", IsDone: false})

MainLoop:
	for {
		clearScreen()

		fmt.Println("################################")
		fmt.Println("###   AYZABAR'S TO-DO CLI    ###")
		fmt.Println("################################")

		listTask(taskList)
		fmt.Println("--------------------------------")

		fmt.Println("1) Add Task")
		fmt.Println("2) Delete Task")
		fmt.Println("3) Mark Done/Undone")
		fmt.Println("4) Save List")
		fmt.Println("5) Quit")

		choice := getChoice(reader)

		switch choice {
		case 1:
			fmt.Print("Task Name: ")
			reader.Scan()
			newTask := reader.Text()

			taskList = reorderTask(append(taskList, addTask(taskList, newTask)))

			fmt.Println("\n✅ Task added successfully!")
			pause(reader)

		case 2:
			taskList = deleteTask(reader, taskList)
			pause(reader)

		case 3:
			markTask(reader, taskList)
			pause(reader)

		case 4:
			fmt.Println("\n💾 Saving feature coming soon...")
			pause(reader)

		case 5:
			clearScreen()
			fmt.Println("👋 See you later, Space Cowboy.")
			break MainLoop
		}
	}
}

// Works on Mac and Linux. Stop using windows man fr.
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func pause(scanner *bufio.Scanner) {
	fmt.Print("\nPress Enter to continue...")
	scanner.Scan()
}

func getChoice(scanner *bufio.Scanner) int {
	for {
		fmt.Print("\nChoice [1-5]: ")
		scanner.Scan()
		text := scanner.Text()

		choice, err := strconv.Atoi(text)
		if err == nil && choice >= 1 && choice <= 5 {
			return choice
		}
		fmt.Println("❌ Invalid choice. Please try again.")
	}
}

func addTask(l []Task, s string) Task {
	return Task{
		ID:     len(l) + 1,
		Value:  s,
		IsDone: false,
	}
}

func deleteTask(scanner *bufio.Scanner, l []Task) []Task {
	for {
		fmt.Print("\nEnter ID to delete (0 to cancel): ")
		scanner.Scan()
		text := scanner.Text()
		id, err := strconv.Atoi(text)

		if err == nil {
			if id == 0 {
				fmt.Println("Operation cancelled.")
				return l
			}
			if id > 0 && id <= len(l) {
				l = append(l[:id-1], l[id:]...)
				fmt.Println("\n🗑️  Task deleted.")
				return reorderTask(l)
			}
		}
		fmt.Println("❌ Invalid ID! Please enter a number from the list.")
	}
}

func markTask(scanner *bufio.Scanner, l []Task) {
	for {
		fmt.Print("\nEnter ID to mark (0 to cancel): ")
		scanner.Scan()
		text := scanner.Text()
		id, err := strconv.Atoi(text)

		if err == nil {
			if id == 0 {
				fmt.Println("Operation cancelled.")
				return
			}
			if id > 0 && id <= len(l) {
				l[id-1].IsDone = !l[id-1].IsDone

				status := "unmarked"
				if l[id-1].IsDone {
					status = "completed"
				}
				fmt.Printf("\n✨ Task marked as %s!", status)
				return
			}
		}
		fmt.Println("❌ Invalid ID! Please enter a number from the list.")
	}
}

func reorderTask(l []Task) []Task {
	for i := range l {
		l[i].ID = i + 1
	}
	return l
}

func listTask(l []Task) {
	if len(l) == 0 {
		fmt.Println("   (List is empty, add something!)")
		return
	}

	fmt.Println("ID   STATUS   TASK")
	fmt.Println("--   ------   ----")
	for _, t := range l {
		status := "[ ]"
		if t.IsDone {
			status = "[X]"
		}
		fmt.Printf("%-4d %-8s %s\n", t.ID, status, t.Value)
	}
}
