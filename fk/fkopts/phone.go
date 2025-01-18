package fkopts

import "github.com/47monad/fakery/fk/fkdata"

func PhoneNumber() *GeneralBuilder[fkdata.PhoneNumber] {
	return &GeneralBuilder[fkdata.PhoneNumber]{lw: &LangWrapper{}}
}

func MobileNumber() *GeneralBuilder[fkdata.PhoneNumber] {
	return &GeneralBuilder[fkdata.PhoneNumber]{lw: &LangWrapper{}}
}
