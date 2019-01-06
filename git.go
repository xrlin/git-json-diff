package git_json_diff

import (
	"os/exec"
	"log"
	"io/ioutil"
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

func RetriveFileContentWithCommitId(filePath, commitid string) (string, error) {
	if isGitInstall() {
		log.Fatalln("cannot find git command")
	}
	cmd := exec.Command("git", "show " + commitid + ":" + filePath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		stdout.Close()
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(stdout)
	return string(b), err
}

func Compare(jsonText1, jsonText2, outFormat string) (string, error) {
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