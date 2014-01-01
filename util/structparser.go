package util

//
// Copyright Jen Andre (jandre@gmail.com)
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR O

import (
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
)

//
// Given an array of values, e.g. a[0..n]
// Assign each value to a struct at position 0..n
//
// For each array value a[0..n], try to convert it to
//
//
func StructParser(v *reflect.Value, lines []string) error {

	typeOf := v.Type()

	for i := 0; i < v.NumField(); i++ {

		if i > len(lines) {
			break
		}

		field := v.Field(i)
		line := strings.TrimSpace(lines[i])

		switch field.Kind() {
		// TODO: more types
		case reflect.Int:
			if val, err := strconv.Atoi(line); err != nil {
				log.Println("Parsing Error:", typeOf.Field(i).Name, err)
				return err
			} else {
				field.SetInt(int64(val))
			}
		case reflect.Int64:
			if val, err := strconv.ParseInt(line, 10, 64); err != nil {
				log.Println("Parsing Error:", typeOf.Field(i).Name, err)
				return err
			} else {
				field.SetInt(val)
			}

		case reflect.String:
			field.SetString(line)

		case reflect.Struct:
			//
			// in the case of a jiffy timestamp value, we abstract it to a type
			// that auto converts jiffies to epoch timestamps
			//
			if field.Type().Name() == "EpochTimestamp" {
				if jiffies, err := strconv.ParseInt(line, 10, 64); err != nil {
					log.Println("Parsing Error:", typeOf.Field(i).Name, err)
					return err
				} else {
					e := NewEpochTimestamp(jiffies)
					field.Set(reflect.ValueOf(e))
				}
			} else {
				return errors.New("Unknown type: " + field.Type().Name())
			}
		default:
			return errors.New("Unknown type: " + typeOf.Field(i).Name + "; type=" + field.Type().Name())
		}

	}
	return nil

}
