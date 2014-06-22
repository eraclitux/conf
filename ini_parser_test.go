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

type hasKeyCase struct {
	Path           string
	TestSection    string
	TestKey        string
	ExpectedResult bool
}

var hasKeyCases []hasKeyCase

func TestParseINI(t *testing.T) {
	for _, testCase := range iniParseTestCases {
		conf := Conf{}
		conf.IniData = make(iniDataType)
		conf.parseINI(testCase.Path)
		isEqual := reflect.DeepEqual(testCase.ExpectedResult.IniData, conf.IniData)
		if !isEqual {
			if testing.Verbose() {
				fmt.Println("Expect:", testCase.ExpectedResult.IniData)
				fmt.Println("Got:", conf.IniData)
			}
			t.Fail()
		}
	}

}

func TestHasSection(t *testing.T) {
	for _, testCase := range hasSectionCases {
		conf, err := Parse(testCase.Path)
		if err != nil {
			t.Fail()
		}
		testResult := conf.HasSection(testCase.TestSection)
		if testResult != testCase.ExpectedResult {
			if testing.Verbose() {
				fmt.Printf("Case: %+v\n", testCase)
				fmt.Println("testResult:", testResult)
			}
			t.Fail()
		}
	}
}

func TestHasKey(t *testing.T) {
	for _, testCase := range hasKeyCases {
		conf, err := Parse(testCase.Path)
		if err != nil {
			t.Fail()
		}
		testResult := conf.HasKey(testCase.TestSection, testCase.TestKey)
		if testResult != testCase.ExpectedResult {
			if testing.Verbose() {
				fmt.Printf("Case: %+v\n", testCase)
				fmt.Println("testResult:", testResult)
			}
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

	// TestHasKey
	hasKeyTestCase := hasKeyCase{
		"test_data/one.ini",
		"questions",
		"wrong-answer",
		true,
	}
	hasKeyCases = append(hasKeyCases, hasKeyTestCase)
	hasKeyTestCase = hasKeyCase{
		"test_data/one.ini",
		"questions",
		"cache",
		false,
	}
	hasKeyCases = append(hasKeyCases, hasKeyTestCase)
}
