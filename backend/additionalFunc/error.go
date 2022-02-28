package additionalFunc

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)
}

func CheckErr(err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
