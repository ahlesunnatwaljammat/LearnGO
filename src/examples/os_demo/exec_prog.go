package main

import (
	"os"
	"os/exec"
)

func main() {
	ping := exec.Command("ping", "-t", "noman", "-n", "5")
	/*ping := exec.Command("cmd")
	ping.Stdin = strings.NewReader("dir")*/
	ping.Stdout = os.Stdout
	ping.Stderr = os.Stderr
	ping.Run()
}
