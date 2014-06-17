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
		conf := Parse(testCase.Path)
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
