package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	remote := os.Args[1]
	tag := os.Args[2]
	dir := os.Args[3]
	e := os.RemoveAll(dir)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	repo, e := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      remote,
		Progress: os.Stdout,
	})
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	branch, e := repo.Tag(tag)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	w, e := repo.Worktree()
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	e = w.Checkout(&git.CheckoutOptions{
		Branch: branch.Name(),
	})
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	e = os.RemoveAll(filepath.Join(dir, ".git"))
	e = os.RemoveAll(filepath.Join(dir, ".github"))
	e = os.RemoveAll(filepath.Join(dir, ".vscode"))
	// e = os.RemoveAll(filepath.Join(dir, "cmd"))
	// check(e)
	e = os.Remove(filepath.Join(dir, "go.mod"))
	e = os.Remove(filepath.Join(dir, "go.sum"))
	e = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		file, e := os.ReadFile(path)
		if e != nil {
			return nil
		}
		_, filename := filepath.Split(path)
		if filename == "go.mod" ||
			filename == "go.sum" ||
			strings.HasSuffix(filename, "_test.go") ||
			strings.HasSuffix(filename, ".md") {
			e = os.Remove(path)
		} else {
			if filename == "Makefile" {
				file = []byte(strings.ReplaceAll(string(file),
					"github.com/lightningnetwork/lnd",
					"git-indra.lan/indra-labs/lnd/lnd"))
			}
			e = os.WriteFile(path,
				[]byte(strings.ReplaceAll(string(file),
					"\"github.com/lightningnetwork/lnd",
					"\"git-indra.lan/indra-labs/lnd/lnd")), 0755)
		}
		return nil
	})
	runCmdWithoutOutput("go", "mod", "tidy")
}

func runCmdWithoutOutput(cmd ...string) {
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Run()
}
