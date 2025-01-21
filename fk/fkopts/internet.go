package fkopts

// Domain Options
type DomainOpts struct{}

type DomainBuilder struct {
	Opts []func(*DomainOpts) error
}

func Domain() *DomainBuilder {
	return &DomainBuilder{}
}

func (b *DomainBuilder) List() []func(*DomainOpts) error {
	return b.Opts
}

// IPv4 Options
type IPv4Opts struct{}

type IPv4Builder struct {
	Opts []func(*IPv4Opts) error
}

func IPv4() *IPv4Builder {
	return &IPv4Builder{}
}

func (b *IPv4Builder) List() []func(*IPv4Opts) error {
	return b.Opts
}

// IPv6 Options
type IPv6Opts struct{}

type IPv6Builder struct {
	Opts []func(*IPv6Opts) error
}

func IPv6() *IPv6Builder {
	return &IPv6Builder{}
}

func (b *IPv6Builder) List() []func(*IPv6Opts) error {
	return b.Opts
}

// Mac Address Options
type MacAddrOpts struct{}

type MacAddrBuilder struct {
	Opts []func(*MacAddrOpts) error
}

func MacAddr() *MacAddrBuilder {
	return &MacAddrBuilder{}
}

func (b *MacAddrBuilder) List() []func(*MacAddrOpts) error {
	return b.Opts
}

// Email Options
type EmailOpts struct{}

type EmailBuilder struct {
	Opts []func(*EmailOpts) error
}

func Email() *EmailBuilder {
	return &EmailBuilder{}
}

func (b *EmailBuilder) List() []func(*EmailOpts) error {
	return b.Opts
}

// Username Options
type UsernameOpts struct{}

type UsernameBuilder struct {
	Opts []func(*UsernameOpts) error
}

func Username() *UsernameBuilder {
	return &UsernameBuilder{}
}

func (b *UsernameBuilder) List() []func(*UsernameOpts) error {
	return b.Opts
}

// Browser Options
type BrowserOpts struct{}

type BrowserBuilder struct {
	Opts []func(*BrowserOpts) error
}

func Browser() *BrowserBuilder {
	return &BrowserBuilder{}
}

func (b *BrowserBuilder) List() []func(*BrowserOpts) error {
	return b.Opts
}

// Browser Engine Options
type BrowserEngineOpts struct{}

type BrowserEngineBuilder struct {
	Opts []func(*BrowserEngineOpts) error
}

func BrowserEngine() *BrowserEngineBuilder {
	return &BrowserEngineBuilder{}
}

func (b *BrowserEngineBuilder) List() []func(*BrowserEngineOpts) error {
	return b.Opts
}
