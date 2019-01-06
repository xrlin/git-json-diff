package git_json_diff

import (
	"testing"
	"log"
)

func TestRetrieveFileContentWithCommitId(t *testing.T) {
	_, err := RetrieveFileContentWithCommitId("test.json", "HEAD")
	if err != nil {
		t.Errorf("failed with error: %v\n", err)
	}
}

func TestCompare(t *testing.T) {
	testCases := []struct{input1 string; input2 string}{
		{ `{
 "item0": [
    {"up": 1, "down": 2}
  ],
  "item1": [
    {"up": 1, "down": 2}
  ]}
	`,
			`
{
  "item1": [
    {"up": 1, "down": 2}
  ]
}
	`},
		{ `{
 "item0": [
    {"up": 1, "down": 2}
  ]}
	`,
			`
{
  "item0": [
	{"up": 3, "down": 4},
    {"up": 1, "down": 2}
  ]
}
	`},
	}
	for _, c := range testCases {
		diffStr, err := Compare(c.input1, c.input2, "ascii")
		log.Println(diffStr)
		if err != nil {
			t.Errorf("error: %v\n", err)
		}
	}

}