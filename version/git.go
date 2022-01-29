package version

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"path/filepath"
	"strings"
)

func GetCurrentBranchFromRepository(repository *git.Repository) (string, error) {
	branchRefs, err := repository.Branches()
	if err != nil {
		return "", err
	}

	headRef, err := repository.Head()
	if err != nil {
		return "", err
	}

	var currentBranchName string
	err = branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
		if branchRef.Hash() == headRef.Hash() {
			currentBranchName = branchRef.Name().String()

			return nil
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(currentBranchName, "refs/heads/"), nil
}

func GetCurrentCommitFromRepository(repository *git.Repository) (string, error) {
	headRef, err := repository.Head()
	if err != nil {
		return "", err
	}
	headSha := headRef.Hash().String()

	return headSha[0:9], nil
}

func GetTagName(tag *plumbing.Reference) string {
	if tag != nil {
		return strings.TrimPrefix(tag.Name().String(), "refs/tags/")
	}

	return "v0.0.0"
}

func GetLatestTagFromRepository(repository *git.Repository) (*plumbing.Reference, error) {
	tagRefs, err := repository.Tags()
	if err != nil {
		return nil, err
	}

	var latestTagCommit *object.Commit
	var latestTag *plumbing.Reference
	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		revision := plumbing.Revision(tagRef.Name().String())
		tagCommitHash, err := repository.ResolveRevision(revision)
		if err != nil {
			return err
		}

		commit, err := repository.CommitObject(*tagCommitHash)
		if err != nil {
			return err
		}

		if latestTagCommit == nil {
			latestTagCommit = commit
			latestTag = tagRef
		}

		if commit.Committer.When.After(latestTagCommit.Committer.When) {
			latestTagCommit = commit
			latestTag = tagRef
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return latestTag, nil
}

func CalculateAheadCommits(repository *git.Repository, targetRef *plumbing.Reference) (int, error) {
	head, err := repository.Head()
	if err != nil {
		return -1, err
	}
	cIter, err := repository.Log(&git.LogOptions{
		// All:   true,
		From:  head.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return -1, err
	}

	var count = 0
	err = cIter.ForEach(func(c *object.Commit) error {
		if targetRef != nil {
			t_hash := targetRef.Hash()
			if bytes.Equal(c.Hash[:], t_hash[:]) {
				// Found!
				return storer.ErrStop
			}
		}
		count++
		// No luck continue searching.
		return nil
	})

	return count, nil
}

func Version(location string) (string, error) {
	// Opening the repo
	cwd, err := filepath.Abs(location)
	if err != nil {
		return "", err
	}
	r, err := git.PlainOpen(cwd)
	if err != nil {
		return "", err
	}

	branchName, err := GetCurrentBranchFromRepository(r)
	if err != nil {
		return "", err
	}

	t, err := GetLatestTagFromRepository(r)
	if err != nil {
		return "", err
	}
	tagName := GetTagName(t)

	headHash, err := GetCurrentCommitFromRepository(r)
	if err != nil {
		return "", err
	}

	c, err := CalculateAheadCommits(r, t)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%d-%s (%s)", tagName, c, headHash, branchName), nil
}
