package glocal

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/sohaha/zlsgo/zlog"
	"os"
)

type Client struct {
	url, username, password string
}

func NewGitClient(url string, username string, password string) *Client {
	return &Client{url: url, username: username, password: password}
}

const dir = "E:\\codestemp"

func (g *Client) Branches() (branches []string, err error) {
	zlog.Infof("git ls-remote --branches %s \n", g.url)
	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{g.url},
	})
	refs, err := rem.List(&git.ListOptions{
		PeelingOption: git.AppendPeeled,
		Auth: &http.BasicAuth{
			Username: g.username,
			Password: g.password,
		},
	})
	if err != nil {
		return nil, err
	}
	for _, ref := range refs {
		if ref.Name().IsBranch() {
			branches = append(branches, ref.Name().Short())
		}
	}
	return
}

func (g *Client) Clone(branch string) error {
	zlog.Infof("git clone %s branch:%s \n", g.url, branch)
	local := fmt.Sprintf("%s/%s", dir, branch)
	zlog.Infof("delete local project %s \n", local)
	os.RemoveAll(local)
	_, err := git.PlainClone(local, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: g.username,
			Password: g.password,
		},
		URL:           g.url,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		Progress:      zlog.Log,
	})
	return err
}
