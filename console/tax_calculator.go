package console

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
