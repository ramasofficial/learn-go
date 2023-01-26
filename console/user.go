package console

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
