package mvn

import (
	"fmt"
	"testing"
)

func Test_doCheck(t *testing.T) {
	m := NewClient("D:/Projects/IdeaProjects/TestDeploy")
	m.Build()
	fmt.Println(m.IsWar)
}
