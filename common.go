// cfgp - go configuration file parser package
// Copyright (c) 2014 Andrea Masi

// Package cfgp is a configuration file parser fo Go.
//
// It tries to be modular and easily extendible to support different formats.
//
// Only INI format supported for now. Files must follows INI informal standard:
//
//	https://en.wikipedia.org/wiki/INI_file
//
// This is a work in progress, better packages are out there.
package cfgp

import (
	"errors"
	"regexp"
)

// Parse guesses configuration type by file extention and call specific parser to pupulate Conf.
// (.ini|.txt|.cfg) are evaluated as INI files.
func Parse(path string, confPtr interface{}) error {
	if match, _ := regexp.MatchString(`\.(ini|txt|cfg)$`, path); match {
		err := parseINI(path, confPtr)
		if err != nil {
			return err
		}
		return nil
	} else if match, _ := regexp.MatchString(`\.(yaml)$`, path); match {
		return errors.New("YAML not yet implemented")
	}
	return errors.New("unrecognized file format")
}
