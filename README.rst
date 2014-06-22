
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
Given the example ini file (es my-conf.ini)::

        [main]
        max-hype = 10

        [cache]
        timeout = 60s
        cache-size=100M

You can parse it from your code  

.. code:: go

        pkg main

        import (
                "cfgp"
                "fmt"
        )

        func main() {
                conf, err := cfgp.Parse("my-conf.ini")
                if err != nil {
                        panic("Unable to parse configuration file")
                }
                //Retrive a specific key
                if conf.HasKey("main", "cache-size") {
                        key := conf.GetKey("main", "cache-size")
                        fmt.Print(key)
                }
                //Print all keys from a specific section
                section := conf.GetSection("cache")
                for k, v := range section {
                        fmt.Print("key:", k, "value:", v)
                }
        }



Notes
=====
Work in progress