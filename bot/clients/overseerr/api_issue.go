/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr_go

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// IssueAPIService IssueAPI service
type IssueAPIService service

type IssueAPIIssueCommentCommentIdDeleteRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
	commentId string
}

func (r IssueAPIIssueCommentCommentIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.IssueCommentCommentIdDeleteExecute(r)
}

/*
IssueCommentCommentIdDelete Delete issue comment

Deletes an issue comment. Only users with `MANAGE_ISSUES` or the user who created the comment can perform this action.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commentId Issue Comment ID
 @return IssueAPIIssueCommentCommentIdDeleteRequest
*/
func (a *IssueAPIService) IssueCommentCommentIdDelete(ctx context.Context, commentId string) IssueAPIIssueCommentCommentIdDeleteRequest {
	return IssueAPIIssueCommentCommentIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		commentId: commentId,
	}
}

// Execute executes the request
func (a *IssueAPIService) IssueCommentCommentIdDeleteExecute(r IssueAPIIssueCommentCommentIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssueCommentCommentIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issueComment/{commentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"commentId"+"}", url.PathEscape(parameterValueToString(r.commentId, "commentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type IssueAPIIssueCommentCommentIdGetRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
	commentId string
}

func (r IssueAPIIssueCommentCommentIdGetRequest) Execute() (*IssueComment, *http.Response, error) {
	return r.ApiService.IssueCommentCommentIdGetExecute(r)
}

/*
IssueCommentCommentIdGet Get issue comment

Returns a single issue comment in JSON format.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commentId
 @return IssueAPIIssueCommentCommentIdGetRequest
*/
func (a *IssueAPIService) IssueCommentCommentIdGet(ctx context.Context, commentId string) IssueAPIIssueCommentCommentIdGetRequest {
	return IssueAPIIssueCommentCommentIdGetRequest{
		ApiService: a,
		ctx: ctx,
		commentId: commentId,
	}
}

// Execute executes the request
//  @return IssueComment
func (a *IssueAPIService) IssueCommentCommentIdGetExecute(r IssueAPIIssueCommentCommentIdGetRequest) (*IssueComment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IssueComment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssueCommentCommentIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issueComment/{commentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"commentId"+"}", url.PathEscape(parameterValueToString(r.commentId, "commentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type IssueAPIIssueCommentCommentIdPutRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
	commentId string
	issueCommentCommentIdPutRequest *IssueCommentCommentIdPutRequest
}

func (r IssueAPIIssueCommentCommentIdPutRequest) IssueCommentCommentIdPutRequest(issueCommentCommentIdPutRequest IssueCommentCommentIdPutRequest) IssueAPIIssueCommentCommentIdPutRequest {
	r.issueCommentCommentIdPutRequest = &issueCommentCommentIdPutRequest
	return r
}

func (r IssueAPIIssueCommentCommentIdPutRequest) Execute() (*IssueComment, *http.Response, error) {
	return r.ApiService.IssueCommentCommentIdPutExecute(r)
}

/*
IssueCommentCommentIdPut Update issue comment

Updates and returns a single issue comment in JSON format.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commentId
 @return IssueAPIIssueCommentCommentIdPutRequest
*/
func (a *IssueAPIService) IssueCommentCommentIdPut(ctx context.Context, commentId string) IssueAPIIssueCommentCommentIdPutRequest {
	return IssueAPIIssueCommentCommentIdPutRequest{
		ApiService: a,
		ctx: ctx,
		commentId: commentId,
	}
}

// Execute executes the request
//  @return IssueComment
func (a *IssueAPIService) IssueCommentCommentIdPutExecute(r IssueAPIIssueCommentCommentIdPutRequest) (*IssueComment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IssueComment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssueCommentCommentIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issueComment/{commentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"commentId"+"}", url.PathEscape(parameterValueToString(r.commentId, "commentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.issueCommentCommentIdPutRequest == nil {
		return localVarReturnValue, nil, reportError("issueCommentCommentIdPutRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.issueCommentCommentIdPutRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type IssueAPIIssueCountGetRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
}

func (r IssueAPIIssueCountGetRequest) Execute() (*IssueCountGet200Response, *http.Response, error) {
	return r.ApiService.IssueCountGetExecute(r)
}

/*
IssueCountGet Gets issue counts

Returns the number of open and closed issues, as well as the number of issues of each type.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return IssueAPIIssueCountGetRequest
*/
func (a *IssueAPIService) IssueCountGet(ctx context.Context) IssueAPIIssueCountGetRequest {
	return IssueAPIIssueCountGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IssueCountGet200Response
func (a *IssueAPIService) IssueCountGetExecute(r IssueAPIIssueCountGetRequest) (*IssueCountGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IssueCountGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssueCountGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issue/count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type IssueAPIIssueGetRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
	take *float32
	skip *float32
	sort *string
	filter *string
	requestedBy *float32
}

func (r IssueAPIIssueGetRequest) Take(take float32) IssueAPIIssueGetRequest {
	r.take = &take
	return r
}

func (r IssueAPIIssueGetRequest) Skip(skip float32) IssueAPIIssueGetRequest {
	r.skip = &skip
	return r
}

func (r IssueAPIIssueGetRequest) Sort(sort string) IssueAPIIssueGetRequest {
	r.sort = &sort
	return r
}

func (r IssueAPIIssueGetRequest) Filter(filter string) IssueAPIIssueGetRequest {
	r.filter = &filter
	return r
}

func (r IssueAPIIssueGetRequest) RequestedBy(requestedBy float32) IssueAPIIssueGetRequest {
	r.requestedBy = &requestedBy
	return r
}

func (r IssueAPIIssueGetRequest) Execute() (*IssueGet200Response, *http.Response, error) {
	return r.ApiService.IssueGetExecute(r)
}

/*
IssueGet Get all issues

Returns a list of issues in JSON format.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return IssueAPIIssueGetRequest
*/
func (a *IssueAPIService) IssueGet(ctx context.Context) IssueAPIIssueGetRequest {
	return IssueAPIIssueGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IssueGet200Response
func (a *IssueAPIService) IssueGetExecute(r IssueAPIIssueGetRequest) (*IssueGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IssueGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssueGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "take", r.take, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	} else {
		var defaultValue string = "added"
		r.sort = &defaultValue
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	} else {
		var defaultValue string = "open"
		r.filter = &defaultValue
	}
	if r.requestedBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "requestedBy", r.requestedBy, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type IssueAPIIssueIssueIdCommentPostRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
	issueId float32
	issueIssueIdCommentPostRequest *IssueIssueIdCommentPostRequest
}

func (r IssueAPIIssueIssueIdCommentPostRequest) IssueIssueIdCommentPostRequest(issueIssueIdCommentPostRequest IssueIssueIdCommentPostRequest) IssueAPIIssueIssueIdCommentPostRequest {
	r.issueIssueIdCommentPostRequest = &issueIssueIdCommentPostRequest
	return r
}

func (r IssueAPIIssueIssueIdCommentPostRequest) Execute() (*Issue, *http.Response, error) {
	return r.ApiService.IssueIssueIdCommentPostExecute(r)
}

/*
IssueIssueIdCommentPost Create a comment

Creates a comment and returns associated issue in JSON format.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId
 @return IssueAPIIssueIssueIdCommentPostRequest
*/
func (a *IssueAPIService) IssueIssueIdCommentPost(ctx context.Context, issueId float32) IssueAPIIssueIssueIdCommentPostRequest {
	return IssueAPIIssueIssueIdCommentPostRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
	}
}

// Execute executes the request
//  @return Issue
func (a *IssueAPIService) IssueIssueIdCommentPostExecute(r IssueAPIIssueIssueIdCommentPostRequest) (*Issue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Issue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssueIssueIdCommentPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issue/{issueId}/comment"
	localVarPath = strings.Replace(localVarPath, "{"+"issueId"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.issueIssueIdCommentPostRequest == nil {
		return localVarReturnValue, nil, reportError("issueIssueIdCommentPostRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.issueIssueIdCommentPostRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type IssueAPIIssueIssueIdDeleteRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
	issueId string
}

func (r IssueAPIIssueIssueIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.IssueIssueIdDeleteExecute(r)
}

/*
IssueIssueIdDelete Delete issue

Removes an issue. If the user has the `MANAGE_ISSUES` permission, any issue can be removed. Otherwise, only a users own issues can be removed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId Issue ID
 @return IssueAPIIssueIssueIdDeleteRequest
*/
func (a *IssueAPIService) IssueIssueIdDelete(ctx context.Context, issueId string) IssueAPIIssueIssueIdDeleteRequest {
	return IssueAPIIssueIssueIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
	}
}

// Execute executes the request
func (a *IssueAPIService) IssueIssueIdDeleteExecute(r IssueAPIIssueIssueIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssueIssueIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issue/{issueId}"
	localVarPath = strings.Replace(localVarPath, "{"+"issueId"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type IssueAPIIssueIssueIdGetRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
	issueId float32
}

func (r IssueAPIIssueIssueIdGetRequest) Execute() (*Issue, *http.Response, error) {
	return r.ApiService.IssueIssueIdGetExecute(r)
}

/*
IssueIssueIdGet Get issue

Returns a single issue in JSON format.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId
 @return IssueAPIIssueIssueIdGetRequest
*/
func (a *IssueAPIService) IssueIssueIdGet(ctx context.Context, issueId float32) IssueAPIIssueIssueIdGetRequest {
	return IssueAPIIssueIssueIdGetRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
	}
}

// Execute executes the request
//  @return Issue
func (a *IssueAPIService) IssueIssueIdGetExecute(r IssueAPIIssueIssueIdGetRequest) (*Issue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Issue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssueIssueIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issue/{issueId}"
	localVarPath = strings.Replace(localVarPath, "{"+"issueId"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type IssueAPIIssueIssueIdStatusPostRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
	issueId string
	status string
}

func (r IssueAPIIssueIssueIdStatusPostRequest) Execute() (*Issue, *http.Response, error) {
	return r.ApiService.IssueIssueIdStatusPostExecute(r)
}

/*
IssueIssueIdStatusPost Update an issue's status

Updates an issue's status to approved or declined. Also returns the issue in a JSON object.

Requires the `MANAGE_ISSUES` permission or `ADMIN`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId Issue ID
 @param status New status
 @return IssueAPIIssueIssueIdStatusPostRequest
*/
func (a *IssueAPIService) IssueIssueIdStatusPost(ctx context.Context, issueId string, status string) IssueAPIIssueIssueIdStatusPostRequest {
	return IssueAPIIssueIssueIdStatusPostRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
		status: status,
	}
}

// Execute executes the request
//  @return Issue
func (a *IssueAPIService) IssueIssueIdStatusPostExecute(r IssueAPIIssueIssueIdStatusPostRequest) (*Issue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Issue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssueIssueIdStatusPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issue/{issueId}/{status}"
	localVarPath = strings.Replace(localVarPath, "{"+"issueId"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"status"+"}", url.PathEscape(parameterValueToString(r.status, "status")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type IssueAPIIssuePostRequest struct {
	ctx context.Context
	ApiService *IssueAPIService
	issuePostRequest *IssuePostRequest
}

func (r IssueAPIIssuePostRequest) IssuePostRequest(issuePostRequest IssuePostRequest) IssueAPIIssuePostRequest {
	r.issuePostRequest = &issuePostRequest
	return r
}

func (r IssueAPIIssuePostRequest) Execute() (*Issue, *http.Response, error) {
	return r.ApiService.IssuePostExecute(r)
}

/*
IssuePost Create new issue

Creates a new issue


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return IssueAPIIssuePostRequest
*/
func (a *IssueAPIService) IssuePost(ctx context.Context) IssueAPIIssuePostRequest {
	return IssueAPIIssuePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Issue
func (a *IssueAPIService) IssuePostExecute(r IssueAPIIssuePostRequest) (*Issue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Issue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssueAPIService.IssuePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.issuePostRequest == nil {
		return localVarReturnValue, nil, reportError("issuePostRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.issuePostRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
