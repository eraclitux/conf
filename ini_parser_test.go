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

type getKeyCase struct {
	Path           string
	TestSection    string
	TestKey        string
	ExpectedResult string
	ExpectedError  error
}

var getKeyCases []getKeyCase

type getSectionCase struct {
	Path           string
	TestSection    string
	ExpectedResult []map[string]string
	ExpectedError  error
}

var getSectionCases []getSectionCase

type getSectionsCase struct {
	Path           string
	ExpectedResult []string
	ExpectedError  error
}

var getSectionsCases []getSectionsCase

type isIniCase struct {
	Path           string
	ExpectedResult bool
}

var isIniCases []isIniCase

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
			t.FailNow()
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
			t.FailNow()
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

func TestGetKey(t *testing.T) {
	for _, testCase := range getKeyCases {
		conf, err := Parse(testCase.Path)
		if err != nil {
			if testing.Verbose() {
				fmt.Print("Unexpected error:", err)
			}
			t.FailNow()
		}
		testResult, err := conf.GetKey(testCase.TestSection, testCase.TestKey)
		// TODO compare errors's strings
		if err == nil && testCase.ExpectedError != nil ||
			err != nil && testCase.ExpectedError == nil {
			if testing.Verbose() {
				fmt.Printf("Case: %+v\n", testCase)
				fmt.Print("Errors mismatch:", err, testCase.ExpectedError)
			}
			t.FailNow()
		} else {
			if testResult != testCase.ExpectedResult {
				if testing.Verbose() {
					fmt.Printf("Case: %+v\n", testCase)
					fmt.Println("testResult:", testResult)
				}
				t.Fail()
			}
		}
	}
}

func TestGetSection(t *testing.T) {
	for _, testCase := range getSectionCases {
		conf, err := Parse(testCase.Path)
		if err != nil {
			if testing.Verbose() {
				fmt.Print("Unexpected error:", err)
			}
			t.FailNow()
		}
		testResult, err := conf.GetSection(testCase.TestSection)
		// TODO compare errors's strings
		if err == nil && testCase.ExpectedError != nil ||
			err != nil && testCase.ExpectedError == nil {
			if testing.Verbose() {
				fmt.Printf("Case: %+v\n", testCase)
				fmt.Print("Errors mismatch:", err, testCase.ExpectedError)
			}
			t.FailNow()
		} else {
			isEqual := reflect.DeepEqual(testCase.ExpectedResult, testResult)
			if !isEqual {
				if testing.Verbose() {
					fmt.Printf("Case: %+v\n", testCase)
					fmt.Println("testResult:", testResult)
				}
				t.Fail()
			}
		}
	}
}

func TestGetSections(t *testing.T) {
	for _, testCase := range getSectionsCases {
		conf, err := Parse(testCase.Path)
		if err != nil {
			if testing.Verbose() {
				fmt.Print("Unexpected error:", err)
			}
			t.FailNow()
		}
		testResult, err := conf.GetSections()
		// TODO compare errors's strings
		if err == nil && testCase.ExpectedError != nil ||
			err != nil && testCase.ExpectedError == nil {
			if testing.Verbose() {
				fmt.Printf("Case: %+v\n", testCase)
				fmt.Print("Errors mismatch:", err, testCase.ExpectedError)
			}
			t.FailNow()
		} else {
			isEqual := reflect.DeepEqual(testCase.ExpectedResult, testResult)
			if !isEqual {
				if testing.Verbose() {
					fmt.Printf("Case: %+v\n", testCase)
					fmt.Println("testResult:", testResult)
				}
				t.Fail()
			}
		}
	}
}

func TestIsIni(t *testing.T) {
	for _, testCase := range isIniCases {
		conf, err := Parse(testCase.Path)
		if err != nil {
			if testing.Verbose() {
				fmt.Printf("Error @ case: %+v\n", testCase)
				fmt.Println(err)
			}
			t.FailNow()
		}
		if conf.IsIni() != testCase.ExpectedResult {
			if testing.Verbose() {
				fmt.Printf("Case: %+v\n", testCase)
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

	// TestGetKey
	getKeyTestCase := getKeyCase{
		"test_data/one.ini",
		"questions",
		"wrong-answer",
		"43",
		nil,
	}
	getKeyCases = append(getKeyCases, getKeyTestCase)
	getKeyTestCase = getKeyCase{
		"test_data/one.cfg",
		"default",
		"the_answer",
		"42",
		nil,
	}
	getKeyCases = append(getKeyCases, getKeyTestCase)
	getKeyTestCase = getKeyCase{
		"test_data/one.ini",
		"questions",
		"cache",
		"",
		fmt.Errorf("Error"),
	}
	getKeyCases = append(getKeyCases, getKeyTestCase)

	// TestGetSection
	getSectionTestCase := getSectionCase{
		"test_data/one.ini",
		"questions",
		[]map[string]string{{"answer": "42"}, {"wrong-answer": "43"}},
		nil,
	}
	getSectionCases = append(getSectionCases, getSectionTestCase)
	getSectionTestCase = getSectionCase{
		"test_data/one.cfg",
		"default",
		[]map[string]string{{"the_answer": "42"}},
		nil,
	}
	getSectionCases = append(getSectionCases, getSectionTestCase)
	getSectionTestCase = getSectionCase{
		"test_data/one.ini",
		"default",
		nil,
		fmt.Errorf("Section default not found"),
	}
	getSectionCases = append(getSectionCases, getSectionTestCase)

	// TestGetSections
	getSectionsTestCase := getSectionsCase{
		"test_data/one.ini",
		[]string{"main", "questions"},
		nil,
	}
	getSectionsCases = append(getSectionsCases, getSectionsTestCase)
	// TestIsIni
	isIniTestCase := isIniCase{
		"test_data/one.ini",
		true,
	}
	isIniCases = append(isIniCases, isIniTestCase)
	isIniTestCase = isIniCase{
		"test_data/one.yaml",
		false,
	}
	// TODO Enable this after YAML implementation.
	//isIniCases = append(isIniCases, isIniTestCase)
	isIniTestCase = isIniCase{
		"test_data/one.cfg",
		true,
	}
	isIniCases = append(isIniCases, isIniTestCase)
}
