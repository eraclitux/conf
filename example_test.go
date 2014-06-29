package cfgp_test

import (
	"fmt"
	"github.com/eraclitux/cfgp"
)

func Example() {
	// Parse will detect configutation file type (INI, YAML) via extention.
	// Only INI supported for now.
	conf, err := cfgp.Parse("test_data/example.ini")
	if err != nil {
		panic("Unable to parse configuration file")
	}
	// Check if a specific section exists
	section := "main"
	if conf.IniHasSection(section) {
		fmt.Printf("Section %q exists\n", section)
	}
	// Check if a specific key exists
	key, section := "wrong-answer", "questions"
	if conf.IniHasKey(section, key) {
		fmt.Printf("Key %q exists\n", key)
	}
	// Retrieve a specific key in a section
	key, section = "answer", "questions"
	if value, err := conf.IniGetKey(section, key); err == nil {
		fmt.Printf("Key %q is %q\n", key, value)
	}
	// Retrieve all keys in a section
	section = "questions"
	if section, err := conf.IniGetSection(section); err == nil {
		for _, kv := range section {
			for k, v := range kv {
				fmt.Printf("Key:%q,value:%q; ", k, v)
			}
		}
		fmt.Println("")
	}
	// Output:
	// Section "main" exists
	// Key "wrong-answer" exists
	// Key "answer" is "42"
	// Key:"answer",value:"42"; Key:"wrong-answer",value:"43";
}
