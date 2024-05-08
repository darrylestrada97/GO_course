package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Bank!") // Print welcome message2
	balance := float64(0)
Loop:
	for {
		fmt.Println("1. Check your balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		var choice int
		fmt.Print("Please enter your choice:")
		_, err := fmt.Scan(&choice)
		if err != nil {
			return
		}

		switch choice {
		case 1:
			fmt.Printf("Your balance is %v\n", balance)
		case 2:
			fmt.Print("Please enter the amount you want to deposit:")
			var deposit float64
			_, err := fmt.Scan(&deposit)
			if err != nil {
				return
			}

			if deposit < 0 {
				fmt.Println("Invalid amount, Must be greater than 0")
				return
			}
			balance += deposit
			fmt.Println("You have deposited:", deposit)
			fmt.Println("Your balance is:", balance)
		case 3:
			fmt.Print("Please enter the amount you want to withdraw:")
			var withdraw float64
			_, err = fmt.Scan(&withdraw)
			if err != nil {
				return
			}
			if withdraw < 0 {
				fmt.Println("Invalid amount, Must be greater than 0")
				return
			}
			if withdraw > balance {
				fmt.Println("You don't have enough money to withdraw")
			} else {
				balance -= withdraw
				fmt.Println("You have withdrawn:", withdraw)
				fmt.Println("Your balance:", balance)
			}
		case 4:
			fmt.Println("Thank you for using our service!")
			break Loop
		default:
			fmt.Println("Invalid choice")
		}
	}
}
