// cfgp - go configuration file parser package
// Copyright (c) 2014 Andrea Masi

// Package cfgp is a configuration parser fo Go.
//
// It tries to be modular and easily extendible to support different formats.
//
// For file, only INI format supported for now. Files must follows INI informal standard:
//
//	https://en.wikipedia.org/wiki/INI_file
//
// This is a work in progress, better packages are out there.
package cfgp

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/eraclitux/stracer"
)

var ErrNeedPointer = errors.New("cfgp: pointer to struct expected")
var ErrFileFormat = errors.New("cfgp: unrecognized file format, only (ini|txt|cfg) supported")
var ErrUnknownFlagType = errors.New("cfgp: unknown kind flag")

func getStructValue(confPtr interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(confPtr)
	if v.Kind() == reflect.Ptr {
		return v.Elem(), nil
	}
	return reflect.Value{}, ErrNeedPointer
}

// myFlag implements Flag.Value.
// TODO is filed needed?
type myFlag struct {
	field      reflect.StructField
	fieldValue reflect.Value
	isBool     bool
}

func (s *myFlag) String() string {
	return s.field.Name
}

// IsBoolFlag istructs the command-line parser
// to makes -name equivalent to -name=true rather than
// using the next command-line argument.
func (s *myFlag) IsBoolFlag() bool {
	return s.isBool
}

func (s *myFlag) Set(arg string) error {
	stracer.Traceln("setting flag", s.field.Name)
	switch s.fieldValue.Kind() {
	case reflect.Int:
		n, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		s.fieldValue.SetInt(int64(n))
	case reflect.String:
		s.fieldValue.SetString(arg)
	case reflect.Bool:
		b, err := strconv.ParseBool(arg)
		if err != nil {
			return err
		}
		s.fieldValue.SetBool(b)
	default:
		return ErrUnknownFlagType
	}
	return nil
}

func makeHelpMessage(f reflect.StructField) string {
	// TODO use reflect' tags to get help message
	stracer.Traceln("tag:", f.Tag.Get("cfgp"))
	switch f.Type.Kind() {
	case reflect.Int:
		return "set an int value"
	case reflect.String:
		return "set a string value"
	case reflect.Bool:
		return "set a bool value"
	}
	return "give a value to flag"
}

func isBool(v reflect.Value) bool {
	if v.Kind() == reflect.Bool {
		return true
	}
	return false
}

func createFlag(f reflect.StructField, fieldValue reflect.Value, fs *flag.FlagSet) {
	name := strings.ToLower(f.Name)
	stracer.Traceln("Creating flag:", name)
	fs.Var(&myFlag{f, fieldValue, isBool(fieldValue)}, name, makeHelpMessage(f))
}

func parseFlags(s reflect.Value) error {
	flagSet := flag.NewFlagSet("cfgp", flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flagSet.PrintDefaults()
	}
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		fieldValue := s.Field(i)
		if fieldValue.CanSet() {
			createFlag(typeOfT.Field(i), fieldValue, flagSet)
		}
	}
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		stracer.Traceln("This is not executed.")
		return err
	}
	return nil
}

// Parse guesses configuration type by file extention and call specific parser.
// (.ini|.txt|.cfg) are evaluated as INI files.
func Parse(path string, confPtr interface{}) error {
	structValue, err := getStructValue(confPtr)
	if err != nil {
		return err
	}
	if path != "" {
		if match, _ := regexp.MatchString(`\.(ini|txt|cfg)$`, path); match {
			err := parseINI(path, structValue)
			if err != nil {
				return err
			}
		} else if match, _ := regexp.MatchString(`\.(yaml)$`, path); match {
			return errors.New("YAML not yet implemented. Want you help?")
		} else {
			return ErrFileFormat
		}
	}
	parseFlags(structValue)
	return nil
}
