package person

import "github.com/47monad/fakery/fk"

// Tier-3 package-level generators. Each one builds a throwaway Faker from
// the supplied options and delegates to it, so callers get zero-setup usage
// with no shared state:
//
//	person.FirstName()
//	person.LastName(person.WithLang("fa"))
//	person.Name(person.WithGender(person.Female))
//
// For a reproducible stream, construct a seeded Faker with New instead.

// FirstName returns a first name using the default Faker.
func FirstName(opts ...fk.Option[Options]) string {
	return New(opts...).FirstName()
}

// LastName returns a last name using the default Faker.
func LastName(opts ...fk.Option[Options]) string {
	return New(opts...).LastName()
}

// Name returns a full name using the default Faker.
func Name(opts ...fk.Option[Options]) string {
	return New(opts...).Name()
}

// NameWithMiddle returns a name with a middle name using the default Faker.
func NameWithMiddle(opts ...fk.Option[Options]) string {
	return New(opts...).NameWithMiddle()
}

// Prefix returns an honorific using the default Faker.
func Prefix(opts ...fk.Option[Options]) string {
	return New(opts...).Prefix()
}

// Suffix returns a name suffix using the default Faker.
func Suffix(opts ...fk.Option[Options]) string {
	return New(opts...).Suffix()
}
