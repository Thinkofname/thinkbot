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
