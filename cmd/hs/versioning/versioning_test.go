package versioning_test

import (
	"github.com/julienmoumne/hotshell/cmd/hs/versioning"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	version, err := versioning.GetVersion()
	a := assert.New(t)
	a.Nil(err)
	a.Equal("0.5.0", string(version))
}
