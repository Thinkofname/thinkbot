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

// Event is anything that can be returned
// by the Event channel
type Event interface{}

// Stop is an event that is sent when the
// bot is stopped (e.g. disconnected)
type Stop struct{}

// Connected is an event when the bot considers
// itself to be fully connected to the server
// and ready for commands
type Connected struct{}

// JoinChannel is an event when the bot joins
// a channel, either by command or forced
type JoinChannel struct {
	Channel string
}

// PartChannel is an event when the bot parts
// a channel, either by command or forced
type PartChannel struct {
	Channel string
}
