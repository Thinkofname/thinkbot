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
	"fmt"
	"reflect"
	"strconv"
)

type TypeHandler interface {
	DefineType(arg string) interface{}
	ParseType(arg string, info interface{}) (interface{}, error)
	Equals(a, b interface{}) bool
}

func (r *Registry) RegisterType(t reflect.Type, handler TypeHandler) {
	if r.typeHandlers == nil {
		r.initTypes()
	}
	_, ok := r.typeHandlers[t]
	if ok {
		panic("type already registered")
	}
	r.typeHandlers[t] = handler
}

func (r *Registry) initTypes() {
	r.typeHandlers = map[reflect.Type]TypeHandler{}

	r.RegisterType(reflect.TypeOf(""), stringHandler{})
}

type stringHandler struct{}

func (stringHandler) DefineType(arg string) interface{} {
	if len(arg) >= 1 {
		i, err := strconv.ParseInt(arg, 10, 32)
		if err != nil {
			panic(err)
		}
		return int(i)
	}
	return -1
}

func (stringHandler) ParseType(arg string, info interface{}) (interface{}, error) {
	limit := info.(int)
	if limit != -1 && len(arg) > limit {
		return nil, fmt.Errorf("string too long (%d > %d)", len(arg), limit)
	}
	return arg, nil
}

func (stringHandler) Equals(a, b interface{}) bool {
	return a.(int) == b.(int)
}
