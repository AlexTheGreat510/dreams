package main

import (
    "fmt"
    "flag"
    "time"
    "os/exec"
)

/* default flag variables */

var (
    wait *int
    command *string
)

/* parse arguments using flag */

func init() {
    wait = flag.Int("wait", 1, "Set the execution time.")
    command = flag.String("command", "random-wallpaper", "Set the command.")
}

/* function to be run by program */

func main() {
    
    /* parse flags */
    
    flag.Parse()
    
    /* start the infinite loop */
    
    for true {

        /* let program sleep */
    
        time.Sleep(time.Duration(*wait) * time.Minute)
        
        /* let program execute the program */

        cmd := exec.Command(*command)
        err := cmd.Run()

        if err != nil {
            fmt.Printf("dreams: Error executing command: %s\n", err)
        }
    }
}
