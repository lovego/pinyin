package utils

import "fmt"

func ExampleInitials() {
	fmt.Println(Initials("长大"))
	fmt.Println(Initials("长城abc"))
	fmt.Println(Initials(" a长b大,c "))
	// Output:
	// ZD
	// ZCABC
	// AZBDC
}
