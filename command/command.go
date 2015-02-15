package command

import (
	"errors"
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

var quotedStringRegex = regexp.MustCompile("[^\\s`]+|`([^`]*)`")

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

func (r *Registry) Execute(caller interface{}, cmd string, extra ...interface{}) error {
	if r.root == nil {
		return ErrCommandNotFound
	}

	parts := quotedStringRegex.FindAllString(cmd, -1)
	for i, p := range parts {
		if strings.HasPrefix(p, "`") && strings.HasSuffix(p, "`") {
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
		args := make([]reflect.Value, 1 + r.ExtraParameters)
		args[0] = reflect.ValueOf(caller)
		for i, e := range extra {
			args[1 + i] = reflect.ValueOf(e)
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
