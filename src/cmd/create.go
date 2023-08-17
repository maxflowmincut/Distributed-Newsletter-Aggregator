package cmd

import (
	"bufio"
	"fmt"
	"net/mail"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"newsletter-aggregator/src/db/sqlite"
)

var reader = bufio.NewReader(os.Stdin)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		var name, email, sendTime string
		var err error

		// Prompt for Name
		for {
			fmt.Print("Enter Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			if name == "" {
				fmt.Println("Name cannot be empty!")
				continue
			}
			break
		}

		// Prompt for Email
		for {
			fmt.Print("Enter Email: ")
			email, err = reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				return
			}

			email = strings.TrimSpace(email)
			if !isValidEmail(email) {
				fmt.Println("Invalid email format!")
				continue
			}

			// Check email uniqueness
			exists, err := sqlite.EmailExists(DB, email)
			if err != nil {
				fmt.Printf("Failed to check email uniqueness: %v\n", err)
				return
			}
			if exists {
				fmt.Println("This email is already registered!")
				continue
			}
			break
		}

		// Prompt for Preferences
		categories := []string{"Tech", "Science", "Art"}
		var selectedCategories []string
		for {
			fmt.Println("Select Preferences (comma-separated, e.g., 1,2):")
			for i, cat := range categories {
				fmt.Printf("%d. %s\n", i+1, cat)
			}
			preferencesInput, _ := reader.ReadString('\n')
			preferencesInput = strings.TrimSpace(preferencesInput)

			// Process preferencesInput to get selected categories
			isValid, categoriesSelected := processPreferences(preferencesInput, len(categories))
			if !isValid {
				fmt.Println("Invalid preferences input!")
				continue
			}
			selectedCategories = categoriesSelected
			break
		}

		// Prompt for SendTime
		for {
			fmt.Print("Enter SendTime in GMT+0 (HH:00 format): ")
			sendTime, err = reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				return
			}
			sendTime = strings.TrimSpace(sendTime)

			if !isValidSendTime(sendTime) {
				fmt.Println("Invalid time format! Enter in HH:00 format.")
				continue
			}
			break
		}

		// Save user details
		joinedPreferences := strings.Join(selectedCategories, ", ")
		_, err = sqlite.CreateUser(DB, name, email, joinedPreferences, sendTime)
		if err != nil {
			fmt.Printf("Failed to create user: %v\n", err)
			return
		}
		fmt.Println("User created successfully!")
	},
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func processPreferences(input string, max int) (bool, []string) {
	split := strings.Split(input, ",")
	selected := make([]string, 0)
	for _, s := range split {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil || i < 1 || i > max {
			return false, nil
		}
		selected = append(selected, strconv.Itoa(i))
	}
	return true, selected
}

func isValidSendTime(time string) bool {
	if len(time) != 5 || time[2] != ':' {
		return false
	}
	hour, err := strconv.Atoi(time[:2])
	if err != nil || hour < 0 || hour > 23 || time[3:] != "00" {
		return false
	}
	return true
}

func init() {
    rootCmd.AddCommand(createCmd)
}