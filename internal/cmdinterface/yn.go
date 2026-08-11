package cmdinterface

import (
	"fmt"

	"github.com/fatih/color"
)

// YesOrNoWithColor shows y/N or Y/n question with color
func YesOrNoWithColor(
	mainColor *color.Color,
	question string,
	defaultChoice bool,
) (bool, error) {
	_, _ = mainColor.Print(question)
	fmt.Print(" (")
	if defaultChoice {
		colorTrue := color.New(color.FgGreen).Add(color.Bold)
		_, _ = colorTrue.Print("Y")
	} else {
		colorTrue := color.New(color.FgGreen)
		_, _ = colorTrue.Print("y")
	}
	fmt.Print("/")
	if !defaultChoice {
		colorTrue := color.New(color.FgRed).Add(color.Bold)
		_, _ = colorTrue.Print("N")
	} else {
		colorTrue := color.New(color.FgRed)
		_, _ = colorTrue.Print("n")
	}
	fmt.Println(") ")
	var response string
	_, errInput := fmt.Scan(&response)
	if errInput != nil {
		return defaultChoice, errInput
	}
	switch response {
	case "y":
		return true, nil
	case "Y":
		return true, nil
	case "n":
		return false, nil
	case "N":
		return false, nil
	default:
		return defaultChoice, nil
	}
}
