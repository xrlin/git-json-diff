package git_json_diff

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func main() {
	filePath := flag.String("file", "", "file in repo to show diff")
	commit1 := flag.String("commit1", "HEAD", "commit id, as current version, to fetch the file content, default HEAD")
	commit2 := flag.String("commit2", "HEAD^", "commit id, as the old version, to fetch the file content , default HEAD^")
	format := flag.String("format", "ascii", "Diff Output Format (ascii, delta)")
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	jsonText1, err := RetriveFileContentWithCommitId(*filePath, *commit1)
	if err != nil {
		log.Fatalln(err)
	}
	jsonText2, err := RetriveFileContentWithCommitId(*filePath, *commit2)
	if err != nil {
		log.Fatalln(err)
	}

	diffString, err := Compare(jsonText1, jsonText2, *format)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(diffString)

}