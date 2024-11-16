package main

import "DB_Project/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
