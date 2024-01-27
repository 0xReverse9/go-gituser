package logger

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintManual() {
	fmt.Println("Hi there 👋🏼")
	fmt.Println("Welcome to the manual 📘")
	fmt.Println("")
	color.Cyan("Description:")
	fmt.Println(" This program helps switch between different git user accounts easily")
	fmt.Println(" There is 3 modes for this program")
	fmt.Println(" 🏠 | <personal> for a personal account \n 📚 | <school> for a school account \n 💻 | <work> for a professional account ")
	fmt.Println("")
	color.Cyan("Setup:")
	fmt.Println(" To add your accounts you need to run")
	fmt.Println(" gituser setup")
	fmt.Println("")
	color.Cyan("Usage:")
	fmt.Println(" To use the program you just need to call the executable")
	fmt.Println(" gituser <mode>")
	fmt.Println("")
	color.Cyan("Flags:")
	fmt.Println(" gituser help (Flags Usages)")
	fmt.Println(" gituser manual (Prints Manual)")
	fmt.Println(" gituser info (Prints Accounts)")
	fmt.Println(" gituser now (Prints Current Account)")
	fmt.Println("")
}
