package console

import (
    "bufio"
    "os"
    "strings"
)

func askConsoleInput() (string, bool) {
    reader := bufio.NewReader(os.Stdin)
    selection, err := reader.ReadString('\n')

    if (err != nil) {
        return "", false
    }

    return strings.TrimSpace(selection), true
}
