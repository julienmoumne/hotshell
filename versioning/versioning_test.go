package versioning_test

import (
	"github.com/julienmoumne/hotshell/versioning"
	. "gopkg.in/check.v1"
	"testing"
)

func TestBuilder(t *testing.T) { TestingT(t) }

type TestVersioning struct{}

var _ = Suite(&TestVersioning{})

func (s *TestVersioning) TestVersion(c *C) {
	version, err := versioning.GetVersion()
	c.Check(err, IsNil)
	c.Check(string(version), Equals, "0.4.0")
}
