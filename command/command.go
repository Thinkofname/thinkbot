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

package command

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Registry contains information required to store
// and execute commands
//
// ExtraParameters can be set to specify how many
// extra parameters a command function will have
// after the caller argument
type Registry struct {
	ExtraParameters int

	root *commandNode
}

var (
	// ErrCommandNotFound is returned when no matching command was found
	ErrCommandNotFound = errors.New("command not found")
)

var quotedStringRegex = regexp.MustCompile(`[^\s"]+|"([^"]*)"`)

// Register adds the passed function to the command registry
// using the description to decided on its name and location.
//
// f must be a function
//
// This is designed to panic instead of returning an error
// because its intended to be used in init methods and fail
// early on mistakes
func (r *Registry) Register(desc string, f interface{}) {
	if reflect.TypeOf(f).Kind() != reflect.Func {
		panic("f must be a function")
	}
	args := strings.Split(desc, " ")

	if len(args[0]) < 1 {
		panic("Invalid command desc")
	}

	if r.root == nil {
		r.root = &commandNode{
			childNodes: map[string]*commandNode{},
		}
	}

	current := r.root
	for _, arg := range args {
		if arg[0] == '%' {
			panic("Unsupported op")
		}
		current = current.subNode(arg)
	}

	if current.f != nil {
		panic("Double registered command")
	}

	current.f = f
}

// Execute tries to execute the specified command as the passed
// caller.
//
// Panics if the number of extra arguments doesn't match the
// amount specified in Registry's ExtraParameters
func (r *Registry) Execute(caller interface{}, cmd string, extra ...interface{}) (err error) {
	if len(extra) != r.ExtraParameters {
		panic("Incorrect number of extra parameters")
	}

	if r.root == nil {
		return ErrCommandNotFound
	}
	// Catch and return any errors thrown to prevent crashing
	defer func() {
		if e := recover(); e != nil {
			var ok bool
			err, ok = e.(error)
			if !ok {
				err = fmt.Errorf("%v", e)
			}
		}
	}()

	// Unlike most command systems this supports quoting of
	// arguments using ". The regex used leaves the quotes
	// in the resulting string so we go through a strip them
	parts := quotedStringRegex.FindAllString(cmd, -1)
	for i, p := range parts {
		if strings.HasPrefix(p, `"`) && strings.HasSuffix(p, `"`) {
			parts[i] = p[1 : len(p)-1]
		}
	}

	// Currently we just start at the root node and work our
	// way through until we hit our command or a dead end.
	// When complex commands are supported this will have to
	// change to support rewinding
	current := r.root
	pos := 0
	for pos < len(parts) {
		part := parts[pos]
		if cn, ok := current.childNodes[strings.ToLower(part)]; ok {
			current = cn
			pos++
		} else {
			return ErrCommandNotFound
		}
	}

	// Its possible to reach a node which doesn't have
	// a command assigned (e.g. part of a sub-command)
	// so we have to check that too
	if current != nil && current.f != nil {
		// No checks are preformed on the function here
		// as they should have been checked in Register
		f := reflect.ValueOf(current.f)
		args := make([]reflect.Value, 1+r.ExtraParameters)
		args[0] = reflect.ValueOf(caller)
		for i, e := range extra {
			args[1+i] = reflect.ValueOf(e)
		}
		f.Call(args)
		return nil
	}

	return ErrCommandNotFound
}

type commandNode struct {
	childNodes map[string]*commandNode
	f          interface{}
}

func (cn *commandNode) subNode(name string) *commandNode {
	node, ok := cn.childNodes[name]
	if !ok {
		node = &commandNode{
			childNodes: map[string]*commandNode{},
		}
		cn.childNodes[name] = node
	}
	return node
}
