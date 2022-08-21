package main

import (
	"fmt"

	"github.com/projectdiscovery/nuclei/v2/pkg/templates"
)

func main() {
	fmt.Println("hello world")
	templates.GetTemplateDoc().Encode()
}
