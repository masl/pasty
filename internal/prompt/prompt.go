package prompt

import (
	"log"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

// Load loads the prompt
func Load() {
	for {
		command := ""
		prompt := &survey.Input{}
		icons := survey.WithIcons(func(icons *survey.IconSet) {
			icons.Question.Text = ">"
		})

		err := survey.AskOne(prompt, &command, icons)

		if err != nil {
			if err == terminal.InterruptErr {
				return 
			}
			log.Println(err.Error())
		}

		fmt.Println(":: " + command)
	}
}