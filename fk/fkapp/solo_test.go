package fkapp_test

import (
	"fmt"
	"testing"

	"github.com/47monad/fakery/fk/fkapp"
)

func TestName(t *testing.T) {
	fmt.Println(fkapp.Name())
}

func TestSemVer(t *testing.T) {
	fmt.Println(fkapp.SemVer())
}
