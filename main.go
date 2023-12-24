package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))

	os.Exit(1)
}

func main() {

	// We instantiate a new repository targeting the given path (the .git folder)
	// _, err := git.PlainOpen(".")
	r, _ := git.PlainOpen(".")

	// ref, _ := r.Head()

	// c, _ := object.GetCommit(r.Storer, ref.Hash())

	// fmt.Println(c.Message)

	// b, _ := git.Blame(c, "test/unit/FundAdapter.t.sol")

	// fmt.Println(b.String())

	// lines := b.Lines

	// fmt.Println(len(lines))

	// fmt.Println(lines[0])
	// fmt.Println(lines[1])
	// fmt.Println(lines[2])
	// fmt.Println(lines[3])
	// fmt.Println(lines[4])
	// fmt.Println(lines[5])
	// fmt.Println(lines[6])
	// fmt.Println(lines[7])
	// fmt.Println(lines[8])
	// fmt.Println(lines[9])

	i, _ := r.CommitObjects()

	c1, _ := i.Next()
	// c2, _ := i.Next()

	files, _ := c1.Files()

	f1, _ := files.Next()

	fmt.Println(f1.Name)

	b, _ := git.Blame(c1, f1.Name)

	fmt.Println(b.String())

	lines := b.Lines

	fmt.Println(len(lines))

	fmt.Println(lines[0])
	// fmt.Println(lines[1])
	// fmt.Println(lines[2])
	// fmt.Println(lines[3])
	// fmt.Println(lines[4])
	// fmt.Println(lines[5])
	// fmt.Println(lines[6])
	// fmt.Println(lines[7])
	// fmt.Println(lines[8])
	// fmt.Println(lines[9])

	f2, _ := files.Next()

	fmt.Println(f2.Name)

	// diffs := diff.Do(c1.String(), c2.String())

	// length := len(f)
	// for x := 0; x < length; x++ {
	// 	fmt.Println(f[x].asdf)
	// 	// fmt.Println(diffs[x].Text)

	// 	// fmt.Println(diffs[x])
	// 	// fmt.Println(diffs[x].Type.String())
	// }

	// diffs := diff.Do(c1.String(), c2.String())

	// length := len(diffs)
	// for x := 0; x < length; x++ {
	// 	fmt.Println(diffs[x].Type)
	// 	fmt.Println(diffs[x].Text)

	// 	fmt.Println(diffs[x])
	// 	fmt.Println(diffs[x].Type.String())
	// }

	// fmt.Println("Hash:")
	// fmt.Println(d.Hash)
	// fmt.Println(i)
	// fmt.Println(" - - - - - - - - - - ")

	// ... retrieves the commit history
	// since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	// until := time.Date(2019, 7, 30, 0, 0, 0, 0, time.UTC)
	// cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	// CheckIfError(err)

	// // ... just iterates over the commits, printing it
	// err = cIter.ForEach(func(c *object.Commit) error {
	// 	fmt.Println(c)

	// 	return nil
	// })
	// CheckIfError(err)

	// git.diffs

	// fmt.Println(err)
	// c, _ := r.Config()

	// fmt.Println(c)
	// fmt.Println(c.Core.IsBare)

	// fmt.Println(c.Core.Worktree)
	// fmt.Println(c.Core.CommentChar)

	// fmt.Println("Hello, World!")

	// fmt.Println(c.Init.DefaultBranch)
}
