package main

import "fmt"

func isError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
