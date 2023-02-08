package azfile_test

import (
	"strings"

	"github.com/OmarKhatib158/azure-storage-file-go/azfile"

	//"home/omar/repos/ILDC/azure-storage-file-go/azfile/azfile"
	//"github.com/Azure/azure-storage-file-go/azfile"
	chk "gopkg.in/check.v1"
)

func createS2SFileSharesWithCredential(c *chk.C, credential azfile.Credential) (source, dest azfile.ShareURL) {
	fsu := getFSU()
	fsu.WithPipeline(azfile.NewPipeline(credential, azfile.PipelineOptions{}))

	source, dest = fsu.NewShareURL(azfile.newUUID().String()), fsu.NewShareURL(azfile.newUUID().String())

	_, err := source.Create(ctx, nil)
	c.Assert(err, chk.IsNil)
	_, err = dest.Create(ctx, nil)
	c.Assert(err, chk.IsNil)

	return
}

func (s *aztestsSuite) TestFileShareS2SOAuth(c *chk.C) {
	ocred, err := getOAuthCredential("", "")
	c.Assert(err, chk.IsNil)
	source, dest := createS2SFileSharesWithCredential(c, ocred)

	sourceShare := source.NewDirectoryURL("SourceShare")
	sourceFile := sourceShare.NewFileURL("SourceFile")

	_, err = sourceFile.UploadRange(ctx, 0, strings.NewReader("data"), nil)
	c.Assert(err, chk.IsNil)

	destShare := dest.NewDirectoryURL("DestShare")
	destFile := destShare.NewFileURL("DestFile")

}
