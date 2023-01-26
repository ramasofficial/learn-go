package console

import (
    "fmt"
    "strconv"
)

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

    fmt.Print(questions[menuSelection])
    selection, ok := askConsoleInput();

    if (!ok) {
        fmt.Printf("Bad input, input again\n");
        return generateQuestion(menuSelection, userInput)
    }

    selectionFloat, err := strconv.ParseFloat(selection, 64)

    if (err != nil) {
        fmt.Printf("Bad input, input again\n\n");
        return generateQuestion(menuSelection, userInput)
    }

    return selectionFloat
}
