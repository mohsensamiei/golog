package main

import (
	"github.com/mohsensamiei/golog"
	"github.com/pkg/errors"
)

type TestData struct {
	Name string
}

func main() {
	log.SetLevel(log.DebugLevel)

	log.Data("name", "mohsen").Debug("error message")

	log.With("mohsen").Info("info message")

	log.With(TestData{Name: "mohsen"}).Warning("warning message")

	log.With(errors.New("error message")).Error("error message")

	log.Fatal("fatal error")
}
