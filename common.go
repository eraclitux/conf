// cfgp - go configuration file parser package
// Copyright (c) 2014 Andrea Masi

// Configuration files parser package fo Go.
// Try to be modular and extendible to support different formats.
// Only INI format supported for now, YAML on the roadmap.
package cfgp

import (
	"regexp"
)

type iniDataType map[string][]map[string]string

type Conf struct {
	// iniDataType is map[string][]map[string]string
	IniData iniDataType
	// Store the configuration file format (INI, YAML etc)
	// Actually only INI supported
	ConfType string
}

var debug bool = false

// Parse guesses configuration type by file extention and call specific parser.
// .ini|.txt|.cfg are evaluated as INI files.
func Parse(path string) *Conf {
	conf := Conf{}
	if match, _ := regexp.MatchString(`\.(ini|txt|cfg)$`, path); match {
		conf.IniData = make(iniDataType)
		conf.ConfType = "INI"
		conf.ParseINI(path)
		return &conf
	} else if match, _ := regexp.MatchString(`\.(yaml)$`, path); match {
		// TODO FIXME
		conf.ConfType = "YAML"
		//conf.ParseYAML(path)
		//return &conf
		return nil
	}
	return nil
}
