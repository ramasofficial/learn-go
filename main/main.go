package main

import (
    "ramasdev/console"
)

func main() {
    console.Execute(console.UserInput{YearlyIncome:1000, AdditionalIncome: 5000, TaxExemption:0})
}
