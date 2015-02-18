package thinkbot

import (
	"encoding/json"
	"fmt"
	"github.com/thinkofdeath/thinkbot/command"
	"net/http"
)

var commands = command.Registry{
	// User and target parameters
	ExtraParameters: 2,
}

func init() {
	commands.Register("latest", latest)
	commands.Register("latest spigot", latest)
	commands.Register("latest bukkit", latestBukkit)
}

func (b *Bot) handleCommand(user User, target, msg string) {
	err := commands.Execute(b, msg, user, target)
	if err != nil {
		b.Reply(user, target, err.Error())
	}
}

func latest(b *Bot, user User, target string) {
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

func latestBukkit(b *Bot, user User, target string) {
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
	Id string `json:"id"`
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
	return reply.Values[0].Id, err
}
