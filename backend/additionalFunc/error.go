package additionalFunc

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
