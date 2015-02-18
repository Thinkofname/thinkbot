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
	"github.com/thinkofdeath/thinkbot/command"
	"net/http"
)

func initCommands(commands *command.Registry) {
	commands.Register("latest", latest)
	commands.Register("latest spigot", latest)
	commands.Register("latest bukkit", latestBukkit)
}

func latest(b *thinkbot.Bot, user thinkbot.User, target string) {
	hash, err := getCommitHash("spigot")
	if err != nil {
		b.Reply(user, target, fmt.Sprintf("Failed to get the latest version: %s", err))
		return
	}
	bhash, err := getCommitHash("craftbukkit")
	if err != nil {
		b.Reply(user, target, fmt.Sprintf("Failed to get the latest version: %s", err))
		return
	}
	b.Reply(user, target, fmt.Sprintf("The latest is git-Spigot-%s-%s", hash[0:7], bhash[0:7]))
}

func latestBukkit(b *thinkbot.Bot, user thinkbot.User, target string) {
	hash, err := getCommitHash("craftbukkit")
	if err != nil {
		b.Reply(user, target, fmt.Sprintf("Failed to get the latest version: %s", err))
		return
	}
	b.Reply(user, target, fmt.Sprintf("The latest is git-Bukkit-%s", hash[0:7]))
}

type versionReply struct {
	Values []versionCommit `json:"values"`
}

type versionCommit struct {
	ID string `json:"id"`
}

func getCommitHash(repo string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(
		"https://hub.spigotmc.org/stash/rest/api/1.0/projects/SPIGOT/repos/%s/commits?limit=1",
		repo,
	))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var reply versionReply
	err = json.NewDecoder(resp.Body).Decode(&reply)
	return reply.Values[0].ID, err
}
