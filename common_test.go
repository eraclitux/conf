package cfgp

import (
	"testing"
)

type parseCaseResponse struct {
	Path           string
	ExpectedResult Conf
}

var parseTestCases []parseCaseResponse

func TestParse(t *testing.T) {
	for _, testCase := range parseTestCases {
		conf, err := Parse(testCase.Path)
		if err != nil {
			t.Fail()
		}
		if conf == nil || conf.ConfType != testCase.ExpectedResult.ConfType {
			t.Fail()
		}
	}
}

func init() {
	//var testCase parseCaseResponse
	testCase := parseCaseResponse{"test_data/one.ini", Conf{ConfType: "INI"}}
	parseTestCases = append(parseTestCases, testCase)
}
