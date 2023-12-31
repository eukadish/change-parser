package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/object/commitgraph"
)

func main() {

	r, _ := git.PlainOpen("../client-verifier")
	cni := commitgraph.NewObjectCommitNodeIndex(r.Storer)

	ref, _ := r.Head()
	cn, _ := cni.Get(ref.Hash())

	// Ordered

	i := commitgraph.NewCommitNodeIterCTime(cn, nil, nil)

	i.ForEach(func(cn commitgraph.CommitNode) error {
		if cn == nil {
			return fmt.Errorf("nil commit node")
		}

		c, _ := cn.Commit()
		fmt.Println(fmt.Sprintf(" * Commit hash %s - %s: ", c.Hash, cn.CommitTime()))

		files, _ := c.Files()
		files.ForEach(func(f *object.File) error {
			if f == nil {
				return fmt.Errorf("nil file")
			}

			fmt.Println(f.Name)

			// TODO: Parse file imports

			// b, _ := git.Blame(c, f.Name)
			// fmt.Println(b.String())

			// lines := b.Lines
			// fmt.Println(len(lines))

			return nil
		})

		return nil
	})

	// Unordered

	// i, _ := r.CommitObjects()

	// i.ForEach(func(c *object.Commit) error {
	// 	if c == nil {
	// 		return fmt.Errorf("nil commit")
	// 	}

	// 	fmt.Println(c.Hash)
	// 	fmt.Println(" > hash files")

	// 	files, _ := c.Files()

	// 	files.ForEach(func(f *object.File) error {
	// 		if f == nil {
	// 			return fmt.Errorf("nil file")
	// 		}

	// 		fmt.Println(f.Name)

	// 		// b, _ := git.Blame(c1, f1.Name)
	// 		// fmt.Println(b.String())

	// 		// lines := b.Lines
	// 		// fmt.Println(len(lines))

	// 		return nil
	// 	})

	// 	return nil
	// })
}
