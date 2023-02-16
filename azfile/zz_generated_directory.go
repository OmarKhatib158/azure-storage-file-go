package azfile

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Azure/azure-pipeline-go/pipeline"
)

// directoryClient is the client for the Directory methods of the Azfile service.
type directoryClient struct {
	managementClient
}

// newDirectoryClient creates an instance of the directoryClient client.
func newDirectoryClient(url url.URL, p pipeline.Pipeline) directoryClient {
	return directoryClient{newManagementClient(url, p)}
}

// Create creates a new directory under the specified share or parent directory.
//
// fileAttributes is if specified, the provided file attributes shall be set. Default value: ‘Archive’ for file and
// ‘Directory’ for directory. ‘None’ can also be specified as default. fileCreationTime is creation time for the
// file/directory. Default value: Now. fileLastWriteTime is last write time for the file/directory. Default value: Now.
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
// filePermission is if specified the permission (security descriptor) shall be set for the directory/file. This header
// can be used if Permission size is <= 8KB, else x-ms-file-permission-key header shall be used. Default value:
// Inherit. If SDDL is specified as input, it must have owner, group and dacl. Note: Only one of the
// x-ms-file-permission or x-ms-file-permission-key should be specified. filePermissionKey is key of the permission to
// be set for the directory/file. Note: Only one of the x-ms-file-permission or x-ms-file-permission-key should be
// specified.
func (client directoryClient) Create(ctx context.Context, fileAttributes string, fileCreationTime string, fileLastWriteTime string, timeout *int32, metadata map[string]string, filePermission *string, filePermissionKey *string) (*DirectoryCreateResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createPreparer(fileAttributes, fileCreationTime, fileLastWriteTime, timeout, metadata, filePermission, filePermissionKey)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DirectoryCreateResponse), err
}

// createPreparer prepares the Create request.
func (client directoryClient) createPreparer(fileAttributes string, fileCreationTime string, fileLastWriteTime string, timeout *int32, metadata map[string]string, filePermission *string, filePermissionKey *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "directory")
	req.URL.RawQuery = params.Encode()
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if filePermission != nil {
		req.Header.Set("x-ms-file-permission", *filePermission)
	}
	if filePermissionKey != nil {
		req.Header.Set("x-ms-file-permission-key", *filePermissionKey)
	}
	req.Header.Set("x-ms-file-attributes", fileAttributes)
	req.Header.Set("x-ms-file-creation-time", fileCreationTime)
	req.Header.Set("x-ms-file-last-write-time", fileLastWriteTime)
	req.Header.Set("x-ms-file-request-intent", "backup")

	return req, nil
}

// createResponder handles the response to the Create request.
func (client directoryClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &DirectoryCreateResponse{rawResponse: resp.Response()}, err
}

// Delete removes the specified empty directory. Note that the directory must be empty before it can be deleted.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client directoryClient) Delete(ctx context.Context, timeout *int32) (*DirectoryDeleteResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.deletePreparer(timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DirectoryDeleteResponse), err
}

// deletePreparer prepares the Delete request.
func (client directoryClient) deletePreparer(timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("DELETE", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "directory")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	req.Header.Set("x-ms-file-request-intent", "backup")

	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client directoryClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &DirectoryDeleteResponse{rawResponse: resp.Response()}, err
}

// ForceCloseHandles closes all handles open for given directory.
//
// handleID is specifies handle ID opened on the file or directory to be closed. Asterix (‘*’) is a wildcard that
// specifies all handles. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> marker is a string value that identifies the portion of the list to be
// returned with the next list operation. The operation returns a marker value within the response body if the list
// returned was not complete. The marker value may then be used in a subsequent call to request the next set of list
// items. The marker value is opaque to the client. sharesnapshot is the snapshot parameter is an opaque DateTime value
// that, when present, specifies the share snapshot to query. recursive is specifies operation should apply to the
// directory specified in the URI, its files, its subdirectories and their files.
func (client directoryClient) ForceCloseHandles(ctx context.Context, handleID string, timeout *int32, marker *string, sharesnapshot *string, recursive *bool) (*DirectoryForceCloseHandlesResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.forceCloseHandlesPreparer(handleID, timeout, marker, sharesnapshot, recursive)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.forceCloseHandlesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DirectoryForceCloseHandlesResponse), err
}

