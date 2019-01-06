package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"github.com/xrlin/git-json-diff"
)

func main() {
	filePath := flag.String("file", "", "file in repo to show diff")
	commit1 := flag.String("commit1", "HEAD", "commit id, as current version, to fetch the file content.")
	commit2 := flag.String("commit2", "HEAD~", "commit id, as the old version, to fetch the file content.")
	format := flag.String("format", "ascii", "Diff Output Format (ascii, delta)")
	flag.Parse()

	if *filePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	jsonText1, err := git_json_diff.RetrieveFileContentWithCommitId(*filePath, *commit1)
	if err != nil {
		log.Fatalln(err)
	}
	jsonText2, err := git_json_diff.RetrieveFileContentWithCommitId(*filePath, *commit2)
	if err != nil {
		log.Fatalln(err)
	}

	diffString, err := git_json_diff.Compare(jsonText1, jsonText2, *format)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(diffString)

}