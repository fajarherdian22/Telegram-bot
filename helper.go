package main

func isError(err error) {
	if err != nil {
		panic(err)
	}
}
