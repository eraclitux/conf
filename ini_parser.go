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
	// Adds default section "default" in case no one is specified
	section := "default"
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

// IsIni returns true whenever Conf has been parsed as INI file.
func (c *Conf) IsIni() bool {
	//TODO add tests
	if c.ConfType == "INI" {
		return true
	}
	return false
}

// IniHasSection returns true if file has a specific section.
func (c *Conf) IniHasSection(section string) bool {
	if !c.IsIni() {
		return false
	}
	_, ok := c.IniData[section]
	return ok
}

// IniHasKey returns true if key is present in section.
func (c *Conf) IniHasKey(section, key string) bool {
	if !c.IsIni() {
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

// IniGetKey returns value given section/key
func (c *Conf) IniGetKey(section, key string) (string, error) {
	if !c.IsIni() {
		return "", fmt.Errorf("Not an INI file")
	}
	if sectionKeys, ok := c.IniData[section]; ok {
		for _, kv := range sectionKeys {
			if value, ok := kv[key]; ok {
				return value, nil
			}
		}
		return "", fmt.Errorf("Key %s not found", key)
	}
	return "", fmt.Errorf("Section %s not found", section)
}

// IniGetSection returns all key/vaule for specific section.
func (c *Conf) IniGetSection(section string) ([]map[string]string, error) {
	if !c.IsIni() {
		return nil, fmt.Errorf("Not an INI file")
	}
	if sectionKeys, ok := c.IniData[section]; ok {
		return sectionKeys, nil
	}
	return nil, fmt.Errorf("Section %s not found", section)
}

// IniGetSections returns all sections's names found in file as a slice of string.
func (c *Conf) IniGetSections() ([]string, error) {
	if !c.IsIni() {
		return nil, fmt.Errorf("Not an INI file")
	}
	sections := make([]string, len(c.IniData))
	i := 0
	for k, _ := range c.IniData {
		sections[i] = k
		i++
	}
	return sections, nil
}
