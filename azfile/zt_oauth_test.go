package azfile_test

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/OmarKhatib158/azure-storage-file-go/azfile"

	chk "gopkg.in/check.v1"
)

type OAuthSuite struct{}

var _ = chk.Suite(&OAuthSuite{})

func getShareName() string {
	name := os.Getenv("SHARE_NAME")

	if name == "" {
		panic("SHARE_NAME environment vars must be set before running oauth tests")
	}

	return name
}

func (s *OAuthSuite) TestFileShareOAuth(c *chk.C) {
	fsu := getFSUWithOAuth()
	shareURL := fsu.NewShareURL(getShareName())

	fileSize := 1024 //1024 bytes

	file, _ := createNewFileFromShare(c, shareURL, int64(fileSize))
	defer delFile(c, file)

	contentR, contentD := getRandomDataAndReader(fileSize)

	_, err := file.UploadRange(context.Background(), 0, contentR, nil)
	c.Assert(err, chk.IsNil)

	resp, err := file.Download(context.Background(), 0, 1024, true)
	c.Assert(err, chk.IsNil)

	download, err := ioutil.ReadAll(resp.Body(azfile.RetryReaderOptions{}))
	c.Assert(err, chk.IsNil)
	c.Assert(download, chk.DeepEquals, contentD[:1024])
}
