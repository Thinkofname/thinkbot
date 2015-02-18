// Package command implements a system for parsing and executing
// commands
//
// When registering commands a description of the command is
// required. For basic commands the format of this is simple:
//
//     commandname sub1 sub2 sub...
//
// BUG(Think): Complex commands are NYI
//
// Complex commands can be created by using % to specify
// arguments. The type of the argument will be inferred
// from the type over the passed function pointer. Extra
// constraints can be added after the % to gain finer control
// over the argument.
//
// Executing commands works by treating whitespace at delimiters
// between arguments with the exception of whitespace contained
// within quotes (") as that will be treated as a single argument
package command
