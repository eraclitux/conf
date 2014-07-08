====
cfgp
====

|image0|_

.. |image0| image:: https://godoc.org/github.com/eraclitux/cfgp?status.png
.. _image0: https://godoc.org/github.com/eraclitux/cfgp

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
                if conf.IniHasSection(section) {
                        fmt.Printf("Section %s exists\n", section)
                }
                // Check if a specific key exists
                key, section := "wrong-answer", "questions"
                if conf.IniHasKey(section, key) {
                        fmt.Printf("Key %s exists\n", key)
                }
                // Retrieve a specific key in a section
                key, section = "answer", "questions"
                if value, err := conf.IniGetKey(section, key); err == nil {
                        fmt.Printf("Key %s is %s\n", key, value)
                }
                // Retrieve all keys in a section
                section = "questions"
                if section, err := conf.IniGetSection(section); err == nil {
                        for _, kv := range section {
                                for k, v := range kv {
                                        fmt.Printf("Key:%s,value:%s;", k, v)
                                }
                        }
                        fmt.Println("")
                }
        }

Notes
=====
YAML not yet supported.
