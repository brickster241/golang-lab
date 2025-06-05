package learnlab

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() {
	// Open a file
	file, err := os.Open("file_read.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return
	}

	defer func() {
		fmt.Println("Closing File ....")
		file.Close()
	}()

	fmt.Println("File opened successfully....")
	
	// Read line by line
	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		line := scnr.Text()
		fmt.Println("Line :", line)
	}

	err = scnr.Err()
	if err != nil {
		fmt.Println("Error Reading File :", err)
		return
	}
}

func writeFile() {
	// os.Create function
	file, err := os.Create("file_byte_output.txt")
	if err != nil {
		fmt.Println("Error creating File :", err)
		return
	}
	defer file.Close()

	// file.Write function -> Writing byte data
	data := []byte("Hello Golang :) \n\nYou are indeed better than Python.\n\n")
	_, err = file.Write(data)
	fmt.Println("Byte Data written to file successfully...")

	file, err = os.Create("file_st_output.txt")
	if err != nil {
		fmt.Println("Error creating File :", err)
		return
	}
	defer file.Close()

	// file.WriteString function -> Writing string data
	st_data := "Hi Everyone, bye everyone. Golang here. \n\n\n\n\nHi, I am Python."
	_, err = file.WriteString(st_data)
	fmt.Println("String Data written to file successfully...")

}