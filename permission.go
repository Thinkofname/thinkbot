package thinkbot

// Permission is a key map of permissions for users
type Permission struct {
	Name    string
	Default bool
}

// HasPermission returns the user's set value for the
// permission or its default value if not set
func (b *Bot) HasPermission(user User, perm Permission) bool {
	if b.permissionContainer != nil {
		return b.permissionContainer.HasPermission(b, user, perm)
	}
	return perm.Default
}

// SetPermissionContainer sets the container used for checking
// permissions for irc users. Should only be set during the
// bot's init
func (b *Bot) SetPermissionContainer(pc PermissionContainer) {
	b.permissionContainer = pc
}

// PermissionContainer is used for looking up permissions
// for irc users
type PermissionContainer interface {
	HasPermission(b *Bot, user User, perm Permission) bool
}