// forceCloseHandlesPreparer prepares the ForceCloseHandles request.
func (client directoryClient) forceCloseHandlesPreparer(handleID string, timeout *int32, marker *string, sharesnapshot *string, recursive *bool) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	if marker != nil && len(*marker) > 0 {
		params.Set("marker", *marker)
	}
	if sharesnapshot != nil && len(*sharesnapshot) > 0 {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	params.Set("comp", "forceclosehandles")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-handle-id", handleID)
	if recursive != nil {
		req.Header.Set("x-ms-recursive", strconv.FormatBool(*recursive))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	req.Header.Set("x-ms-file-request-intent", "backup")

	return req, nil
}

// forceCloseHandlesResponder handles the response to the ForceCloseHandles request.
func (client directoryClient) forceCloseHandlesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &DirectoryForceCloseHandlesResponse{rawResponse: resp.Response()}, err
}

// GetProperties returns all system properties for the specified directory, and can also be used to check the existence
// of a directory. The data returned does not include the files in the directory or any subdirectories.
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client directoryClient) GetProperties(ctx context.Context, sharesnapshot *string, timeout *int32) (*DirectoryGetPropertiesResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPropertiesPreparer(sharesnapshot, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DirectoryGetPropertiesResponse), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client directoryClient) getPropertiesPreparer(sharesnapshot *string, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if sharesnapshot != nil && len(*sharesnapshot) > 0 {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "directory")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	req.Header.Set("x-ms-file-request-intent", "backup")
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client directoryClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &DirectoryGetPropertiesResponse{rawResponse: resp.Response()}, err
}

