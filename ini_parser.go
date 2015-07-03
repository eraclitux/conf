//INI files specific functions

package cfgp

import (
	"bufio"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/eraclitux/stracer"
)

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

func putInStruct(structValue reflect.Value, kv []string) error {
	// FIXME handle different types.
	stracer.Traceln("handling pair:", kv)
	f := strings.Title(kv[0])
	fieldValue := structValue.FieldByName(f)
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
		case reflect.Bool:
			b, err := strconv.ParseBool(kv[1])
			if err != nil {
				return err
			}
			fieldValue.SetBool(b)
		default:
			return ErrUnknownFlagType
		}
	}
	return nil
}

// parseINI opens configuration file specified by path and populate
// passed struct.
// Files must follows INI informal standard:
//
//	https://en.wikipedia.org/wiki/INI_file
//
// FIXME Current implementation stores info about section but
// discards it. Use reflection tags to specify which section to use.
func parseINI(path string, structValue reflect.Value) error {
	conf := make(map[string][][]string)
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
