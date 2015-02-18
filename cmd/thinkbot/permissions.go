package main

import (
	"github.com/thinkofdeath/thinkbot"
)

type configPermissions struct{}

var _ thinkbot.PermissionContainer = configPermissions{}

func (configPermissions) HasPermission(
	b *thinkbot.Bot,
	user thinkbot.User,
	perm thinkbot.Permission) bool {
	configLock.RLock()
	defer configLock.RUnlock()
	u, ok := config.Users[user.Host]
	if !ok {
		return perm.Default
	}
	p, ok := u.Permissions[perm.Name]
	if !ok {
		return perm.Default
	}
	return p
}
