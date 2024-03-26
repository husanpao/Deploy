package mvn

import (
	"fmt"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zshell"
	"path/filepath"
	"strings"
)

type Client struct {
	dir   string
	IsWar bool
}

func NewClient(dir string) *Client {
	return &Client{dir: dir}
}

func (c *Client) doCheckJava() bool {
	_, outStr, errStr, err := zshell.Run("java -version")
	if err != nil {
		zlog.Errorf("Java Not Found.")
		return false
	}
	zlog.Debugf("%s\n%s\n", outStr, errStr)
	return strings.Index(outStr, "java version") != -1 || strings.Index(errStr, "java version") != -1
}
func (c *Client) Build() (flag bool) {
	if !c.doCheckJava() {
		return false
	}
	statusChan, _, err := zshell.CallbackRun(fmt.Sprintf("mvn  -f %s/pom.xml clean install -Dmaven.test.skip=true -Dcheckstyle.skip=true", c.dir), func(out string, isBasic bool) {
		zlog.Debugf(out)
		if !flag {
			flag = strings.Contains(out, "BUILD SUCCESS")
		}
	})
	if err != nil {
		return false
	}
	<-statusChan
	if flag {
		err = c._initWar()
		if err != nil {
			zlog.Errorf("Get Packaging Type Error:%s", err.Error())
			return false
		}
	}
	zlog.Infof("Build Source Success.")
	return
}
func (c *Client) _initWar() error {
	target := fmt.Sprintf("%s/target", c.dir)
	files, err := zfile.Ls(target)
	if err != nil {
		return err
	}
	for _, file := range files {
		c.IsWar = filepath.Ext(file) == ".war"
		if c.IsWar {
			break
		}
	}
	return nil
}
