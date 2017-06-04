package fileloader_test

import (
	"fmt"
	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
	"github.com/julienmoumne/hotshell/cmd/hs/fileloader"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"testing"
)

var (
	dummyPayload = []byte{0xFF}
	t            *testing.T
	a            *assert.Assertions
	loader       fileloader.Loader
	fs           vfs.Filesystem
	tests        = []testCase{
		// empty path
		{
			ctx: ctx{file: file{"directory/sub", "hs.js"}},
			out: out{true},
		},
		// directory
		{
			ctx{file: file{"directory/sub", "hs.js"}},
			in{"directory"},
			out{true},
		},
		// non-existent file
		{
			in: in{"directory/sub/hs.js"},
			out: out{true},
		},
		// valid local file
		{
			ctx{file: file{"directory/sub", "file"}},
			in{"directory/sub/file"},
			out{false},
		},
		//non-existing remote file
		{
			ctx{remote: true},
			in{"http://localhost/hs.js"},
			out{true},
		},
		// existing remote file
		{
			ctx{file: file{filename: "http://localhost/hs.js"}, remote: true},
			in{"http://localhost/hs.js"},
			out{false},
		},
	}
)

type (
	file struct {
		dir      string
		filename string
	}
	ctx struct {
		file
		remote bool
	}
	in struct {
		path string
	}
	out struct {
		error bool
	}
	testCase struct {
		ctx
		in
		out
	}
)

func TestFileLoader(_t *testing.T) {
	t = _t
	a = assert.New(_t)
	for _, t := range tests {
		runTest(t)
	}
}

func runTest(t testCase) {
	setupTest(t)
	validateTest(t)
	cleanupTest()
}

func setupTest(t testCase) {
	httpmock.Activate()
	fs = memfs.Create()
	loader = fileloader.Loader{}
	if t.ctx.file == (file{}) {
		return
	}
	if t.ctx.remote {
		httpmock.RegisterResponder(
			"GET",
			t.ctx.file.filename,
			httpmock.NewBytesResponder(200, dummyPayload),
		)
	} else {
		registerDirectory(t.ctx.file.dir)
		registerFile(fmt.Sprintf("%s/%s", t.ctx.file.dir, t.ctx.file.filename))
	}
}

func validateTest(t testCase) {
	content, err := loader.Load(fs, t.in.path)
	if t.out.error {
		a.NotNil(err)
	} else {
		a.Nil(err)
		a.Equal(dummyPayload, content)
	}
}

func cleanupTest() {
	httpmock.DeactivateAndReset()
}

func registerFile(path string) {
	raiseFatalIfErr(vfs.WriteFile(fs, path, dummyPayload, 0))
}

func registerDirectory(path string) {
	raiseFatalIfErr(vfs.MkdirAll(fs, path, 0))
}

func raiseFatalIfErr(err error) {
	if err != nil {
		t.Fatal(err)
	}
}
