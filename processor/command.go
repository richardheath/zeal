package processor

import (
	"fmt"

	"github.com/richardheath/zeal/config"
)

func ExecuteCommand(command string, targetPackage string) error {
	fmt.Print(targetPackage)
	fmt.Println("In ExecuteCommand")
	configData, err := config.LoadDir(targetPackage)

	filter := configData.Filter("install", "contents", "")
	fmt.Println(filter)
	if err != nil {
		return err
	}
	return nil
}
