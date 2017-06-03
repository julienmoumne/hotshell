//go:generate mockery -name UserGetter -inpkg -case underscore
package definitionloader_test

import (
	"errors"
	"fmt"
	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
	"github.com/julienmoumne/hotshell/cmd/hs/definitionloader"
	"github.com/julienmoumne/hotshell/cmd/hs/fileloader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"os/user"
	"testing"
)

var (
	dummyPayload   = []byte{0xFF}
	dummyUser      = user.User{HomeDir: "/home/user"}
	homeHotshell   = fsEntry{dummyUser.HomeDir + "/.hs", "hs.js"}
	defaultMenu, _ = ioutil.ReadFile("../../../examples/default/default.hs.js")
	a              *assert.Assertions
	dl             definitionloader.DefinitionLoader
	tests          = []testCase{
		// default menu explicitly requested
		{
			in:  in{true, "directory/sub/hs.js"},
			out: out{false, "default.hs.js", defaultMenu},
		},
		{
			ctx{fs: []fsEntry{{"directory/sub", "hs.js"}}},
			in{true, "directory/sub/hs.js"},
			out{false, "default.hs.js", defaultMenu},
		},
		{
			ctx{fs: []fsEntry{{"directory/sub", "hs.js"}}},
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
			ctx{fs: []fsEntry{{"directory/sub", "hs.js"}}},
			in{false, "directory/sub"},
			out{false, "directory/sub/hs.js", dummyPayload},
		},
		{
			ctx{fs: []fsEntry{{"directory/sub", "hs.js"}}},
			in{false, "directory/sub/hs.js"},
			out{false, "directory/sub/hs.js", dummyPayload},
		},
		// "-f" option, missing file
		{
			in:  in{false, "directory/sub"},
			out: out{true, "", nil},
		},
		{
			in:  in{false, "directory/sub/hs.js"},
			out: out{true, "", nil},
		},
		// default locations
		{
			ctx{fs: []fsEntry{{".", "hs.js"}}},
			in{false, ""},
			out{false, "./hs.js", dummyPayload},
		},
		{
			ctx{fs: []fsEntry{{".", "hs.js"}, homeHotshell}, userFound: true},
			in{false, ""},
			out{false, "./hs.js", dummyPayload},
		},
		{
			ctx{fs: []fsEntry{homeHotshell}, userFound: true},
			in{false, ""},
			out{false, homeHotshell.dir + "/" + homeHotshell.filename, dummyPayload},
		},
	}
)

type (
	fsEntry struct {
		dir      string
		filename string
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
	fl := new(fileloader.MockFileLoader)
	ug := new(definitionloader.MockUserGetter)
	dl = definitionloader.DefinitionLoader{FileLoader: fl, Fs: memfs.Create(), UserGetter: ug}
	for _, entry := range t.fs {
		setupFsEntry(fl, entry)
	}
	fl.On("Load", mock.AnythingOfType("string")).Return(nil, errors.New("file not found"))
	setupCurrentUser(t, ug)
}

func setupCurrentUser(t testCase, ug *definitionloader.MockUserGetter) {
	var err error
	if !t.ctx.userFound {
		err = errors.New("user not found")
	}
	ug.On("Get").Return(&dummyUser, err)
}

func setupFsEntry(fl *fileloader.MockFileLoader, entry fsEntry) {
	path := fmt.Sprintf("%s/%s", entry.dir, entry.filename)
	fl.On("Load", path).Return(dummyPayload, nil)
	vfs.MkdirAll(dl.Fs, entry.dir, 0)
	vfs.WriteFile(dl.Fs, path, dummyPayload, 0)
}

func validateTest(t testCase) {
	d, err := dl.Load(t.in.loadDefaultMenu, t.in.path)
	if t.out.error {
		a.NotNil(err)
	} else {
		a.Nil(err)
		a.Equal(t.out.filename, d.Filename)
		a.Equal(t.out.content, d.Dsl)
	}
}
