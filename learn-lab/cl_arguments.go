package learnlab

import (
	"flag"
	"fmt"
	"os"
)

func commandArgs() {

	// Print command line arguments
	cl_args := os.Args
	fmt.Println("Command Arguments : ", cl_args)

	// Define flags
	var name string
	var age int
	var male bool
	flag.StringVar(&name, "name", "John", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", true, "Is user male ?")

	flag.Parse()
	fmt.Println(name, age, male)
}

func commandSubCommands() {

	// command line subcommands
	stringFlag := flag.String("user", "Guest", "Name of the user")
	flag.Parse()
	fmt.Println(stringFlag)
	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subcommand1.Bool("processing", false, "Command processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subcommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language:", *flagsc2)
	default:
		fmt.Println("no subcommand entered!")
		os.Exit(1)
	}
}