package main

import (
	"fmt"

	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
	"context"
	"time"
)
func main() {
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()
	done := make(chan struct{})
	go func ()  {
		stdout, stderr, err := shell.Execute("Start-Sleep -Seconds 10; Get-WmiObject -Class Win32_Processor")
		if err != nil {
			panic(err)
		}

		fmt.Println(stdout)
		fmt.Println(stderr)
		fmt.Println(err)
		close(done)
	}()
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
	select {
		case <-ctx.Done():
			fmt.Println("context done")
			shell.Exit()
			fmt.Println("exiting")
		case <-done:
			fmt.Println("done")
	}
	
}