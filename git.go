package git_json_diff

import (
	"os/exec"
	"log"
	diff "github.com/yudai/gojsondiff"
	"fmt"
	"encoding/json"
	"github.com/yudai/gojsondiff/formatter"
	)

func isGitInstall() bool {
	if _, err := exec.LookPath("git"); err != nil {
		return false
	}
	return true
}

func RetrieveFileContentWithCommitId(filePath, commitId string) (ret string, err error) {
	if !isGitInstall() {
		log.Fatalln("cannot find git command")
	}
	cmd := exec.Command("git", "show",  commitId + ":" + filePath)
	stdout, err := cmd.CombinedOutput()
	defer func() {
		if err != nil {
			fmt.Printf("running %s with args %s riase error: %s\n", cmd.Path, cmd.Args, stdout)
		}
	}()

	if err != nil {
		return "", err
	}
	return string(stdout), err
}

func checkJSONUnmarshal(jsonText string) error {
	var m  map[string]interface{}
	err := json.Unmarshal([]byte(jsonText), &m)
	return err
}

func Compare(jsonText1, jsonText2, outFormat string) (string, error) {
	if jsonText1 == "" {
		jsonText1 = "{}"
	}

	if jsonText2 == "" {
		jsonText2 = "{}"
	}

	if err := checkJSONUnmarshal(jsonText1); err != nil {
		return "", fmt.Errorf("json content:\n %s \n decode failed with error: %s\n", jsonText1, err)
	}

	if err := checkJSONUnmarshal(jsonText2); err != nil {
		return "", fmt.Errorf("json content:\n %s \n decode failed with error: %s\n", jsonText2, err)
	}

	differ := diff.New()
	d, err := differ.Compare([]byte(jsonText1), []byte(jsonText2))
	if err != nil {
		return "", err
	}
	if !d.Modified() {
		return "", err
	}
	return formatDiff(d, jsonText1, outFormat, true)
}

func formatDiff(d diff.Diff, prevJsonText string, format string, enableColor bool) (ret string, err error) {
	if format == "ascii" {
		var aJson map[string]interface{}
		json.Unmarshal([]byte(prevJsonText), &aJson)

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       enableColor,
		}

		formatTransfer := formatter.NewAsciiFormatter(aJson, config)
		return formatTransfer.Format(d)
	}

	if format == "delta" {
		formatTransfer := formatter.NewDeltaFormatter()
		return formatTransfer.Format(d)
	}
	return "", fmt.Errorf("unkonw format %s", format)
}