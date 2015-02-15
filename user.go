package thinkbot

import (
	"fmt"
	"strings"
)

type User struct {
	Nickname string
	Ident    string
	Host     string
}

func parseUser(u string) User {
	if !strings.ContainsRune(u, '!') || !strings.ContainsRune(u, '@') {
		return User{
			Nickname: u,
		}
	}
	nick := u[:strings.IndexRune(u, '!')]
	u = u[len(nick)+1:]
	ident := u[:strings.IndexRune(u, '@')]
	return User{
		Nickname: nick,
		Ident:    ident,
		Host:     u[len(ident)+1:],
	}
}

func (u User) String() string {
	return fmt.Sprintf("%s!%s@%s", u.Nickname, u.Ident, u.Host)
}
