package azfile_test

import (
	"context"
	"io/ioutil"

	"github.com/OmarKhatib158/azure-storage-file-go/azfile"

	//"home/omar/repos/ILDC/azure-storage-file-go/azfile/azfile"
	//"github.com/Azure/azure-storage-file-go/azfile"
	chk "gopkg.in/check.v1"
)

func createS2SFileSharesWithTokenCredential(c *chk.C, credential azfile.TokenCredential) (source, dest azfile.ShareURL) {
	fsu := getFSU()
	fsu.WithPipeline(azfile.NewPipeline(credential, azfile.PipelineOptions{}))

	source, dest = fsu.NewShareURL(azfile.NewUUID().String()), fsu.NewShareURL(azfile.NewUUID().String())

	_, err := source.Create(ctx, azfile.Metadata{}, 0)
	c.Assert(err, chk.IsNil)
	//_, err = dest.Create(ctx, azfile.Metadata{}, 0)
	//c.Assert(err, chk.IsNil)

	return
}

// TestUploadRangeFromURL check UploadRangeFromURL

func (s *aztestsSuite) TestFileShareS2SOAuth(c *chk.C) {
	SetEnvVarsazfilesoauth()
	ocred, err := getOAuthCredential("", "")
	c.Assert(err, chk.IsNil)
	source, _ := createS2SFileSharesWithTokenCredential(c, ocred)

	_, error := createNewFileFromShare(c, source, 2048)
	//sourceShare := source.NewDirectoryURL("SourceShare")

	//sourceShare := source.NewDirectoryURL("SourceShare")
	//sourceFile := sourceShare.NewFileURL("SourceFile")

	//_, err = sourceFile.UploadRangeFromURL(ctx, 0, strings.NewReader("data"), nil)
	//c.Assert(err, chk.IsNil)
	c.Assert(error, chk.IsNil)

}

func (s *aztestsSuite) TestFileShareOAuth(c *chk.C) {
	SetEnvVarsPlayGround()
	fsu := getFSUWithOauth()
	fsu = getFSU()
	ocred, _ := getOAuthCredential("", "")
	fsu.WithPipeline(azfile.NewPipeline(ocred, azfile.PipelineOptions{}))

	shareURL, _ := createNewShare(c, fsu)
	defer delShare(c, shareURL, azfile.DeleteSnapshotsOptionNone)

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
