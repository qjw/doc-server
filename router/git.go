package router

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/qjw/kelly"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type User struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func gitFile(r *git.Repository, commitStr, name string) (io.ReadCloser, error) {
	// ... retrieving the commit object
	commit, err := r.CommitObject(plumbing.NewHash(commitStr))
	if err != nil {
		return nil, err
	}

	f, err := commit.File(name)
	if err != nil {
		return nil, err
	}

	return f.Reader()
}

func gitLatestFile(r *git.Repository, name string) (io.ReadCloser, error) {
	h, err := r.Head()
	if err != nil {
		return nil, err
	}

	return gitFile(r, h.Hash().String(), name)
}

func matchFile(r *git.Repository, commit *object.Commit, file string) error {
	p, err := commit.Parent(0)
	if err != nil {
		if err == object.ErrParentNotFound {
			pr, err := commit.Files()
			if err != nil {
				return err
			}

			for {
				cf, err := pr.Next()
				if err != nil {
					return fmt.Errorf("not match")
				}
				if cf.Name == file {
					return nil
				}
			}
		} else {
			return err
		}
	}

	pr, err := commit.Patch(p)
	if err != nil {
		return err
	}

	for _, v := range pr.FilePatches() {
		f, t := v.Files()
		if f != nil && f.Path() == file || t != nil && t.Path() == file {
			return nil
		}
	}
	return fmt.Errorf("not match")
}

type Global struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
}

type Spec struct {
	Groups []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Projects    []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Spec        string `json:"spec"`
		} `json:"projects"`
	} `json:"groups"`
	Global
}

type Current struct {
	Global
}

type Commit struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Log   string `json:"log"`
	Hash  string `json:"hash"`
	Time  string `json:"time"`
}

func initGitRouter(repository *git.Repository, r kelly.Router) {

	r.GET("/description", func(c *kelly.Context) {
		rd, err := gitLatestFile(repository, "spec.json")
		if err != nil {
			c.Abort(http.StatusInternalServerError, err.Error())
			return
		}

		var spec Current
		dec := json.NewDecoder(rd)
		if err := dec.Decode(&spec); err != nil {
			c.Abort(http.StatusInternalServerError, err.Error())
			return
		}

		c.WriteIndentedJson(http.StatusOK, kelly.H{
			"code":    0,
			"message": "OK",
			"data":    &spec,
		})
	})

	authR := r
	authR.GET("/spec", func(c *kelly.Context) {
		rd, err := gitLatestFile(repository, "spec.json")
		if err != nil {
			c.Abort(http.StatusInternalServerError, err.Error())
			return
		}

		var spec Spec
		dec := json.NewDecoder(rd)
		if err := dec.Decode(&spec); err != nil {
			c.Abort(http.StatusInternalServerError, err.Error())
			return
		}

		c.WriteIndentedJson(http.StatusOK, kelly.H{
			"code":    0,
			"message": "OK",
			"data":    &spec,
		})
	})

	authR.GET("/file", func(c *kelly.Context) {
		filename, err := c.GetQueryVarible("file")
		if err != nil {
			c.Abort(http.StatusUnprocessableEntity, err.Error())
			return
		}

		var rd io.ReadCloser = nil
		commit, err := c.GetQueryVarible("commit")
		if err != nil {
			rd, err = gitLatestFile(repository, filename)
		} else {
			rd, err = gitFile(repository, commit, filename)
		}

		if err != nil {
			c.Abort(http.StatusInternalServerError, err.Error())
			return
		}
		d, err := ioutil.ReadAll(rd)
		if err != nil {
			c.Abort(http.StatusInternalServerError, err.Error())
			return
		}

		c.WriteRawJson(http.StatusOK, d)
	})

	authR.GET("/history", func(c *kelly.Context) {
		filename, err := c.GetQueryVarible("file")
		if err != nil {
			c.Abort(http.StatusUnprocessableEntity, err.Error())
			return
		}

		ref, err := repository.Head()
		if err != nil {
			c.Abort(http.StatusInternalServerError, err.Error())
			return
		}
		cIter, err := repository.Log(&git.LogOptions{From: ref.Hash()})
		if err != nil {
			c.Abort(http.StatusInternalServerError, err.Error())
			return
		}

		res := []*Commit{}
		err = cIter.ForEach(func(commit *object.Commit) error {

			if matchFile(repository, commit, filename) == nil {
				res = append(res, &Commit{
					Name:  commit.Author.Name,
					Email: commit.Author.Email,
					Time:  commit.Author.When.Format("2006-01-02 03:04:05"),
					Hash:  commit.Hash.String(),
					Log:   commit.Message,
				})
			}

			return nil
		})

		c.WriteIndentedJson(http.StatusOK, kelly.H{
			"code":    0,
			"message": "OK",
			"data":    &res,
		})
	})
}
