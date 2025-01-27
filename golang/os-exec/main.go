package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func basic() {
	// create command for looping number from 0 to 10 and sleep each iteration for 2s
	cmd := exec.Command("bash", "-c", "bin/loop")

	stdin, _ := cmd.StdinPipe()
	defer stdin.Close()

	// connect to standart output command
	// will closing the pipe after seeing the command exit
	stdout, _ := cmd.StdoutPipe()
	defer stdout.Close()

	// start new subprocess
	fmt.Println("start")
	if err := cmd.Start(); err != nil {
		fmt.Println("error: ", err)
	}

	// listening output from "stdout"
	reader := bufio.NewReader(stdout)
	for {
		out, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(out)
		fmt.Println(line)

		// do something when the output is 3
		if line == "3" {
			// given input to subprocess
			// '\n' is equal to pressing the enter. without this the subproccess will waiting forever
			io.WriteString(stdin, "test\n")
			// kill subprocess
			cmd.Process.Kill()
		}
	}

	// wait for the subproccess exit
	if err := cmd.Wait(); err != nil {
		fmt.Println("error during waiting the process: ", err)
	}

	if cmd.ProcessState.Success() {
		fmt.Println("done")
	} else {
		fmt.Println("failed")
	}
}
