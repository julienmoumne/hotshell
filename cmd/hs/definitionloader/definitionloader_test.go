package definitionloader_test

import (
	"fmt"
	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
	"github.com/julienmoumne/hotshell/cmd/hs/definitionloader"
	"github.com/julienmoumne/hotshell/cmd/hs/fileloader"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	dummyPayload   = []byte{0xFF}
	defaultMenu, _ = ioutil.ReadFile("../../../examples/default/default.hs.js")
	a              *assert.Assertions
	dl             definitionloader.DefinitionLoader
	tests          = []testCase{
		// default menu explicitly requested
		{
			ctx{},
			in{true, "directory/sub/hs.js"},
			out{true, false, "default.hs.js", defaultMenu},
		},
		{
			ctx{"directory/sub", "hs.js"},
			in{true, "directory/sub/hs.js"},
			out{true, false, "default.hs.js", defaultMenu},
		},
		{
			ctx{"directory/sub", "hs.js"},
			in{true, ""},
			out{true, false, "default.hs.js", defaultMenu},
		},
		// "-f" option, file found
		{
			ctx{"directory/sub", "hs.js"},
			in{false, "directory/sub"},
			out{false, false, "directory/sub/hs.js", dummyPayload},
		},
		{
			ctx{"directory/sub", "hs.js"},
			in{false, "directory/sub/hs.js"},
			out{false, false, "directory/sub/hs.js", dummyPayload},
		},
		// "-f" option, missing file
		{
			ctx{},
			in{false, "directory/sub"},
			out{false, true, "", nil},
		},
		{
			ctx{},
			in{false, "directory/sub/hs.js"},
			out{false, true, "", nil},
		},
		// todo default locations with fallback to default
	}
)

type (
	ctx struct {
		dir      string
		filename string
	}
	in struct {
		loadDefaultMenu bool
		path            string
	}
	out struct {
		defaultMenuLoaded bool
		error             bool
		filename          string
		content           []byte
	}
	testCase struct {
		ctx
		in
		out
	}
)

func TestDefinitionLoader(t *testing.T) {
	a = assert.New(t)
	for _, t := range tests {
		runTest(t)
	}
}

func runTest(t testCase) {
	setupTest(t)
	validateTest(t)
}

func setupTest(t testCase) {
	fl := fileloader.MockFileLoader{}
	fs := memfs.Create()
	dl = definitionloader.DefinitionLoader{FileLoader: &fl, Fs: fs}

	if t.ctx == (ctx{}) {
		return
	}
	path := fmt.Sprintf("%s/%s", t.ctx.dir, t.ctx.filename)
	fl.On("Load", path).Return(dummyPayload, nil)
	vfs.MkdirAll(fs, t.ctx.dir, 0)
	vfs.WriteFile(fs, path, dummyPayload, 0)
}

func validateTest(t testCase) {
	d, err := dl.Load(t.in.loadDefaultMenu, t.in.path)
	if t.out.error {
		a.NotNil(err)
	} else {
		a.Nil(err)
		a.Equal(t.out.defaultMenuLoaded, d.DefaultMenuLoaded)
		a.Equal(t.out.filename, d.Filename)
		a.Equal(t.out.content, d.Dsl)
	}
}
