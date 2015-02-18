/*
 * Copyright 2015 Matthew Collins
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package thinkbot

import (
	"fmt"
	"strings"
)

// User contains basic information on a single
// irc user.
//
// Ident/Host may not always be known
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
