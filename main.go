package main 

import(
	"fmt"
	"github.com/AlecAivazis/survey"
)

func askName() (string, error){
	name := ""
	prompt := &survey.Input{
		Message: "What is your name?",
	}
	if err := survey.AskOne(prompt, &name); err != nil {
		return name, err
	}
	return name, nil
}

func main() {
	fmt.Println("hello world")
	name, err := askName()
	if err != nil {
		fmt.Println("error occured: ", err.Error())
		return
	}
	fmt.Printf("Hello, %s.\n", name)
}