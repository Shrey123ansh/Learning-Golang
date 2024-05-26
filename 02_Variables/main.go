package main

import "fmt"

// uint8-  the set of all unsigned 8-bit integers (0 to 255)
// uint16-  set of all unsigned 16-bit integers (to 65535)
// uint32 -the the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64 -the set of all unsigned 64-bit integers (0 to 18446744073709551615)
// int8  -the set of all signed 8-bit integers (-128 to 127)
// int16- the set of all signed 16-bit integers (-32768 to 32767)
// int32- the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64- the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
// float32- the set of all IEEE-754 32-bit floating-point numbers
// float64 - the set of all IEEE-754 64-bit floating-point numbers
// complex64- the set of all complex numbers with float32 real and imaginary parts complex128 the set of all complex numbers with float64 real and imaginary parts
// byte alias- for uint8
// rune alias- for int32

//Please use documentation

// Public (if u put L capital in beginning) and global variable
const LoginToken string = "ghabbhhjd" // Public

func main() {
	//Can use global also
	var username string = "hitesh"
	// will add default /n in last of Println
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int8 = 127
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style

	//Can;t use global also
	numberOfUser := 300000.1243
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
