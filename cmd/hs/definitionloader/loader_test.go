package definitionloader_test

import (
	"errors"
	"fmt"
	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
	"github.com/julienmoumne/hotshell/cmd/hs/definitionloader"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os/user"
	"testing"
)

var (
	dummyPayload   = []byte{0xFF}
	dummyUser      = user.User{HomeDir: "/home/user"}
	homeHotshell   = fsEntry{dir: dummyUser.HomeDir + "/.hs", name: "hs.js"}
	defaultMenu, _ = ioutil.ReadFile("../../../examples/default/default.hs.js")
	a              *assert.Assertions
	dl             definitionloader.Loader
	fs             vfs.Filesystem
	ug             *definitionloader.MockUserGetter
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
			in: in{false, "https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/hs.js"},
			out: out{true, "", nil},
		},
		// default locations
		{
			ctx{fs: []fsEntry{{dir: ".", name: "hs.js"}}},
			in{false, ""},
			out{false, "./hs.js", dummyPayload},
		},
		{
			ctx{fs: []fsEntry{{dir: ".", name: "hs.js"}, homeHotshell}, userFound: true},
			in{false, ""},
			out{false, "./hs.js", dummyPayload},
		},
		{
			ctx{fs: []fsEntry{homeHotshell}, userFound: true},
			in{false, ""},
			out{false, homeHotshell.dir + "/" + homeHotshell.name, dummyPayload},
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
	ug = new(definitionloader.MockUserGetter)
	fs = memfs.Create()
	dl = definitionloader.Loader{}
	for _, entry := range t.fs {
		setupFsEntry(entry)
	}
	setupCurrentUser(t)
}

func setupCurrentUser(t testCase) {
	var err error
	if !t.ctx.userFound {
		err = errors.New("user not found")
	}
	ug.On("Get").Return(&dummyUser, err)
}

func setupFsEntry(e fsEntry) {
	path := fmt.Sprintf("%s/%s", e.dir, e.name)
	vfs.MkdirAll(fs, e.dir, 0)
	vfs.WriteFile(fs, path, dummyPayload, 0)
}

func validateTest(t testCase) {
	d, err := dl.Load(fs, ug, t.in.loadDefaultMenu, t.in.path)
	if t.out.error {
		a.NotNil(err)
	} else {
		a.Nil(err)
		a.Equal(t.out.filename, d.Filename)
		a.Equal(t.out.content, d.Dsl)
	}
}