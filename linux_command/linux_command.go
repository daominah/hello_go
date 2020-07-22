package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	//cmd := exec.Command("ls", "-l")
	cmd := exec.Command("pwd")
	//cmd := exec.Command("fuck", "-l")

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout, cmd.Stderr = &outBuf, &errBuf
	err := cmd.Run()
	fmt.Println("ERROR cmd.Run: ", err)
	fmt.Println("Stderr: ", errBuf.String())
	fmt.Println("Stdout:", outBuf.String())
}
