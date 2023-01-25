package console

import (
    "fmt"
    "strconv"
)

type UserInput struct {
    YearlyIncome, TaxExemption, AdditionalIncome float64
}

func Execute(userInput UserInput) {
    menuSelection, ok := ShowMenu()

    if (menuSelection == 5) {
        fmt.Println("Terminal closed.\nGood Bye!")
        return
    }

    if (!ok) {
        Execute(userInput)
        return
    }

    answer := GenerateQuestion(menuSelection, userInput)
    userInput = SetUserInput(menuSelection, userInput, answer)

    Execute(userInput)
    return
}

func SetUserInput(menuSelection int, userInput UserInput, answer float64) UserInput {
    if (menuSelection == 1) {
        userInput.YearlyIncome = answer
    }

    if (menuSelection == 2) {
        userInput.TaxExemption = answer
    }

    if (menuSelection == 3) {
        userInput.AdditionalIncome = answer
    }

    return userInput
}

func calculateTax(userInput UserInput) {
    var totalIncome float64
    totalIncome = userInput.YearlyIncome + userInput.AdditionalIncome

    fmt.Printf("\nTax: %f\n\n", totalIncome)
}

func GenerateQuestion(menuSelection int, userInput UserInput) float64 {
    if (menuSelection == 4) {
        calculateTax(userInput)
        return 100.00
    }

    questions := map[int]string {
        1: "Input yearly income: ",
        2: "Input the amount of tax exemption: ",
        3: "Input your additional income: ",
    }

    var selection float64

    fmt.Printf(questions[menuSelection] + "\n")
    _, err := fmt.Scanln(&selection)

    if (err != nil) {
        fmt.Printf("Bad input, input gain\n");
        return GenerateQuestion(menuSelection, userInput)
    }

    return selection
}

func ShowMenu() (int, bool) {
    fmt.Printf("Select menu:\n%s", GetMenuOptions())

    var selection int

    _, err := fmt.Scanln(&selection)

    if (err != nil) {
        return 0, false
    }

    return selection, true
}

func GetMenuOptions() string {
    // menuOptions := map[int]string {
    //     1: "Input your yearly income (salary based)",
    //     2: "Input the amount of tax exemption (if any)",
    //     3: "Input your additional income",
    //     4: "Calculate tax",
    //     5: "Quit",
    // }

    var menuOptions [5]string
    menuOptions[0] = "Input your yearly income (salary based)";
    menuOptions[1] = "Input the amount of tax exemption (if any)";
    menuOptions[2] = "Input your additional income";
    menuOptions[3] = "Calculate tax";
    menuOptions[4] = "Quit";
    
    var menu string
    for key, menuOption := range menuOptions {
        menu = menu + strconv.Itoa(key + 1) + ") " + menuOption + "\n"
    }

    return menu
}
