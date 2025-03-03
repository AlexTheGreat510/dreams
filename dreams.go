package main

import (
    "fmt"
    "flag"
    "time"
    "os/exec"
)

/* function to be run by program */

func main() {
 
    /* parse arguments using flag */

    wait := flag.Int("wait", 1, "Set the execution time.")
    command := flag.String("command", "random-wallpaper", "Set the command.")
    
    /* parse flags */
    
    flag.Parse()
    
    /* start the infinite loop */
    
    for true {

        /* let program execute the program */

        cmd := exec.Command(*command)
        err := cmd.Run()

        if err != nil {
            fmt.Printf("dreams: Error executing command: %s\n", err)
        }
        
        /* let program sleep */

        time.Sleep(time.Duration(*wait) * time.Minute)

        /* show the message if running command */

        fmt.Println("dreams: Running command", *command)
    }
}
