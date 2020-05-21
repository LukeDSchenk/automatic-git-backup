package main

import (
    "os/exec"
    "log"
    "fmt"
    )

func main() {
    cmd := exec.Command("git", "status")
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", out)
    /*for _, b := range out {
        fmt.Printf("%q", b)
    }*/
}
