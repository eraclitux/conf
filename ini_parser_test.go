package cfgp

import (
	"fmt"
	"reflect"
	"testing"
)

type iniParseCaseResponse struct {
	Path           string
	ExpectedResult Conf
}

var iniParseTestCases []iniParseCaseResponse

type hasSectionCase struct {
	Path           string
	TestSection    string
	ExpectedResult bool
}

var hasSectionCases []hasSectionCase

func TestParseINI(t *testing.T) {
	for _, testCase := range iniParseTestCases {
		conf := Conf{}
		conf.IniData = make(iniDataType)
		conf.parseINI(testCase.Path)
		isEqual := reflect.DeepEqual(testCase.ExpectedResult.IniData, conf.IniData)
		if !isEqual {
			fmt.Println("Expect:", testCase.ExpectedResult.IniData)
			fmt.Println("Got:", conf.IniData)
			t.Fail()
		}
	}

}

func TestHasSection(t *testing.T) {
	for _, testCase := range hasSectionCases {
		conf := Parse(testCase.Path)
		testResult := conf.HasSection(testCase.TestSection)
		if testResult != testCase.ExpectedResult {
			fmt.Println("Case:", testCase)
			t.Fail()
		}
	}
}

func init() {
	// Enable verbose output
	debug = false

	// TestParseINI
	expectedMap := iniDataType{"main": []map[string]string{
		{"one": "42"},
		{"three": "Zaphod"},
	},
		"questions": []map[string]string{
			{"answer": "42"},
			{"wrong-answer": "43"},
		},
	}
	iniTestCase := iniParseCaseResponse{
		"test_data/one.ini",
		Conf{IniData: expectedMap},
	}
	iniParseTestCases = append(iniParseTestCases, iniTestCase)

	// TestHasSection
	hasSectionTestCase := hasSectionCase{
		"test_data/one.ini",
		"questions",
		true,
	}
	hasSectionCases = append(hasSectionCases, hasSectionTestCase)
	hasSectionTestCase = hasSectionCase{
		"test_data/one.ini",
		"cache",
		false,
	}
	hasSectionCases = append(hasSectionCases, hasSectionTestCase)
}
