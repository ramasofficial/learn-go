package console

import (
    "fmt"
    "strconv"
    "bufio"
    "os"
    "strings"
)

type UserInput struct {
    YearlyIncome, AdditionalIncome, TaxExemption float64
}

func Execute(userInput UserInput) {
    menuSelection, ok := showMenu()

    if (menuSelection == 5) {
        fmt.Println("Terminal closed.\nGood Bye!\n\n")
        return
    }

    if (!ok) {
        fmt.Println("\nBad input, please select correct input from menu!\n")
        Execute(userInput)
        return
    }

    answer := generateQuestion(menuSelection, userInput)
    userInput = setUserInput(menuSelection, userInput, answer)

    Execute(userInput)
    return
}

func setUserInput(menuSelection int, userInput UserInput, answer float64) UserInput {
    switch menuSelection {
    case 1:
        userInput.YearlyIncome = answer
    case 2:
        userInput.TaxExemption = answer
    case 3:
        userInput.AdditionalIncome = answer
    }

    return userInput
}

const SMALLER_TAXES_THRESHOLD int = 30000
const SMALLER_TAXES_MULTIPLIER float64 = 0.20
const HIGHER_TAXES_MULTIPLIER float64 = 0.25

func calculateTax(userInput UserInput) float64 {
    totalIncome := (userInput.YearlyIncome + userInput.AdditionalIncome) - userInput.TaxExemption

    if (totalIncome < float64(SMALLER_TAXES_THRESHOLD)) {
        return totalIncome * SMALLER_TAXES_MULTIPLIER
    }
    
    return totalIncome * HIGHER_TAXES_MULTIPLIER
}

const CALCULATE_TAXES_OPTION int = 4

func generateQuestion(menuSelection int, userInput UserInput) float64 {
    if (menuSelection == CALCULATE_TAXES_OPTION) {
        taxExemption := calculateTax(userInput)
        fmt.Printf("\nTax: %f\n\n", taxExemption)
        return 0
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
        fmt.Printf("Bad input, input again\n");
        return generateQuestion(menuSelection, userInput)
    }

    return selection
}

func showMenu() (int, bool) {
    fmt.Printf("\nSelect menu:\n\n%s", getMenuOptions())

    reader := bufio.NewReader(os.Stdin)
    selection, err := reader.ReadString('\n')

    if (err != nil) {
        return 0, false
    }

    selection = strings.TrimSpace(selection)
    selectionInteger, err := strconv.Atoi(selection)

    if (err != nil) {
        return 0, false
    }

    return selectionInteger, true
}

func getMenuOptions() string {
    var menuOptions = [...]string {
        "Input your yearly income (salary based)",
        "Input the amount of tax exemption (if any)",
        "Input your additional income",
        "Calculate tax",
        "Quit",
    }
    
    var menu string
    for key, menuOption := range menuOptions {
        key++
        menu = menu + strconv.Itoa(key) + ") " + menuOption + "\n"
    }

    return menu
}
