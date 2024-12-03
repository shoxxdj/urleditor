package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	// Trim the newline character
	input = strings.TrimSpace(input)

	// Split input into URL and arguments
	parts := strings.Split(input, "?")
	baseURL := parts[0]
	fmt.Println("URL:", baseURL)

	var args []string
	if len(parts) > 1 {
		args = strings.Split(parts[1], "&")
		fmt.Println("Arguments:")
		for i, arg := range args {
			fmt.Printf("%d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("No arguments provided.")
	}

	for len(args) > 0 {
		fmt.Print("Enter parameter number to edit (or press Enter to finish): ")
		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		choice = strings.TrimSpace(choice)

		if choice == "" {
			break
		}

		index, err := strconv.Atoi(choice)
		if err != nil || index < 1 || index > len(args) {
			fmt.Println("Invalid choice. Please enter a valid parameter number.")
			continue
		}

		fmt.Printf("Current value: %s\n", args[index-1])
		fmt.Print("Enter new value: ")
		newValue, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		
		parts := strings.SplitN(args[index-1], "=", 2)
		if len(parts) == 2 {
			args[index-1] = parts[0] + "=" + strings.TrimSpace(newValue)
		} else {
			args[index-1] = strings.TrimSpace(newValue)
		}
	}

	if len(args) > 0 {
		updatedURL := baseURL + "?" + strings.Join(args, "&")
		fmt.Println("Updated URL:", updatedURL)
	} else {
		fmt.Println("Updated URL:", baseURL)
	}
}

