package fkinternet

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
	"golang.org/x/text/language"
)

func Domain(opts ...fkopts.Builder[fkopts.DomainOpts]) string {
	//_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext[fkdata.Internet](language.English, nil).SetBinder(fkdata.NewInternet)
	return domainName(ctx)
}

func IPv4(opts ...fkopts.Builder[fkopts.IPv4Opts]) string {
	return ipv4Address()
}

func IPv6(opts ...fkopts.Builder[fkopts.IPv6Opts]) string {
	return ipv6Address()
}

func MacAddr(opts ...fkopts.Builder[fkopts.MacAddrOpts]) string {
	return macAddress()
}

func Email(opts ...fkopts.Builder[fkopts.EmailOpts]) string {
	ctx := fk.NewContext[fkdata.Internet](language.English, nil).SetBinder(fkdata.NewInternet)
	return email(ctx)
}

func Username(opts ...fkopts.Builder[fkopts.UsernameOpts]) string {
	ctx := fk.NewContext[fkdata.Internet](language.English, nil).SetBinder(fkdata.NewInternet)
	return username(ctx)
}

func Browser(opts ...fkopts.Builder[fkopts.BrowserOpts]) string {
	ctx := fk.NewContext[fkdata.Internet](language.English, nil).SetBinder(fkdata.NewInternet)
	return browser(ctx)
}

func BrowserEngine(opts ...fkopts.Builder[fkopts.BrowserEngineOpts]) string {
	ctx := fk.NewContext[fkdata.Internet](language.English, nil).SetBinder(fkdata.NewInternet)
	return browserEngine(ctx)
}
