package glocal

import (
	"github.com/sohaha/zlsgo/zlog"
	"testing"
)

func TestBranches(t *testing.T) {
	repo := NewGitClient("https://gitee.com/moneygogo/wmscode.git", "17000006705@163.com", "111222333.")
	branches, err := repo.Branches()
	if err != nil {
		zlog.Errorf("git branches error %s\n", err.Error())
		return
	}
	zlog.Infof("branch count: %d \n", len(branches))
	repo.Clone("dev_69code")
}
