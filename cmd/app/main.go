package main

import (
	"github.com/dis0ps/golang-project-template/internal/helpers"
	"github.com/dis0ps/golang-project-template/pkg/apitools"
)

func main() {
	const url = "https://jsonplaceholder.typicode.com/todos/1"
	var response string = apitools.FetchAPI(url)
	helpers.PrintMessage(response)

}
