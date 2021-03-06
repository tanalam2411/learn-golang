package main

import (
	"fmt"
)

/*
Agenda - Primitives type
- Boolean type
- Numeric types
  - Integers
  - Floating point
  - Complex numbers
- Text types

For more checkout - https://golang.org/pkg/builtin/
*/

func main() {

	// Boolean types - true/false
	// Boolean is not an alias of any numeric or other type. So you can't convert to its equivalent other type like how you do in Python
	// https://golang.org/ref/spec#Boolean_types
	fmt.Println("**************** BOOLEAN ********************")

	var b bool = true

	fmt.Printf("%v, %T \n", b, b)

	n := 10 == 10
	m := 11 == 12

	fmt.Printf("%v, %T \n", n, n)
	fmt.Printf("%v, %T \n", m, m)

	//Default value for bool is false
	var flag bool
	fmt.Printf("%v, %T \n", flag, flag)

	fmt.Println("**************** INT ********************")

	// numeric type
	// https://golang.org/ref/spec#Numeric_types
	// Integer type (Signed numbers are both positive nad negative, Unsigned numbers are always positive numbers)
	// int range:
	// int 		-   min 32bit
	// int8		-	-128	    					+127
	// int16	- 	-32768							+32767
	// int32	- 	-2 147 483 648					+2 147 483 647          	// Approx neg to pos 2 Billion
	// int64	- 	-9 223 372 036 854 775 808		+9 223 372 036 854 775 807  // Approx neg to pos 9 Quintillion

	// Also have unsigned int - uint, uint8, uint16, uint32, uint64, uintptr
	// uint8	-	0 - 255
	// uint16	-	0 - 65 535
	// uint32	-	0 - 4 294 967 295
	// uint64	-

	// byte - is a unsigned 8bit int type (alias for uint8)
	// rune - is a alias for int32(represents a unicode point)

	// Operation on same type of different range will through erro - int + int8 not possible,
	// so in this senario explicit type cast is required int + int(int8)

	var n1 int // default
	var n2 int8 = 11

	println(n1 + int(n2)) // as n1 + n2 will through error (invalid operation: n1 + n2 (mismatched types int and int8))

	// Bit operator
	b1 := 10              // 1010
	b2 := 3               // 0011
	fmt.Println(b1 & b2)  // &(and) when both has 1 it sets 1, so 1010 & 0011 - 0010
	fmt.Println(b1 | b2)  // |(or) when has atleast one 1s , so 1010 | 0011 - 1011
	fmt.Println(b1 ^ b2)  // ^(exclusive or - xor) only one should have one 1, so 1010 ^ 0011 - 1001
	fmt.Println(b1 &^ b2) // &^(and not) if both has 0, so 1010 &^ 0011 - 0100

	// Bit shifting
	b3 := 8              // 2^3
	fmt.Println(b3 << 3) // left shifting -> 2^3 * 2^3 = 2^6
	fmt.Println(b3 >> 3) // right shifting -> 2^3 / 2^3 = 2^0

	fmt.Println("**************** FLOAT ********************")

	// Floating point number
	// float32 = +or-1.18E-38 to +or-3.4E38  //E-38 -> 10^-3 and E38 -> 10^3, so 1.18E-38 = 1.18*10^-38
	// float64 = +or-2.23E-308 to +or-1.8E308
	// you can't perform % operation, bit operator and bit shifting on floating numbers

	f1 := 3.14   // var f1 float32 = 3.14, f1 := 3.14 will assign type float64 as default
	f1 = 13.7e72 //13.7*10^72
	f1 = 2.1e14
	fmt.Printf("%v, %T \n", f1, f1)

	fmt.Println("**************** COMPLEX ********************")

	// Complex numbers
	// Zero value of complex number is = 0 + 0i
	// complex64 = float32(real) + float32(imag)
	// complex128 = float64(real) + float64(imag)
	// Builtin functions - complex, real, imag

	var c1 complex64 = 1 + 2i
	fmt.Printf("%v, %T \n", c1, c1)
	fmt.Printf("%v, %T \n", real(c1), real(c1))
	fmt.Printf("%v, %T \n", imag(c1), imag(c1))

	c2 := 1 + 2i
	c3 := 2 + 5.2i
	fmt.Println(c2 + c3)
	fmt.Println(c2 - c3)
	fmt.Println(c2 * c3)
	fmt.Println(c2 / c3)

	// Create complex number
	var c4 complex64 = complex(10, 20)
	fmt.Printf("%v, %T \n", c4, c4)

	fmt.Println("**************** TEXT ********************")

	// https://golang.org/ref/spec#String_types
	// Text types :
	// 1) String - Any utf-8 character
	// String in go are aliases of bytes
	// String are immutable

	s1 := "First"
	s2 := "Second"
	fmt.Printf("%v, %T\n", s1, s1)
	fmt.Printf("%v, %T\n", s1[3], s1[3])         // s1[3] gives ascii val of that char
	fmt.Printf("%v, %T\n", string(s1[3]), s1[3]) // to get string val, call string func

	fmt.Printf("%v, %T\n", s1+s2, s1+s2)

	by1 := []byte(s1) //convert string to byte
	fmt.Printf("%v, %T \n", by1, by1)

	// 2) rune - Any utf-32 char
	// rune is a type alias for int32
	// single quoiting creates rune type, r1 := 'a' will create r1 of type rune(int32)
	// look strings.Reader#ReadRune

	r1 := 'a'
	var r2 rune = 'a'
	fmt.Printf("%v, %T\n", r1, r1)
	fmt.Printf("%v, %T\n", r2, r2)

}
