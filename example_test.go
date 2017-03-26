// Copyright (c) 2015 Andrea Masi. All rights reserved.
// Use of this source code is governed by MIT license
// which that can be found in the LICENSE.txt file.

package conf_test

import (
	"fmt"
	"log"

	"github.com/eraclitux/conf"
)

type myConf struct {
	Address string
	Port    string
	// A command line flag "-users", which expects an int value,
	// will be created.
	// Same key name will be searched in configuration file.
	NumberOfUsers int `conf:"users,number of users,"`
	Daemon        bool
	Message       string
}

func Example() {
	// To create a dafault value for a flag
	// assign it when instantiate the conf struct.
	c := myConf{Message: "A default value"}
	conf.Path = "test_data/one.ini"
	err := conf.Parse(&c)
	if err != nil {
		log.Fatal("Unable to parse configuration", err)
	}
	fmt.Println("address:", c.Address)
	fmt.Println("port:", c.Port)
	fmt.Println("number of users:", c.NumberOfUsers)

	//Output:
	//address: localhost
	//port: 8080
	//number of users: 42

}
