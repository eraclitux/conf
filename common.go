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

// FIXME make this public?
type iniDataType map[string][]map[string]string

// Conf stores parsed data from configuration file
type Conf struct {
	// iniDataType is map[string][]map[string]string
	IniData iniDataType
	// Store the configuration file format (INI, YAML etc)
	// Actually only INI supported
	ConfType string
}

// Parse guesses configuration type by file extention and call specific parser to pupulate Conf.
//
// (.ini|.txt|.cfg) are evaluated as INI files.
func Parse(path string) (*Conf, error) {
	conf := Conf{}
	if match, _ := regexp.MatchString(`\.(ini|txt|cfg)$`, path); match {
		conf.IniData = make(iniDataType)
		conf.ConfType = "INI"
		err := conf.parseINI(path)
		if err != nil {
			return nil, err
		}
		return &conf, nil
	} else if match, _ := regexp.MatchString(`\.(yaml)$`, path); match {
		// TODO
		conf.ConfType = "YAML"
		//conf.ParseYAML(path)
		//return &conf
		return nil, errors.New("YAML not yet implemented")
	}
	return nil, errors.New("Unrecognized format")
}
