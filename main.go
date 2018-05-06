package main

import (
	"fmt"
)

func main(){
	cmd1 := parseCmd()

	if cmd1.versionFlag{
		fmt.Println("version 0.0.1")
	}else if cmd1.helpFlag {
		printUsage()
	}else{
		startJVM(cmd1)
	}

}


