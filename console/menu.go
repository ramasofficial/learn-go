package console

import (
    "fmt"
    "strconv"
)

func showMenu() (int, bool) {
    fmt.Printf("\nSelect menu:\n\n%s", getMenuOptions())
    selection, ok := askConsoleInput();

    if (!ok) {
        return 0, false
    }

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
