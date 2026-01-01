package main

import (
	"io"
	"os"

	"go-rest-api/internal/model"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&model.User{},
	)
	if err != nil {
		panic(err)
	}
	io.WriteString(os.Stdout, stmts)
}
