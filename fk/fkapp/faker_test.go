package fkapp_test

import (
	"fmt"
	"testing"

	"github.com/47monad/fakery/fk/fkapp"
)

func TestFaker(t *testing.T) {
	faker := fkapp.New()
	fmt.Println(faker.SemVer())
	fmt.Println(faker.Name())
}
