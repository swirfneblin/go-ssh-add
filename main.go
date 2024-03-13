package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    scriptPath := "./scripts/add_ssh_keys.sh"

    if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
        fmt.Println("The selected script doesn't exist.")
        return
    }

    cmd := exec.Command("bash", scriptPath)

    stdoutStderr, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("The script failed:", err)
        return
    }

    fmt.Println(string(stdoutStderr))
}
