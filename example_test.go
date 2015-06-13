package cfgp_test

import (
	"fmt"
	"log"

	"github.com/eraclitux/cfgp"
)

type myConf struct {
	Address         string
	Port            string
	Number_of_users int
}

func Example() {
	c := myConf{}
	err := cfgp.Parse("test_data/one.ini", &c)
	if err != nil {
		log.Fatal("Unable to parse configuration file", err)
	}
	fmt.Println("address:", c.Address)
	fmt.Println("port:", c.Port)
	fmt.Println("number_of_users:", c.Number_of_users)

	//Output:
	//address: localhost
	//port: 8080
	//number_of_users: 42

}
