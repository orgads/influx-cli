/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	_context "context"
	_fmt "fmt"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"time"
)

// Linger please
var (
	_ _context.Context
)

type BackupApi interface {

	/*
			 * GetBackupKV Download snapshot of metadata stored in the server's embedded KV store. Don't use with InfluxDB versions greater than InfluxDB 2.1.x.
			 * Retrieves a snapshot of metadata stored in the server's embedded KV store.
		InfluxDB versions greater than 2.1.x don't include metadata stored in embedded SQL;
		avoid using this endpoint with versions greater than 2.1.x.

			 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 * @return ApiGetBackupKVRequest
	*/
	GetBackupKV(ctx _context.Context) ApiGetBackupKVRequest

	/*
	 * GetBackupKVExecute executes the request
	 * @return *os.File
	 */
	GetBackupKVExecute(r ApiGetBackupKVRequest) (*_nethttp.Response, error)

	/*
	 * GetBackupKVExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return *os.File
	 */
	GetBackupKVExecuteWithHttpInfo(r ApiGetBackupKVRequest) (*_nethttp.Response, *_nethttp.Response, error)

	/*
	 * GetBackupMetadata Download snapshot of all metadata in the server
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetBackupMetadataRequest
	 */
	GetBackupMetadata(ctx _context.Context) ApiGetBackupMetadataRequest

	/*
	 * GetBackupMetadataExecute executes the request
	 * @return *os.File
	 */
	GetBackupMetadataExecute(r ApiGetBackupMetadataRequest) (*_nethttp.Response, error)

	/*
	 * GetBackupMetadataExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return *os.File
	 */
	GetBackupMetadataExecuteWithHttpInfo(r ApiGetBackupMetadataRequest) (*_nethttp.Response, *_nethttp.Response, error)

	/*
	 * GetBackupShardId Download snapshot of all TSM data in a shard
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param shardID The shard ID.
	 * @return ApiGetBackupShardIdRequest
	 */
	GetBackupShardId(ctx _context.Context, shardID int64) ApiGetBackupShardIdRequest

	/*
	 * GetBackupShardIdExecute executes the request
	 * @return *os.File
	 */
	GetBackupShardIdExecute(r ApiGetBackupShardIdRequest) (*_nethttp.Response, error)

	/*
	 * GetBackupShardIdExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return *os.File
	 */
	GetBackupShardIdExecuteWithHttpInfo(r ApiGetBackupShardIdRequest) (*_nethttp.Response, *_nethttp.Response, error)
}

// BackupApiService BackupApi service
type BackupApiService service

type ApiGetBackupKVRequest struct {
	ctx          _context.Context
	ApiService   BackupApi
	zapTraceSpan *string
}

func (r ApiGetBackupKVRequest) ZapTraceSpan(zapTraceSpan string) ApiGetBackupKVRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetBackupKVRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetBackupKVRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetBackupKVExecute(r)
}

func (r ApiGetBackupKVRequest) ExecuteWithHttpInfo() (*_nethttp.Response, *_nethttp.Response, error) {
	return r.ApiService.GetBackupKVExecuteWithHttpInfo(r)
}

/*
 * GetBackupKV Download snapshot of metadata stored in the server's embedded KV store. Don't use with InfluxDB versions greater than InfluxDB 2.1.x.
 * Retrieves a snapshot of metadata stored in the server's embedded KV store.
InfluxDB versions greater than 2.1.x don't include metadata stored in embedded SQL;
avoid using this endpoint with versions greater than 2.1.x.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetBackupKVRequest
*/
func (a *BackupApiService) GetBackupKV(ctx _context.Context) ApiGetBackupKVRequest {
	return ApiGetBackupKVRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *BackupApiService) GetBackupKVExecute(r ApiGetBackupKVRequest) (*_nethttp.Response, error) {
	returnVal, _, err := a.GetBackupKVExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return *os.File
 */
func (a *BackupApiService) GetBackupKVExecuteWithHttpInfo(r ApiGetBackupKVRequest) (*_nethttp.Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *_nethttp.Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupApiService.GetBackupKV")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/backup/kv"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	localVarReturnValue = localVarHTTPResponse

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetBackupMetadataRequest struct {
	ctx            _context.Context
	ApiService     BackupApi
	zapTraceSpan   *string
	acceptEncoding *string
}

func (r ApiGetBackupMetadataRequest) ZapTraceSpan(zapTraceSpan string) ApiGetBackupMetadataRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetBackupMetadataRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetBackupMetadataRequest) AcceptEncoding(acceptEncoding string) ApiGetBackupMetadataRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}
func (r ApiGetBackupMetadataRequest) GetAcceptEncoding() *string {
	return r.acceptEncoding
}

func (r ApiGetBackupMetadataRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetBackupMetadataExecute(r)
}

func (r ApiGetBackupMetadataRequest) ExecuteWithHttpInfo() (*_nethttp.Response, *_nethttp.Response, error) {
	return r.ApiService.GetBackupMetadataExecuteWithHttpInfo(r)
}

/*
 * GetBackupMetadata Download snapshot of all metadata in the server
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetBackupMetadataRequest
 */
func (a *BackupApiService) GetBackupMetadata(ctx _context.Context) ApiGetBackupMetadataRequest {
	return ApiGetBackupMetadataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *BackupApiService) GetBackupMetadataExecute(r ApiGetBackupMetadataRequest) (*_nethttp.Response, error) {
	returnVal, _, err := a.GetBackupMetadataExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return *os.File
 */
func (a *BackupApiService) GetBackupMetadataExecuteWithHttpInfo(r ApiGetBackupMetadataRequest) (*_nethttp.Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *_nethttp.Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupApiService.GetBackupMetadata")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/backup/metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"multipart/mixed", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	if r.acceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = parameterToString(*r.acceptEncoding, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	localVarReturnValue = localVarHTTPResponse

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetBackupShardIdRequest struct {
	ctx            _context.Context
	ApiService     BackupApi
	shardID        int64
	zapTraceSpan   *string
	acceptEncoding *string
	since          *time.Time
}

func (r ApiGetBackupShardIdRequest) ShardID(shardID int64) ApiGetBackupShardIdRequest {
	r.shardID = shardID
	return r
}
func (r ApiGetBackupShardIdRequest) GetShardID() int64 {
	return r.shardID
}

func (r ApiGetBackupShardIdRequest) ZapTraceSpan(zapTraceSpan string) ApiGetBackupShardIdRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetBackupShardIdRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetBackupShardIdRequest) AcceptEncoding(acceptEncoding string) ApiGetBackupShardIdRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}
func (r ApiGetBackupShardIdRequest) GetAcceptEncoding() *string {
	return r.acceptEncoding
}

func (r ApiGetBackupShardIdRequest) Since(since time.Time) ApiGetBackupShardIdRequest {
	r.since = &since
	return r
}
func (r ApiGetBackupShardIdRequest) GetSince() *time.Time {
	return r.since
}

func (r ApiGetBackupShardIdRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetBackupShardIdExecute(r)
}

func (r ApiGetBackupShardIdRequest) ExecuteWithHttpInfo() (*_nethttp.Response, *_nethttp.Response, error) {
	return r.ApiService.GetBackupShardIdExecuteWithHttpInfo(r)
}

/*
 * GetBackupShardId Download snapshot of all TSM data in a shard
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param shardID The shard ID.
 * @return ApiGetBackupShardIdRequest
 */
func (a *BackupApiService) GetBackupShardId(ctx _context.Context, shardID int64) ApiGetBackupShardIdRequest {
	return ApiGetBackupShardIdRequest{
		ApiService: a,
		ctx:        ctx,
		shardID:    shardID,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *BackupApiService) GetBackupShardIdExecute(r ApiGetBackupShardIdRequest) (*_nethttp.Response, error) {
	returnVal, _, err := a.GetBackupShardIdExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return *os.File
 */
func (a *BackupApiService) GetBackupShardIdExecuteWithHttpInfo(r ApiGetBackupShardIdRequest) (*_nethttp.Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *_nethttp.Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupApiService.GetBackupShardId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/backup/shards/{shardID}"
	localVarPath = strings.Replace(localVarPath, "{"+"shardID"+"}", _neturl.PathEscape(parameterToString(r.shardID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.since != nil {
		localVarQueryParams.Add("since", parameterToString(*r.since, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	if r.acceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = parameterToString(*r.acceptEncoding, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
			newErr.model = &v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	localVarReturnValue = localVarHTTPResponse

	return localVarReturnValue, localVarHTTPResponse, nil
}
