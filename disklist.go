package main

import(
    "fmt"
    "os/exec"
)

func main(){ 

cmd := exec.Command("cmd", "/c","wmic diskdrive list brief")
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("cmd.Run() failed with %s\n", err)
    }
    fmt.Printf("combined out:\n%s\n", string(out))}