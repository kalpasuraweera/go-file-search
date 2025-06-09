package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Need at Least 2 arguments\nex: run /path text")
		os.Exit(0)
	}
	fmt.Println("Welocome to KP Search!")
	files, err := os.ReadDir(args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, file := range files {
		//fmt.Println("Proccessing: ", file.Name())
		if !file.IsDir() && file.Name() == args[2] {
			fmt.Println("Found file")
			os.Exit(0)
		}
		// TODO: add chanels and search deep inside directories
		/*
			For now we just search for file names
			content,err := os.Open(file.Name())
		*/
	}
	fmt.Println("No files found with the name")
}
