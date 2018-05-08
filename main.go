package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	R "github.com/lixinio/doc-server/router"
	"github.com/qjw/kelly"
	"gopkg.in/src-d/go-git.v4"
)

func main() {
	c := make(chan os.Signal, 10)
	signal.Notify(c, syscall.SIGUSR2)

	config := R.InitConfig()
	var r *git.Repository = nil
	var err error = nil

	if len(config.GitOrigin) > 0 {
		R.RemoveContents(config.LocalDir)
		r, err = git.PlainClone(config.LocalDir, false, &git.CloneOptions{
			URL:      config.GitOrigin,
			Progress: os.Stdout,
		})
	} else {
		r, err = git.PlainOpen(config.LocalDir)
	}
	R.CheckIfError(err)
	w, err := r.Worktree()
	R.CheckIfError(err)

	go func(w *git.Worktree) {
		for {
			//阻塞直至有信号传入
			<-c
			fmt.Println("new pull request")
			err := w.Pull(&git.PullOptions{RemoteName: "origin"})
			if err != nil && err != git.NoErrAlreadyUpToDate {
				R.CheckIfError(err)
			}
			fmt.Println("pull request end")
		}
	}(w)

	router := kelly.New()
	R.Init(r, router, config)
	router.Run(fmt.Sprintf("0.0.0.0:%d", config.Port))
}
