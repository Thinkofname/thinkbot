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

package main

import (
	"bytes"
	"github.com/thinkofdeath/thinkbot"
	"github.com/thinkofdeath/thinkbot/command"
	"regexp"
)

var (
	permJoin      = thinkbot.Permission{Name: "command.join", Default: false}
	permPart      = thinkbot.Permission{Name: "command.part", Default: false}
	permPrefix    = thinkbot.Permission{Name: "command.prefix", Default: false}
	permAutoReply = thinkbot.Permission{Name: "command.autoreply", Default: false}
)

func initCommands(cmd *command.Registry) {
	cmd.Register("join %", join)
	cmd.Register("part %", part)
	cmd.Register("part", partCurrent)
	cmd.Register("prefix add %", addPrefix)
	cmd.Register("prefix remove %", removePrefix)
	cmd.Register("autoreply add % % %", addAutoReply)
	cmd.Register("autoreply remove %", removeAutoReply)
	cmd.Register("autoreply", listAutoReply)
}

func autoReplyFunc(msg string) func(b *thinkbot.Bot, sender thinkbot.User, target, message string) error {
	return func(b *thinkbot.Bot, sender thinkbot.User, target, message string) error {
		b.Reply(sender, target, msg)
		return nil
	}
}

func listAutoReply(b *thinkbot.Bot, sender thinkbot.User, target string) {
	if !b.HasPermission(sender, permAutoReply) {
		panic("you don't have permission for this command")
	}

	configLock.RLock()
	defer configLock.RUnlock()
	var buf bytes.Buffer
	for n, _ := range config.AutoReplies {
		buf.WriteString(n)
		buf.WriteRune(' ')
	}
	b.Reply(sender, target, buf.String())
}

func removeAutoReply(b *thinkbot.Bot, sender thinkbot.User, target, name string) {
	if !b.HasPermission(sender, permAutoReply) {
		panic("you don't have permission for this command")
	}

	// Update the config
	configLock.Lock()
	defer configLock.Unlock()
	_, ok := config.AutoReplies[name]
	if !ok {
		b.Reply(sender, target, "No auto-reply with that name found")
	} else {
		delete(config.AutoReplies, name)
		saveConfig(config)
		b.Reply(sender, target, "Removed")
	}
}

func addAutoReply(b *thinkbot.Bot, sender thinkbot.User, target, name, re, msg string) {
	if !b.HasPermission(sender, permAutoReply) {
		panic("you don't have permission for this command")
	}

	r := regexp.MustCompile(re)

	// Update the config
	configLock.Lock()
	defer configLock.Unlock()
	_, ok := config.AutoReplies[name]
	if !ok {
		b.AddChatHandler(r, autoReplyFunc(msg))
		config.AutoReplies[name] = autoReply{re, msg}
		saveConfig(config)
		b.Reply(sender, target, "Added auto-reply")
	} else {
		b.Reply(sender, target, "A auto-reply with that name already exists")
	}
}

func addPrefix(b *thinkbot.Bot, sender thinkbot.User, target, prefix string) {
	if !b.HasPermission(sender, permPrefix) {
		panic("you don't have permission for this command")
	}
	b.AddCommandPrefix(prefix)

	// Update the config
	configLock.Lock()
	defer configLock.Unlock()
	for _, p := range config.CommandPrefix {
		if p == prefix {
			b.Reply(sender, target, "I already know that prefix")
			return
		}
	}
	config.CommandPrefix = append(config.CommandPrefix, prefix)
	saveConfig(config)
	b.Reply(sender, target, "Prefix added")
}

func removePrefix(b *thinkbot.Bot, sender thinkbot.User, target, prefix string) {
	if !b.HasPermission(sender, permPrefix) {
		panic("you don't have permission for this command")
	}
	b.RemoveCommandPrefix(prefix)

	// Update the config
	configLock.Lock()
	defer configLock.Unlock()
	for i, p := range config.CommandPrefix {
		if p == prefix {
			config.CommandPrefix = append(config.CommandPrefix[:i], config.CommandPrefix[i+1:]...)
			saveConfig(config)
			b.Reply(sender, target, "Prefix removed")
			return
		}
	}
	b.Reply(sender, target, "Prefix not found")
}

func join(b *thinkbot.Bot, sender thinkbot.User, target, channel string) {
	if !b.HasPermission(sender, permJoin) {
		panic("you don't have permission for this command")
	}
	if len(channel) < 0 || channel[0] != '#' {
		panic("invalid channel")
	}
	b.JoinChannel(channel)
}

func part(b *thinkbot.Bot, sender thinkbot.User, target, channel string) {
	if !b.HasPermission(sender, permPart) {
		panic("you don't have permission for this command")
	}
	if len(channel) < 0 || channel[0] != '#' {
		panic("invalid channel")
	}
	b.PartChannel(channel)
}

func partCurrent(b *thinkbot.Bot, sender thinkbot.User, target string) {
	if !b.HasPermission(sender, permPart) {
		panic("you don't have permission for this command")
	}
	if len(target) < 0 || target[0] != '#' {
		panic("not in a channel")
	}
	b.PartChannel(target)
}
