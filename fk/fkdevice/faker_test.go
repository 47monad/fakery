package fkdevice_test

import (
	"fmt"
	"testing"

	"github.com/47monad/fakery/fk/fkdevice"
)

func TestModel(t *testing.T) {
	faker := fkdevice.New()
	fmt.Println(faker.Model())
	fmt.Println(faker.Manufacturer())
	fmt.Println(faker.Platform())
	fmt.Println(faker.Type())
	fmt.Println(faker.SerialNumber())
}
