package filefetcher_test

import (
	"github.com/julienmoumne/hotshell/filefetcher"
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&FilefetcherTestSuite{})

type FilefetcherTestSuite struct {
	ff        filefetcher.Filefetcher
	fs        fileSystemStub
	webClient webClientStub
}

func (s *FilefetcherTestSuite) SetUpTest(c *C) {
	s.ff = filefetcher.Filefetcher{}
	s.fs = fileSystemStub{registeredFiles: make([]fileStub, 0)}
	s.webClient = webClientStub{}
	s.ff.Fs = &s.fs
	s.ff.DefaultHSFilename = "hs.js"
	s.ff.WebClient = &s.webClient
}

func (s *FilefetcherTestSuite) validateOutcome(c *C, path string, expectedContent []byte, expectedPath string, expectedError bool) {
	content, actualPath, err := s.ff.FetchFile(path)
	c.Check(content, DeepEquals, expectedContent)
	c.Check(actualPath, Equals, expectedPath)

	if expectedError {
		c.Check(err, Not(IsNil))
	} else {
		c.Check(err, IsNil)
	}
}

func (s *FilefetcherTestSuite) registerFileEntry(path string, content []byte, fileMode os.FileMode) {
	s.fs.registeredFiles = append(
		s.fs.registeredFiles,
		fileStub{path: path, content: content, fileInfo: fileInfoStub{fileMode}},
	)
}

func (s *FilefetcherTestSuite) registerFile(path string, content []byte) {
	s.registerFileEntry(path, content, 0)
}

func (s *FilefetcherTestSuite) registerDirectory(path string) {
	s.registerFileEntry(path, nil, os.ModeDir)
}

func (s *FilefetcherTestSuite) TestEmptyPath(c *C) {
	s.validateOutcome(c, "", nil, "", true)
}

func (s *FilefetcherTestSuite) TestUnknownLocalFile(c *C) {
	s.validateOutcome(c, "unknown-file", nil, "", true)
}

func (s *FilefetcherTestSuite) TestExistingLocalFile(c *C) {
	s.registerFile("existing-file", []byte{0xFF})
	s.validateOutcome(c, "existing-file", []byte{0xFF}, "existing-file", false)
}

func (s *FilefetcherTestSuite) TestKnownDirectoryWithoutMenu(c *C) {
	s.registerDirectory("existing-directory")
	s.validateOutcome(c, "existing-directory", nil, "", true)
}

func (s *FilefetcherTestSuite) TestKnownDirectoryWithMenu(c *C) {
	s.registerDirectory("existing-directory")
	s.registerFile("existing-directory/hs.js", []byte{0xFF})
	s.validateOutcome(c, "existing-directory", []byte{0xFF}, "existing-directory/hs.js", false)
}

func (s *FilefetcherTestSuite) TestUnknownRemoteFile(c *C) {
	s.validateOutcome(c, "http://localhost/hs.js", nil, "", true)
}

func (s *FilefetcherTestSuite) TestExistingRemoteFile(c *C) {
	s.webClient.registeredFile = "http://localhost/hs.js"
	s.webClient.registeredContent = []byte{0xFF}
	s.validateOutcome(c, "http://localhost/hs.js", []byte{0xFF}, "http://localhost/hs.js", false)
}
