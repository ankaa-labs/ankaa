package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/gomega"

	"github.com/ankaa-labs/ankaa/test"
)

func TestStateMachineUnmarshal(t *testing.T) {
	g := NewWithT(t)

	f := path.Join(test.CurrentProjectPath() + "/hack/asl/map-with-null-itemspath.json")
	data, err := ioutil.ReadFile(f)
	g.Expect(err).ToNot(HaveOccurred())

	s := &StateMachine{}
	err = json.Unmarshal(data, s)
	g.Expect(err).ToNot(HaveOccurred())

	spew.Config.Indent = "\t"
	spew.Dump(s)
	fmt.Printf("%+v\n", spew.Sprint(s))
}
