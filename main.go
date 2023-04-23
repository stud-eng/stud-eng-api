package main

import (
	"fmt"

	"github.com/stud-eng/stud-eng-api/config"
	"github.com/stud-eng/stud-eng-api/internal/app"
)

func main() {
	conf, err := config.New()
	if err != nil {
		fmt.Printf("conf error: %s", err)
	}

	if err := app.Run(conf); err != nil {
		fmt.Printf("app error: %s", err)
	}
}
