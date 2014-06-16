package cfgp

import (
	"testing"
	"fmt"
)

type iniParseCaseResponse struct {
	Path string
	ExpectedResult Conf
}
var iniParseTestCases []iniParseCaseResponse

func TestParseINI(t *testing.T) {
	for _, testCase := range(iniParseTestCases) {
		conf := Conf{}
		conf.IniData = make(iniDataType)
		conf.ParseINI(testCase.Path)
		//FIXME not working
		//compareMaps(map[string][]map[string]string(testCase.ExpectedResult.IniData), conf.IniData)
		for k, v := range(conf.IniData) {
			fmt.Println(k, v)
		}
		t.Fail()
	}

}

//TODO create a separete pckg for this!
func compareMaps(o, p map[string]interface{}) bool {
	//TODO add more types
	if len(o) != len(p) {
		return false
	}
	for k, v := range o {
		switch vv := v.(type) {
		case string:
			//fmt.Println(k, "is string", vv)
			if vv != p[k].(string) {
				return false
			}
		case float64:
			//fmt.Println(k, "is float64", vv)
			if vv != p[k].(float64) {
				return false
			}
		case map[string]interface{}:
			//Recursion rocks
			return compareMaps(vv, p[k].(map[string]interface{}))
		default:
			fmt.Printf("%v is of a type I don't know how to handle %T(%v)\n", k, v, v)
			return false
		}
	}
	return true
}

func init() {
	expectedMap := iniDataType{"main": []map[string]string{
				{"one":"42"},
				{"three":"Zaphod"},
			},
	}
	iniTestCase := iniParseCaseResponse{
		"test_data/one.ini",
		Conf{IniData: expectedMap},
	}
	iniParseTestCases = append(iniParseTestCases, iniTestCase)
}
