package learnlab

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bufioReadWrite() {
	bf_reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\n"))

	// Read byte slice
	byte_data := make([]byte, 15)
	n, err := bf_reader.Read(byte_data)
	if err != nil {
		fmt.Println("Error Reading:", err)
		return
	}	
	fmt.Printf("Read %d bytes: %s\n", n, byte_data[:n])

	// Read string until delimiter
	line, err := bf_reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}
	fmt.Println("Read string:", line)

	// Write byte slice
	byte_data = []byte("Sashiburidana, Mugiwara !!\n")
	bf_writer := bufio.NewWriter(os.Stdout)

	n, err = bf_writer.Write(byte_data)
	if err != nil {
		fmt.Println("Error writing byte data :", err)
		return
	}
	fmt.Printf("Wrote %d bytes :)\n", n)

	// Flush the buffer 
	err = bf_writer.Flush()
	if err != nil {
		fmt.Println("Error Flushing Writer :", err)
		return
	}

	// Writing string
	n, err = bf_writer.WriteString("Crocodile ... !!! Kaizoku Orewa naarrra")
	if err != nil {
		fmt.Println("Error writing string :", err)
	}
	fmt.Printf("Wrote %d bytes :)\n", n)

	// Flush the buffer 
	err = bf_writer.Flush()
	if err != nil {
		fmt.Println("Error Flushing Writer :", err)
		return
	}
}