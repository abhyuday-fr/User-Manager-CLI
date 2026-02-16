package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct{
	Name  string `json:"name"`  // Tags tell Go how to name fields in JSON
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Global slice to store users in-memory
var users []User

// Reader to handle input with spaces (e.g., "John Doe")
var reader = bufio.NewReader(os.Stdin)

// The file where we save data
const fileName = "users.json"

func main(){

	// Load data immediately when the app starts
	loadUsers()

	fmt.Println("Running ...")
	createIntro()

	for {
		printMenu()
		choice := readLine("Enter your choice: ")

		switch choice {
		case "1":
			createUser()
		case "2":
			listUsers()
		case "3":
			searchUsers()
		case "4":
			deleteUser()
		case "5":
			fmt.Println("\nGoodbye! Thanks for using it :)")
			return
		default:
			fmt.Println("\nInvalid choice, please try again :(")
		}
		fmt.Println() // Add a blank line for readability
	}
}

func createIntro(){
	fmt.Println()
	fmt.Println("██╗   ██╗███████╗███████╗██████╗     ███╗   ███╗ █████╗ ███╗   ██╗ █████╗  ██████╗ ███████╗██████╗ ")
	fmt.Println("██║   ██║██╔════╝██╔════╝██╔══██╗    ████╗ ████║██╔══██╗████╗  ██║██╔══██╗██╔════╝ ██╔════╝██╔══██╗")  
	fmt.Println("██║   ██║███████╗█████╗  ██████╔╝    ██╔████╔██║███████║██╔██╗ ██║███████║██║  ███╗█████╗  ██████╔╝")  
	fmt.Println("██║   ██║╚════██║██╔══╝  ██╔══██╗    ██║╚██╔╝██║██╔══██║██║╚██╗██║██╔══██║██║   ██║██╔══╝  ██╔══██╗")  
	fmt.Println("╚██████╔╝███████║███████╗██║  ██║    ██║ ╚═╝ ██║██║  ██║██║ ╚████║██║  ██║╚██████╔╝███████╗██║  ██║")  
	fmt.Println(" ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝    ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝")
	fmt.Println()
	fmt.Println()
}

// --- Helper Functions ---

func printMenu() {
	fmt.Println("--------------------")
	fmt.Println("1. Create User")
	fmt.Println("2. List All Users")
	fmt.Println("3. Search User")
	fmt.Println("4. Delete User")
	fmt.Println("5. Exit")
	fmt.Println("--------------------")
}

func createUser() {
	fmt.Println("\n--- Create New User ---")
	
	name := readLine("Name: ")
	email := readLine("Email: ")
	
	// Handle Age Input (conversion string -> int)
	ageStr := readLine("Age: ")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("\nError: Age must be a number.")
		return
	}

	newUser := User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	// Append the new user to our in-memory slice
	users = append(users, newUser)
	fmt.Println("\nUser created successfully!")

	// Save to file immediately after adding
	saveUsers()
}

func listUsers() {
	fmt.Println("\n--- List of Users ---")
	if len(users) == 0 {
		fmt.Println("\nNo users found.")
		return
	}

	for i, u := range users {
        // 'i' is the index (0, 1, 2). 
        // We print 'i+1' so the user sees 1, 2, 3.
        fmt.Printf("%d. %s (%s, %d years old)\n", i+1, u.Name, u.Email, u.Age)
    }
}

func searchUsers() {
	fmt.Println("\n--- Search Users ---")
	query := strings.ToLower(readLine("\nEnter name to search: "))
	found := false

	for _, u := range users {
		// Check if the query is contained in the Name (case-insensitive)
		if strings.Contains(strings.ToLower(u.Name), query) {
			fmt.Printf("\nFound: %s | Email: %s | Age: %d\n", u.Name, u.Email, u.Age)
			found = true
		}
	}

	if !found {
		fmt.Println("\nNo users found matching that name.")
	}
}

func deleteUser() {
	fmt.Println("\n--- Delete User ---")
	// Show the list so the user knows which ID to pick
	listUsers() 
	if len(users) == 0 {
		return
	}

	idStr := readLine("Enter the ID of the user to delete: ")
	id, err := strconv.Atoi(idStr)
	
	// Validate input (User sees 1-based index, slice is 0-based)
	if err != nil || id < 1 || id > len(users) {
		fmt.Println("Invalid ID.")
		return
	}

	// Convert user-friendly ID (1, 2, 3) to slice index (0, 1, 2)
	index := id - 1

	// Capture the name for the success message before we delete it
	deletedName := users[index].Name

	// THE TRICK: Remove the element at 'index'
	// 1. Take everything up to the index: users[:index]
	// 2. Take everything after the index: users[index+1:]
	// 3. Join them together
	users = append(users[:index], users[index+1:]...)

	saveUsers() // Don't forget to save!
	fmt.Printf("Deleted user: %s\n", deletedName)
}

// readLine is a helper to clean up input reading
func readLine(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// --- File Handling Functions ---

// saveUsers converts the slice to JSON and writes it to a file
func saveUsers() {
	// MarshalIndent makes the JSON pretty and readable (adds spaces/tabs)
	data, err := json.MarshalIndent(&users, "", "  ")
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	// 0644 is the permission code (readable by everyone, writable by owner)
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// loadUsers reads the file and fills the 'users' slice
func loadUsers() {
	// Read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		// If the file doesn't exist, that's fine (first run). 
		// We just start with an empty list.
		return 
	}

	// Unmarshal parses the JSON data and puts it into the &users slice
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error parsing user data:", err)
	}
}