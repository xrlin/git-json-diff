This simple tool is just to get the json diff in git repository conveniently.

## Install

```bash
go install github.com/xrlin/git-json-diff/cmd/git-json-diff/...
```
Besides, you can download the executable files in releases page directly.

```bash
$ git-json-diff
Usage of C:\Users\linxianrong\go\bin\git-json-diff.exe:
  -commit1 string
        commit id, as current version, to fetch the file content. (default "HEAD")
  -commit2 string
        commit id, as the old version, to fetch the file content. (default "HEAD~")
  -file string
        file in repo to show diff
  -format string
        Diff Output Format (ascii, delta) (default "ascii")

```

## Usage

```bash
git-json-diff -file=config/test.json
2019/01/07 11:32:39  {
-  "item0": [
-    : {
-      "down": 2,
-      "up": 1
-    }
-  ],
   "item1": [
     : {
       "down": 2,
       "up": 1
     }
   ]
 }

```

Run `git-json-diff` for more cli arguments.

