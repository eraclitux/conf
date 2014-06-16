package cfgp

import (
	"regexp"
)
type iniDataType map[string][]map[string]string

type Conf struct {
	IniData iniDataType
	//Store the configuration file format (INI, YAML etc)
	//Actually only INI supported 
	ConfType string
}

//Guess configuration type by extention and call specific parser
func Parse(path string) *Conf {
	conf := Conf{}
	if match, _ := regexp.MatchString(`\.(ini|txt|cfg)$`, path); match {
		conf.IniData = make(iniDataType)
		conf.ConfType = "INI"
		conf.ParseINI(path)
		return &conf
	} else if match, _ := regexp.MatchString(`\.(yaml)$`, path); match {
		//TODO FIXME
		conf.ConfType = "YAML"
		//conf.ParseYAML(path)
		//return &conf
		return nil
	}
	return nil
}
