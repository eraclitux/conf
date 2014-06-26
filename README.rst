
.. contents::

Intro
=====
A go package for configuration files parsing.

Aims to be multiformat, only INI files supported for now.

Usage and examples
==================
See `godocs <http://godoc.org/github.com/eraclitux/cfgp>`_ for in depth documentation.

ini files
---------
Given the example ini file::
        [main]
        one = 42
        three = Zaphod

        [questions]
        answer = 42
        wrong-answer = 43

Example code:

.. code:: go

        pkg main

        import (
                "fmt"
	        "github.com/eraclitux/cfgp"
        )

        func main() {
                // Parse will detect configutation file type (INI, YAML) via extention.
                // Only INI supported for now.
                conf, err := cfgp.Parse("test_data/example.ini")
                if err != nil {
                        panic("Unable to parse configuration file")
                }
                // Check if a specific section exists
                section := "main"
                if conf.HasSection(section) {
                        fmt.Printf("Section %s exists\n", section)
                }
                // Check if a specific key exists
                key, section := "wrong-answer", "questions"
                if conf.HasKey(section, key) {
                        fmt.Printf("Key %s exists\n", key)
                }
	        // Retrieve a specific key
                key, section = "answer", "questions"
                if value, err := conf.GetKey(section, key); err == nil {
                        fmt.Printf("Key %s is %s\n", key, value)
                }
        }

Notes
=====
Work in progress
