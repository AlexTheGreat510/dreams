package main

/* import modules */

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
    waitType := flag.String("wait-type", "minutes", "Set the wait type [seconds, minutes, hours].")
    
    /* parse flags */

    flag.Parse()
    
    /* check the time type */
    
    var timeType time.Duration

    switch *waitType {

        case "seconds":
            timeType = time.Second

        case "minutes":
            timeType = time.Minute

        case "hours":
            timeType = time.Hour

        default:
            timeType = time.Minute
    }

    /* start the infinite loop */

    for true {

        /* show the message if running command */

        fmt.Printf("dreams: Running command '%s'\n", *command)

        /* let program execute the program */

        cmd := exec.Command(*command)
        err := cmd.Run()

        if err != nil {
            fmt.Printf("dreams: Error executing command: %s\n", err)
        }
        
        /* let program sleep */

        time.Sleep(time.Duration(*wait) * timeType)
    }
}
