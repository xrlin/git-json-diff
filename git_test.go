package git_json_diff

import "testing"

func TestRetrieveFileContentWithCommitId(t *testing.T) {
	_, err := RetrieveFileContentWithCommitId("test.json", "HEAD")
	if err != nil {
		t.Errorf("failed with error: %v\n", err)
	}
}