package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		input = strings.TrimSuffix(input, "\n")

		if err = execute(input); err!=nil{
			fmt.Println(os.Stderr, err)
		}

	}
}

var ErrNoPath = errors.New("path required")

func execute(input string) error{

	store := strings.Split(input, " ")

	switch store[0]{
	case "cd":
		if len(store)<2{
			return ErrNoPath
		}else{
			return os.Chdir(store[1])
		}
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(store[0], store[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}