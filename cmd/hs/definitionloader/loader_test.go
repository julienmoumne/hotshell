package definitionloader_test

import (
	"fmt"
	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
	"github.com/julienmoumne/hotshell/cmd/hs/definitionloader"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	dummyPayload   = []byte{0xFF}
	defaultMenu, _ = ioutil.ReadFile("../../../examples/default/default.hs.js")
	a              *assert.Assertions
	dl             definitionloader.Loader
	fs             vfs.Filesystem
	tests          = []testCase{
		// default menu explicitly requested
		{
			in:  in{true, "directory/sub/hs.js"},
			out: out{false, "default.hs.js", defaultMenu},
		},
		{
			ctx{fs: []fsEntry{{dir: "directory/sub", name: "hs.js"}}},
			in{true, "directory/sub/hs.js"},
			out{false, "default.hs.js", defaultMenu},
		},
		{
			ctx{fs: []fsEntry{{dir: "directory/sub", name: "hs.js"}}},
			in{true, ""},
			out{false, "default.hs.js", defaultMenu},
		},
		// default fallback
		{
			in:  in{false, ""},
			out: out{false, "default.hs.js", defaultMenu},
		},
		// "-f" option, file found
		{
			ctx{fs: []fsEntry{{dir: "directory/sub", name: "hs.js"}}},
			in{false, "directory/sub/hs.js"},
			out{false, "directory/sub/hs.js", dummyPayload},
		},
		// "-f" option, missing file
		{
			in:  in{false, "directory/sub/hs.js"},
			out: out{true, "", nil},
		},
		// "-f" option, directory
		{
			ctx{fs: []fsEntry{{dir: "directory/sub", name: "hs.js"}}},
			in{false, "directory/sub"},
			out{true, "", nil},
		},
		// "-f" option, remote file
		{
			in:  in{false, "https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/hs.js"},
			out: out{true, "", nil},
		},
		// "-f" option not specified, hs.js in current working directory
		{
			ctx{fs: []fsEntry{{dir: ".", name: "hs.js"}}},
			in{false, ""},
			out{false, "./hs.js", dummyPayload},
		},
	}
)

type (
	fsEntry struct {
		dir    string
		name   string
		remote bool
	}
	ctx struct {
		fs        []fsEntry
		userFound bool
	}
	in struct {
		loadDefaultMenu bool
		path            string
	}
	out struct {
		error    bool
		filename string
		content  []byte
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
	fs = memfs.Create()
	dl = definitionloader.Loader{}
	for _, entry := range t.fs {
		setupFsEntry(entry)
	}
}

func setupFsEntry(e fsEntry) {
	path := fmt.Sprintf("%s/%s", e.dir, e.name)
	vfs.MkdirAll(fs, e.dir, 0)
	vfs.WriteFile(fs, path, dummyPayload, 0)
}

func validateTest(t testCase) {
	d, err := dl.Load(fs, t.in.loadDefaultMenu, t.in.path)
	if t.out.error {
		a.NotNil(err)
	} else {
		a.Nil(err)
		a.Equal(t.out.filename, d.Filename)
		a.Equal(t.out.content, d.Dsl)
	}
}
