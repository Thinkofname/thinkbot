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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

var (
	spigotVersionMatcher = regexp.MustCompile(`(?i)git-Spigot-([0-9a-f]{7})-([0-9a-f]{7})`)
	bukkitVersionMatcher = regexp.MustCompile(`(?i)git-Bukkit-([0-9a-f]{7})`)
)

func (b *Bot) handleMessage(sender User, target, message string) {
	matches := spigotVersionMatcher.FindStringSubmatch(message)
	if matches == nil {
		matches = bukkitVersionMatcher.FindStringSubmatch(message)
	}
	if matches != nil {
		go func() {
			err := b.checkSpigotVersion(sender, target, matches)
			if err != nil {
				b.Reply(sender, target, fmt.Sprintf("Sorry I had an issue checking your version, %s", err))
			}
		}()
		return
	}
}

func (b *Bot) checkSpigotVersion(sender User, target string, info []string) error {
	distance := 0
	extra := ""
	if len(info) == 3 {
		spigot, err := getDistanceFromLatest("spigot", info[1])
		if err != nil {
			return err
		}
		craftbukkit, err := getDistanceFromLatest("craftbukkit", info[2])
		if err != nil {
			return err
		}
		distance = craftbukkit + spigot
		extra = fmt.Sprintf("(%d/%d)", spigot, craftbukkit)
	} else {
		craftbukkit, err := getDistanceFromLatest("craftbukkit", info[1])
		if err != nil {
			return err
		}
		distance = craftbukkit
		extra = fmt.Sprintf("(%d)", craftbukkit)
	}

	if distance == 0 {
		b.Reply(sender, target, "You have the latest version")
		return nil
	}

	s := ""
	if distance != 1 {
		s = "s"
	}
	b.Reply(sender, target, fmt.Sprintf(
		"You are behind by %d version%s, please rerun BuildTools %s",
		distance,
		s,
		extra,
	))
	return nil
}

// Reply sends a message to the user in the same way
// the message was sent.
//
// If the message was sent in a channel the message
// will be sent back to that channel with the sender's
// nickname prefixed.
//
// If the message was sent in a private message then
// this will just reply normally
func (b *Bot) Reply(sender User, target, message string) {
	if target[0] == '#' {
		b.SendMessage(target, fmt.Sprintf("%s: %s", sender.Nickname, message))
	} else {
		b.SendMessage(sender.Nickname, message)
	}
}

func getDistanceFromLatest(repo, hash string) (int, error) {
	resp, err := http.Get(fmt.Sprintf(
		"https://hub.spigotmc.org/stash/rest/api/1.0/projects/SPIGOT/repos/%s/commits?since=%s&withCounts=true",
		repo,
		url.QueryEscape(hash),
	))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	var reply commitsReply
	err = json.NewDecoder(resp.Body).Decode(&reply)
	return reply.TotalCount, err
}

type commitsReply struct {
	TotalCount int `json:"totalCount"`
}
