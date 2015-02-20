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
