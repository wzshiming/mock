package mock

import (
	"github.com/wzshiming/crun"
)

// RandDomain Returns a random domain
func RandDomain() string {
	return domain.Rand()
}

// RandURL Returns a random URL
func RandURL() string {
	return url.Rand()
}

// RandUUID Returns a random UUID
func RandUUID() string {
	return uuid.Rand()
}

// RandEmail Returns a random Email
func RandEmail() string {
	return email.Rand()
}

// RandName Returns a random name
func RandName() string {
	return name.Rand()
}

// RandText Returns a random text
func RandText() string {
	return text.Rand()
}

// RandWord Returns a random word
func RandWord() string {
	return word.Rand()
}

// crun constant
var (
	protocol = crun.MustCompile(_protocol)
	domain   = crun.MustCompile(_domain)
	url      = crun.MustCompile(_url)
	email    = crun.MustCompile(_email)
	uuid     = crun.MustCompile(_uuid)
	name     = crun.MustCompile(_name)
	text     = crun.MustCompile(_text)
	word     = crun.MustCompile(_word)
)
