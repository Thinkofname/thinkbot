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

func (b *Bot) init() {
	b.Commands.Register("join %", join)
}

var (
	permJoin = Permission{Name: "command.join", Default: false}
)

func join(b *Bot, sender User, target, channel string) {
	if !b.HasPermission(sender, permJoin) {
		panic("you don't have permission for this command")
	}
	if len(channel) < 0 || channel[0] != '#' {
		panic("invalid channel")
	}
	b.JoinChannel(channel)
}