// ListFilesAndDirectoriesSegment returns a list of files or directories under the specified share or directory. It
// lists the contents only for a single level of the directory hierarchy.
//
// prefix is filters the results to return only entries whose name begins with the specified prefix. sharesnapshot is
// the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot to query. marker
// is a string value that identifies the portion of the list to be returned with the next list operation. The operation
// returns a marker value within the response body if the list returned was not complete. The marker value may then be
// used in a subsequent call to request the next set of list items. The marker value is opaque to the client.
// maxresults is specifies the maximum number of entries to return. If the request does not specify maxresults, or
// specifies a value greater than 5,000, the server will return up to 5,000 items. timeout is the timeout parameter is
// expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client directoryClient) ListFilesAndDirectoriesSegment(ctx context.Context, prefix *string, sharesnapshot *string, marker *string, maxresults *int32, timeout *int32) (*ListFilesAndDirectoriesSegmentResponse, error) {
	if err := validate([]validation{
		{targetValue: maxresults,
			constraints: []constraint{{target: "maxresults", name: null, rule: false,
				chain: []constraint{{target: "maxresults", name: inclusiveMinimum, rule: 1, chain: nil}}}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.listFilesAndDirectoriesSegmentPreparer(prefix, sharesnapshot, marker, maxresults, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listFilesAndDirectoriesSegmentResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ListFilesAndDirectoriesSegmentResponse), err
}

// listFilesAndDirectoriesSegmentPreparer prepares the ListFilesAndDirectoriesSegment request.
func (client directoryClient) listFilesAndDirectoriesSegmentPreparer(prefix *string, sharesnapshot *string, marker *string, maxresults *int32, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if prefix != nil && len(*prefix) > 0 {
		params.Set("prefix", *prefix)
	}
	if sharesnapshot != nil && len(*sharesnapshot) > 0 {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	if marker != nil && len(*marker) > 0 {
		params.Set("marker", *marker)
	}
	if maxresults != nil {
		params.Set("maxresults", strconv.FormatInt(int64(*maxresults), 10))
	}
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "directory")
	params.Set("comp", "list")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	req.Header.Set("x-ms-file-request-intent", "backup")
	return req, nil
}

// listFilesAndDirectoriesSegmentResponder handles the response to the ListFilesAndDirectoriesSegment request.
func (client directoryClient) listFilesAndDirectoriesSegmentResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ListFilesAndDirectoriesSegmentResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, err
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListHandles lists handles for directory.
//
// marker is a string value that identifies the portion of the list to be returned with the next list operation. The
// operation returns a marker value within the response body if the list returned was not complete. The marker value
// may then be used in a subsequent call to request the next set of list items. The marker value is opaque to the
// client. maxresults is specifies the maximum number of entries to return. If the request does not specify maxresults,
// or specifies a value greater than 5,000, the server will return up to 5,000 items. timeout is the timeout parameter
// is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> sharesnapshot is the snapshot parameter is an opaque DateTime value that,
// when present, specifies the share snapshot to query. recursive is specifies operation should apply to the directory
// specified in the URI, its files, its subdirectories and their files.
func (client directoryClient) ListHandles(ctx context.Context, marker *string, maxresults *int32, timeout *int32, sharesnapshot *string, recursive *bool) (*ListHandlesResponse, error) {
	if err := validate([]validation{
		{targetValue: maxresults,
			constraints: []constraint{{target: "maxresults", name: null, rule: false,
				chain: []constraint{{target: "maxresults", name: inclusiveMinimum, rule: 1, chain: nil}}}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.listHandlesPreparer(marker, maxresults, timeout, sharesnapshot, recursive)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listHandlesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ListHandlesResponse), err
}

// listHandlesPreparer prepares the ListHandles request.
func (client directoryClient) listHandlesPreparer(marker *string, maxresults *int32, timeout *int32, sharesnapshot *string, recursive *bool) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if marker != nil && len(*marker) > 0 {
		params.Set("marker", *marker)
	}
	if maxresults != nil {
		params.Set("maxresults", strconv.FormatInt(int64(*maxresults), 10))
	}
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	if sharesnapshot != nil && len(*sharesnapshot) > 0 {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	params.Set("comp", "listhandles")
	req.URL.RawQuery = params.Encode()
	if recursive != nil {
		req.Header.Set("x-ms-recursive", strconv.FormatBool(*recursive))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	req.Header.Set("x-ms-file-request-intent", "backup")
	return req, nil
}

// listHandlesResponder handles the response to the ListHandles request.
func (client directoryClient) listHandlesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ListHandlesResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, err
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// SetMetadata updates user defined metadata for the specified directory.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
func (client directoryClient) SetMetadata(ctx context.Context, timeout *int32, metadata map[string]string) (*DirectorySetMetadataResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setMetadataPreparer(timeout, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setMetadataResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DirectorySetMetadataResponse), err
}

// setMetadataPreparer prepares the SetMetadata request.
func (client directoryClient) setMetadataPreparer(timeout *int32, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "directory")
	params.Set("comp", "metadata")
	req.URL.RawQuery = params.Encode()
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	req.Header.Set("x-ms-file-request-intent", "backup")

	return req, nil
}

// setMetadataResponder handles the response to the SetMetadata request.
func (client directoryClient) setMetadataResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &DirectorySetMetadataResponse{rawResponse: resp.Response()}, err
}

// SetProperties sets properties on the directory.
//
// fileAttributes is if specified, the provided file attributes shall be set. Default value: ‘Archive’ for file and
// ‘Directory’ for directory. ‘None’ can also be specified as default. fileCreationTime is creation time for the
// file/directory. Default value: Now. fileLastWriteTime is last write time for the file/directory. Default value: Now.
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> filePermission is if specified the permission (security descriptor) shall
// be set for the directory/file. This header can be used if Permission size is <= 8KB, else x-ms-file-permission-key
// header shall be used. Default value: Inherit. If SDDL is specified as input, it must have owner, group and dacl.
// Note: Only one of the x-ms-file-permission or x-ms-file-permission-key should be specified. filePermissionKey is key
// of the permission to be set for the directory/file. Note: Only one of the x-ms-file-permission or
// x-ms-file-permission-key should be specified.
func (client directoryClient) SetProperties(ctx context.Context, fileAttributes string, fileCreationTime string, fileLastWriteTime string, timeout *int32, filePermission *string, filePermissionKey *string) (*DirectorySetPropertiesResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setPropertiesPreparer(fileAttributes, fileCreationTime, fileLastWriteTime, timeout, filePermission, filePermissionKey)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DirectorySetPropertiesResponse), err
}

// setPropertiesPreparer prepares the SetProperties request.
func (client directoryClient) setPropertiesPreparer(fileAttributes string, fileCreationTime string, fileLastWriteTime string, timeout *int32, filePermission *string, filePermissionKey *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "directory")
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	req.Header.Set("x-ms-file-request-intent", "backup")

	if filePermission != nil {
		req.Header.Set("x-ms-file-permission", *filePermission)
	}
	if filePermissionKey != nil {
		req.Header.Set("x-ms-file-permission-key", *filePermissionKey)
	}
	req.Header.Set("x-ms-file-attributes", fileAttributes)
	req.Header.Set("x-ms-file-creation-time", fileCreationTime)
	req.Header.Set("x-ms-file-last-write-time", fileLastWriteTime)
	return req, nil
}

// setPropertiesResponder handles the response to the SetProperties request.
func (client directoryClient) setPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &DirectorySetPropertiesResponse{rawResponse: resp.Response()}, err
}
