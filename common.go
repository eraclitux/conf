// cfgp - go configuration file parser package
// Copyright (c) 2014 Andrea Masi

// Configuration files parser package fo Go.
//
// Tries to be modular and extendible to support different formats.
//
// Only INI format supported for now. Files must follows INI informal standard:
//
//	https://en.wikipedia.org/wiki/INI_file
//
// YAML support is on the roadmap.
package cfgp

import (
	"errors"
	"regexp"
)

//TODO make this public
type iniDataType map[string][]map[string]string

type Conf struct {
	// iniDataType is map[string][]map[string]string
	IniData iniDataType
	// Store the configuration file format (INI, YAML etc)
	// Actually only INI supported
	ConfType string
}

var debug bool = false

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
		// TODO FIXME
		conf.ConfType = "YAML"
		//conf.ParseYAML(path)
		//return &conf
		return nil, errors.New("YAML not yet implemented")
	}
	return nil, errors.New("Unrecognized format")
}
