package main

import (
	"log"
	"os"
	"os/exec"
)

// Shorted commands to execute.
var cmdChain = []*exec.Cmd{
	exec.Command("lib/synonyms"),
	exec.Command("lib/sprinkle"),
	exec.Command("lib/coolify"),
	exec.Command("lib/domainify"),
	exec.Command("lib/available"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout

	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Fatalln(err)
		}

		nextCmd.Stdin = stdout
	}

	for _, cmd := range cmdChain {

		// Start the processes in background
		if err := cmd.Start(); err != nil {
			log.Fatalln(err)

			// And be sure to kill the processes when this program is finished
		} else {
			defer cmd.Process.Kill()
		}
	}

	for _, cmd := range cmdChain {
		// Wait for the processes to finish
		if err := cmd.Wait(); err != nil {
			log.Fatalln(err)
		}
	}
}
