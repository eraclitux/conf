package cfgp_test

import (
	"fmt"
	"log"

	"github.com/eraclitux/cfgp"
)

type myConf struct {
	Address       string
	Port          string
	NumberOfUsers int `cfgp:"users,number of users,"`
	Daemon        bool
}

func Example() {
	c := myConf{}
	err := cfgp.Parse("test_data/one.ini", &c)
	if err != nil {
		log.Fatal("Unable to parse configuration file", err)
	}
	fmt.Println("address:", c.Address)
	fmt.Println("port:", c.Port)
	fmt.Println("number of users:", c.NumberOfUsers)

	//Output:
	//address: localhost
	//port: 8080
	//number of users: 42

}
