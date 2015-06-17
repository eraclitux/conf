//INI files specific functions

package cfgp

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/eraclitux/stracer"
)

var ErrNeedPointer = errors.New("pointer to struct expected")

// parseKeyValue given one line encoded like "key = value" returns corresponding
// []string with "key" > kv[0] and "value" > kv[1].
func parseKeyValue(line string) []string {
	// Check for inline comments.
	if strings.Contains(line, ";") {
		line = strings.Split(line, ";")[0]
	} else if strings.Contains(line, "#") {
		line = strings.Split(line, "#")[0]
	}
	line = strings.Replace(line, " ", "", -1)
	// Does nothing if no "=" sign.
	if strings.Contains(line, "=") {
		return strings.Split(line, "=")

	}
	return nil
}

func getStructValue(confPtr interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(confPtr)
	if v.Kind() == reflect.Ptr {
		return v.Elem(), nil
	}
	return reflect.Value{}, ErrNeedPointer
}

func putInStruct(structValue reflect.Value, kv []string) error {
	// FIXME handle different types.
	stracer.Traceln("handling pair:", kv)
	f := strings.Title(kv[0])
	fieldValue := structValue.FieldByName(f)
	stracer.Traceln("k to title:", f, "kind in struct:", fieldValue.Kind(), "is settable:", fieldValue.CanSet())
	if fieldValue.CanSet() {
		switch fieldValue.Kind() {
		case reflect.Int:
			i, err := strconv.Atoi(kv[1])
			if err != nil {
				return err
			}
			fieldValue.SetInt(int64(i))
		case reflect.String:
			fieldValue.SetString(kv[1])
		}
	}
	return nil
}

// parseINI opens configuration file specified by path and populate Conf.IniData.
// Files must follows INI informal standard:
//
//	https://en.wikipedia.org/wiki/INI_file
//
func parseINI(path string, confPtr interface{}) error {
	conf := make(map[string][][]string)
	structValue, err := getStructValue(confPtr)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sectionExp := regexp.MustCompile(`^(\[).+(\])$`)
	commentExp := regexp.MustCompile(`^(#|;)`)
	// Adds default section "default" in case no one is specified
	section := "default"
	for scanner.Scan() {
		line := scanner.Text()
		stracer.Traceln("raw line to parse:", line)
		if commentExp.MatchString(line) {
			continue
		} else if sectionExp.MatchString(line) {
			// Removes spaces too so "[ section]" is parsed correctly
			section = strings.Trim(line, "[] ")
			continue
		}
		kv := parseKeyValue(line)
		// This even prevents empty line to be added
		if len(kv) > 0 {
			putInStruct(structValue, kv)
			conf[section] = append(conf[section], kv)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	stracer.Traceln("coded map:", conf)
	return nil
}
