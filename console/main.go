package console

import (
    "fmt"
)

func Execute(userInput UserInput) {
    menuSelection, ok := showMenu()

    if (menuSelection == 5) {
        fmt.Print("Terminal closed.\nGood Bye!\n\n")
        return
    }

    if (!ok) {
        fmt.Print("\nBad input, please select correct input from menu!\n")
        Execute(userInput)
        return
    }

    answer := generateQuestion(menuSelection, userInput)
    userInput = setUserInput(menuSelection, userInput, answer)

    Execute(userInput)
    return
}
