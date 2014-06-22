//INI files specific functions

package cfgp

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// parseKeyValue given one line encoded like "key = value" returns corresponding map.
func parseKeyValue(line string) map[string]string {
	kvm := make(map[string]string)
	line = strings.Replace(line, " ", "", -1)
	// Does nothing if no "=" sign.
	if strings.Contains(line, "=") {
		values := strings.Split(line, "=")
		kvm[values[0]] = values[1]
	}
	return kvm
}

// parseINI opens configuration file specified by path and populate Conf.IniData.
// All values as returned as strings, the caller has to make required casting.
// Files must follows INI informal standard:
//
//	https://en.wikipedia.org/wiki/INI_file
//
func (c *Conf) parseINI(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sectionExp, _ := regexp.Compile(`^(\[).+(\])$`)
	commentExp, _ := regexp.Compile(`^(#|;)`)
	// Adds default section "main" in case no one is specified
	section := "main"
	for scanner.Scan() {
		line := scanner.Text()
		if debug {
			fmt.Println("line parsed:", line)
		}
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
			c.IniData[section] = append(c.IniData[section], kv)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if debug {
		fmt.Println("returning map:", c.IniData)
	}
	return nil
}

// HasSection returns true if file has a specific section.
func (c *Conf) HasSection(section string) bool {
	if c.ConfType != "INI" {
		return false
	}
	_, ok := c.IniData[section]
	return ok
}

// HasKey returns true if key is present in section.
func (c *Conf) HasKey(section, key string) bool {
	if c.ConfType != "INI" {
		return false
	}
	if sectionKeys, ok := c.IniData[section]; ok {
		for _, kv := range sectionKeys {
			if _, ok := kv[key]; ok {
				return true
			}
		}
	}
	return false
}

func (c *Conf) GetKey(section, key string) string {
	if c.ConfType != "INI" {
		return ""
	}
	return ""
}

// Returns all key/vaule for specific section
func (c *Conf) GetSection(section string) []map[string]string {
	if c.ConfType != "INI" {
		return nil
	}
	return nil
}

// Returns all sections found in file
func (c *Conf) GetSections(section string) []string {
	if c.ConfType != "INI" {
		return nil
	}
	return nil
}
