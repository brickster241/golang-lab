package main

import (
	"fmt"
)

func main() {
	/*
		Integer Formatting Verbs
	*/
	value := 22

	// Base 2 Output
	fmt.Printf("%b\n", value)
	
	// Base 10 Output
	fmt.Printf("%d\n", value)

	// Base 10 but show sign
	fmt.Printf("%+d\n", value)

	// Base 8 Output
	fmt.Printf("%o\n", value)

	// Base 8 with leading 0x
	fmt.Printf("%O\n", value)

	// Base 16 lowercase Output
	fmt.Printf("%x\n", value)

	// Base 16 uppercase Output
	fmt.Printf("%X\n", value)

	// Base 16 with leading 0x
	fmt.Printf("%#x\n", value)

	// Pad with Spaces right justified : width 4
	fmt.Printf("%4d\n", value)

	// Pad with Spaces left justified : width 4
	fmt.Printf("%-4d\n", value)

	// Pad with zeroes : width 4
	fmt.Printf("%04d\n", value)

	/*
		String Formatting Verbs
	*/ 
	txt := "Hello, World"

	// Prints text as plain string
	fmt.Printf("%s\n", txt)

	// Prints text as double-quoted string
	fmt.Printf("%q\n", txt)

	// Right Justified : Width 20
	fmt.Printf("%20s\n", txt)

	// Left Justified : Width 20
	fmt.Printf("%-20s\n", txt)

	// Prints text as hex dump of byte values
	fmt.Printf("%x\n", txt)

	// Prints text as hex dump with spaces
	fmt.Printf("% x\n", txt)

	/*
		Float Formatting Verbs
	*/
	flt := 918.5673

	// Scientific Notation with 'e' as exponent
	fmt.Printf("%e\n", flt)

	// Decimal Point, no exponent
	fmt.Printf("%f\n", flt)

	// Default Width, Precision 2
	fmt.Printf("%.2f\n", flt)

	// Minimum Width, Precision 2
	fmt.Printf("%6.2f\n", flt)

	// Exponent as Needed, only necessary digits
	fmt.Printf("%g\n", flt)

}