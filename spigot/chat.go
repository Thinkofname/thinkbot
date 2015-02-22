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

package spigot

import (
	"encoding/json"
	"fmt"
	"github.com/thinkofdeath/thinkbot"
	"net/http"
	"net/url"
	"regexp"
)

var (
	spigotVersionMatcher = regexp.MustCompile(`(?i)git-Spigot-([0-9a-f]{7})-([0-9a-f]{7})`)
	bukkitVersionMatcher = regexp.MustCompile(`(?i)git-Bukkit-([0-9a-f]{7})`)
)

func initChat(b *thinkbot.BotConfig) {
	b.AddChatHandler(spigotVersionMatcher, func(b *thinkbot.Bot, sender thinkbot.User, target, message string) error {
		return checkSpigotVersion(b, sender, target, spigotVersionMatcher.FindStringSubmatch(message))
	})
	b.AddChatHandler(bukkitVersionMatcher, func(b *thinkbot.Bot, sender thinkbot.User, target, message string) error {
		return checkSpigotVersion(b, sender, target, bukkitVersionMatcher.FindStringSubmatch(message))
	})
}

func checkSpigotVersion(b *thinkbot.Bot, sender thinkbot.User, target string, info []string) error {
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
