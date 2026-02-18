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

	// [UI UPDATE] Let's add a dummy task so the list isn't empty on start
	taskList = append(taskList, Task{ID: 1, Value: "Become a Systems Engineer", IsDone: false})

MainLoop:
	for {
		// [UI UPDATE] 1. Clear the screen at the start of every loop
		clearScreen()

		// [UI UPDATE] 2. Dashboard Header
		fmt.Println("################################")
		fmt.Println("###   AYZABAR'S TO-DO CLI    ###")
		fmt.Println("################################")

		// [UI UPDATE] 3. Print the list immediately (The Dashboard)
		listTask(taskList)
		fmt.Println("--------------------------------")

		// [UI UPDATE] 4. Menu Options (Removed "Print List" option)
		fmt.Println("1) Add Task")
		fmt.Println("2) Delete Task")
		fmt.Println("3) Mark Done/Undone")
		fmt.Println("4) Save List")
		fmt.Println("5) Quit")

		choice := getChoice(reader)

		switch choice {
		case 1: // ADD
			fmt.Print("Task Name: ") // [UI UPDATE] English prompt
			reader.Scan()
			newTask := reader.Text()

			taskList = reorderTask(append(taskList, addTask(taskList, newTask)))

			fmt.Println("\n✅ Task added successfully!")
			pause(reader) // [UI UPDATE] Wait for user to read the message

		case 2: // DELETE
			taskList = deleteTask(reader, taskList)
			pause(reader) // [UI UPDATE] Wait

		case 3: // MARK
			markTask(reader, taskList)
			pause(reader) // [UI UPDATE] Wait

		case 4: // SAVE
			fmt.Println("\n💾 Saving feature coming soon...")
			pause(reader) // [UI UPDATE] Wait

		case 5: // QUIT
			clearScreen() // [UI UPDATE] Clean exit
			fmt.Println("👋 See you later, Space Cowboy.")
			break MainLoop
		}
	}
}

// --- HELPER FUNCTIONS ---

// [UI UPDATE] Clears the terminal screen (ANSI Escape Codes)
// Works on Mac (your machine) and Linux.
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// [UI UPDATE] Pauses execution until Enter is pressed
func pause(scanner *bufio.Scanner) {
	fmt.Print("\nPress Enter to continue...")
	scanner.Scan()
}

func getChoice(scanner *bufio.Scanner) int {
	for {
		fmt.Print("\nChoice [1-5]: ") // [UI UPDATE] English prompt
		scanner.Scan()
		text := scanner.Text()

		choice, err := strconv.Atoi(text)
		// [UI UPDATE] Check range 1-5
		if err == nil && choice >= 1 && choice <= 5 {
			return choice
		}
		fmt.Println("❌ Invalid choice. Please try again.")
	}
}

// --- LOGIC FUNCTIONS ---

func addTask(l []Task, s string) Task {
	return Task{
		ID:     len(l) + 1,
		Value:  s,
		IsDone: false,
	}
}

func deleteTask(scanner *bufio.Scanner, l []Task) []Task {
	for {
		// [UI UPDATE] English prompts & Cancel option (0)
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
		// [UI UPDATE] English prompts & Cancel option (0)
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
				l[id-1].IsDone = !l[id-1].IsDone // Toggle logic

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

	// [UI UPDATE] Cleaner Table Format
	fmt.Println("ID   STATUS   TASK")
	fmt.Println("--   ------   ----")
	for _, t := range l {
		status := "[ ]"
		if t.IsDone {
			status = "[X]"
		}
		// [UI UPDATE] Formatting: Left align ID (4 spaces), Status (8 spaces)
		fmt.Printf("%-4d %-8s %s\n", t.ID, status, t.Value)
	}
}
