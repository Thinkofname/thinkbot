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

type Registry struct {
	ExtraParameters int

	root *commandNode
}

var (
	ErrCommandNotFound = errors.New("Command not found")
)

var quotedStringRegex = regexp.MustCompile(`[^\s"]+|"([^"]*)"`)

func (r *Registry) Register(desc string, f interface{}) {
	args := strings.Split(desc, " ")

	if len(args) < 1 {
		panic("Invalid command desc")
	}

	if r.root == nil {
		r.root = &commandNode{
			childNodes: map[string]*commandNode{},
		}
	}

	current := r.root
	for _, arg := range args {
		if strings.HasPrefix(arg, "%") {
			panic("Unsupported op")
			continue
		}
		current = current.subNode(arg)
	}

	if current.f != nil {
		panic("Double registered command")
	}

	current.f = f
}

func (r *Registry) Execute(caller interface{}, cmd string, extra ...interface{}) (err error) {
	// No commands registered
	if r.root == nil {
		return ErrCommandNotFound
	}
	// Catch and return any errors throw to prevent crashing
	defer func() {
		if e := recover(); e != nil {
			var ok bool
			err, ok = e.(error)
			if !ok {
				err = fmt.Errorf("%v", e)
			}
		}
	}()


	parts := quotedStringRegex.FindAllString(cmd, -1)
	for i, p := range parts {
		if strings.HasPrefix(p, `"`) && strings.HasSuffix(p, `"`) {
			parts[i] = p[1 : len(p)-1]
		}
	}

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

	if current != nil && current.f != nil {
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
