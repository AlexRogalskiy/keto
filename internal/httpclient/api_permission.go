/*
 * Ory Keto API
 *
 * Documentation for all of Ory Keto's REST APIs. gRPC is documented separately.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type PermissionApi interface {

	/*
	 * CheckPermission Check a permission
	 * To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return PermissionApiApiCheckPermissionRequest
	 */
	CheckPermission(ctx context.Context) PermissionApiApiCheckPermissionRequest

	/*
	 * CheckPermissionExecute executes the request
	 * @return CheckPermissionResult
	 */
	CheckPermissionExecute(r PermissionApiApiCheckPermissionRequest) (*CheckPermissionResult, *http.Response, error)

	/*
	 * CheckPermissionOrError Check a permission
	 * To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return PermissionApiApiCheckPermissionOrErrorRequest
	 */
	CheckPermissionOrError(ctx context.Context) PermissionApiApiCheckPermissionOrErrorRequest

	/*
	 * CheckPermissionOrErrorExecute executes the request
	 * @return CheckPermissionResult
	 */
	CheckPermissionOrErrorExecute(r PermissionApiApiCheckPermissionOrErrorRequest) (*CheckPermissionResult, *http.Response, error)

	/*
	 * ExpandPermissions Expand a Relationship into permissions.
	 * Use this endpoint to expand a relationship tuple into permissions.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return PermissionApiApiExpandPermissionsRequest
	 */
	ExpandPermissions(ctx context.Context) PermissionApiApiExpandPermissionsRequest

	/*
	 * ExpandPermissionsExecute executes the request
	 * @return ExpandedPermissionTree
	 */
	ExpandPermissionsExecute(r PermissionApiApiExpandPermissionsRequest) (*ExpandedPermissionTree, *http.Response, error)

	/*
	 * PostCheckPermission Check a permission
	 * To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return PermissionApiApiPostCheckPermissionRequest
	 */
	PostCheckPermission(ctx context.Context) PermissionApiApiPostCheckPermissionRequest

	/*
	 * PostCheckPermissionExecute executes the request
	 * @return CheckPermissionResult
	 */
	PostCheckPermissionExecute(r PermissionApiApiPostCheckPermissionRequest) (*CheckPermissionResult, *http.Response, error)

	/*
	 * PostCheckPermissionOrError Check a permission
	 * To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return PermissionApiApiPostCheckPermissionOrErrorRequest
	 */
	PostCheckPermissionOrError(ctx context.Context) PermissionApiApiPostCheckPermissionOrErrorRequest

	/*
	 * PostCheckPermissionOrErrorExecute executes the request
	 * @return CheckPermissionResult
	 */
	PostCheckPermissionOrErrorExecute(r PermissionApiApiPostCheckPermissionOrErrorRequest) (*CheckPermissionResult, *http.Response, error)
}

// PermissionApiService PermissionApi service
type PermissionApiService service

type PermissionApiApiCheckPermissionRequest struct {
	ctx                 context.Context
	ApiService          PermissionApi
	namespace           *string
	object              *string
	relation            *string
	subjectId           *string
	subjectSetNamespace *string
	subjectSetObject    *string
	subjectSetRelation  *string
	maxDepth            *int64
}

func (r PermissionApiApiCheckPermissionRequest) Namespace(namespace string) PermissionApiApiCheckPermissionRequest {
	r.namespace = &namespace
	return r
}
func (r PermissionApiApiCheckPermissionRequest) Object(object string) PermissionApiApiCheckPermissionRequest {
	r.object = &object
	return r
}
func (r PermissionApiApiCheckPermissionRequest) Relation(relation string) PermissionApiApiCheckPermissionRequest {
	r.relation = &relation
	return r
}
func (r PermissionApiApiCheckPermissionRequest) SubjectId(subjectId string) PermissionApiApiCheckPermissionRequest {
	r.subjectId = &subjectId
	return r
}
func (r PermissionApiApiCheckPermissionRequest) SubjectSetNamespace(subjectSetNamespace string) PermissionApiApiCheckPermissionRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}
func (r PermissionApiApiCheckPermissionRequest) SubjectSetObject(subjectSetObject string) PermissionApiApiCheckPermissionRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}
func (r PermissionApiApiCheckPermissionRequest) SubjectSetRelation(subjectSetRelation string) PermissionApiApiCheckPermissionRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}
func (r PermissionApiApiCheckPermissionRequest) MaxDepth(maxDepth int64) PermissionApiApiCheckPermissionRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r PermissionApiApiCheckPermissionRequest) Execute() (*CheckPermissionResult, *http.Response, error) {
	return r.ApiService.CheckPermissionExecute(r)
}

/*
 * CheckPermission Check a permission
 * To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return PermissionApiApiCheckPermissionRequest
 */
func (a *PermissionApiService) CheckPermission(ctx context.Context) PermissionApiApiCheckPermissionRequest {
	return PermissionApiApiCheckPermissionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return CheckPermissionResult
 */
func (a *PermissionApiService) CheckPermissionExecute(r PermissionApiApiCheckPermissionRequest) (*CheckPermissionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *CheckPermissionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.CheckPermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check/openapi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.object != nil {
		localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	}
	if r.relation != nil {
		localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	}
	if r.subjectId != nil {
		localVarQueryParams.Add("subject_id", parameterToString(*r.subjectId, ""))
	}
	if r.subjectSetNamespace != nil {
		localVarQueryParams.Add("subject_set.namespace", parameterToString(*r.subjectSetNamespace, ""))
	}
	if r.subjectSetObject != nil {
		localVarQueryParams.Add("subject_set.object", parameterToString(*r.subjectSetObject, ""))
	}
	if r.subjectSetRelation != nil {
		localVarQueryParams.Add("subject_set.relation", parameterToString(*r.subjectSetRelation, ""))
	}
	if r.maxDepth != nil {
		localVarQueryParams.Add("max-depth", parameterToString(*r.maxDepth, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type PermissionApiApiCheckPermissionOrErrorRequest struct {
	ctx                 context.Context
	ApiService          PermissionApi
	namespace           *string
	object              *string
	relation            *string
	subjectId           *string
	subjectSetNamespace *string
	subjectSetObject    *string
	subjectSetRelation  *string
	maxDepth            *int64
}

func (r PermissionApiApiCheckPermissionOrErrorRequest) Namespace(namespace string) PermissionApiApiCheckPermissionOrErrorRequest {
	r.namespace = &namespace
	return r
}
func (r PermissionApiApiCheckPermissionOrErrorRequest) Object(object string) PermissionApiApiCheckPermissionOrErrorRequest {
	r.object = &object
	return r
}
func (r PermissionApiApiCheckPermissionOrErrorRequest) Relation(relation string) PermissionApiApiCheckPermissionOrErrorRequest {
	r.relation = &relation
	return r
}
func (r PermissionApiApiCheckPermissionOrErrorRequest) SubjectId(subjectId string) PermissionApiApiCheckPermissionOrErrorRequest {
	r.subjectId = &subjectId
	return r
}
func (r PermissionApiApiCheckPermissionOrErrorRequest) SubjectSetNamespace(subjectSetNamespace string) PermissionApiApiCheckPermissionOrErrorRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}
func (r PermissionApiApiCheckPermissionOrErrorRequest) SubjectSetObject(subjectSetObject string) PermissionApiApiCheckPermissionOrErrorRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}
func (r PermissionApiApiCheckPermissionOrErrorRequest) SubjectSetRelation(subjectSetRelation string) PermissionApiApiCheckPermissionOrErrorRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}
func (r PermissionApiApiCheckPermissionOrErrorRequest) MaxDepth(maxDepth int64) PermissionApiApiCheckPermissionOrErrorRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r PermissionApiApiCheckPermissionOrErrorRequest) Execute() (*CheckPermissionResult, *http.Response, error) {
	return r.ApiService.CheckPermissionOrErrorExecute(r)
}

/*
 * CheckPermissionOrError Check a permission
 * To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return PermissionApiApiCheckPermissionOrErrorRequest
 */
func (a *PermissionApiService) CheckPermissionOrError(ctx context.Context) PermissionApiApiCheckPermissionOrErrorRequest {
	return PermissionApiApiCheckPermissionOrErrorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return CheckPermissionResult
 */
func (a *PermissionApiService) CheckPermissionOrErrorExecute(r PermissionApiApiCheckPermissionOrErrorRequest) (*CheckPermissionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *CheckPermissionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.CheckPermissionOrError")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.object != nil {
		localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	}
	if r.relation != nil {
		localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	}
	if r.subjectId != nil {
		localVarQueryParams.Add("subject_id", parameterToString(*r.subjectId, ""))
	}
	if r.subjectSetNamespace != nil {
		localVarQueryParams.Add("subject_set.namespace", parameterToString(*r.subjectSetNamespace, ""))
	}
	if r.subjectSetObject != nil {
		localVarQueryParams.Add("subject_set.object", parameterToString(*r.subjectSetObject, ""))
	}
	if r.subjectSetRelation != nil {
		localVarQueryParams.Add("subject_set.relation", parameterToString(*r.subjectSetRelation, ""))
	}
	if r.maxDepth != nil {
		localVarQueryParams.Add("max-depth", parameterToString(*r.maxDepth, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v CheckPermissionResult
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type PermissionApiApiExpandPermissionsRequest struct {
	ctx        context.Context
	ApiService PermissionApi
	namespace  *string
	object     *string
	relation   *string
	maxDepth   *int64
}

func (r PermissionApiApiExpandPermissionsRequest) Namespace(namespace string) PermissionApiApiExpandPermissionsRequest {
	r.namespace = &namespace
	return r
}
func (r PermissionApiApiExpandPermissionsRequest) Object(object string) PermissionApiApiExpandPermissionsRequest {
	r.object = &object
	return r
}
func (r PermissionApiApiExpandPermissionsRequest) Relation(relation string) PermissionApiApiExpandPermissionsRequest {
	r.relation = &relation
	return r
}
func (r PermissionApiApiExpandPermissionsRequest) MaxDepth(maxDepth int64) PermissionApiApiExpandPermissionsRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r PermissionApiApiExpandPermissionsRequest) Execute() (*ExpandedPermissionTree, *http.Response, error) {
	return r.ApiService.ExpandPermissionsExecute(r)
}

/*
 * ExpandPermissions Expand a Relationship into permissions.
 * Use this endpoint to expand a relationship tuple into permissions.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return PermissionApiApiExpandPermissionsRequest
 */
func (a *PermissionApiService) ExpandPermissions(ctx context.Context) PermissionApiApiExpandPermissionsRequest {
	return PermissionApiApiExpandPermissionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return ExpandedPermissionTree
 */
func (a *PermissionApiService) ExpandPermissionsExecute(r PermissionApiApiExpandPermissionsRequest) (*ExpandedPermissionTree, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *ExpandedPermissionTree
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.ExpandPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/expand"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.namespace == nil {
		return localVarReturnValue, nil, reportError("namespace is required and must be specified")
	}
	if r.object == nil {
		return localVarReturnValue, nil, reportError("object is required and must be specified")
	}
	if r.relation == nil {
		return localVarReturnValue, nil, reportError("relation is required and must be specified")
	}

	localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	if r.maxDepth != nil {
		localVarQueryParams.Add("max-depth", parameterToString(*r.maxDepth, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type PermissionApiApiPostCheckPermissionRequest struct {
	ctx                     context.Context
	ApiService              PermissionApi
	maxDepth                *int64
	postCheckPermissionBody *PostCheckPermissionBody
}

func (r PermissionApiApiPostCheckPermissionRequest) MaxDepth(maxDepth int64) PermissionApiApiPostCheckPermissionRequest {
	r.maxDepth = &maxDepth
	return r
}
func (r PermissionApiApiPostCheckPermissionRequest) PostCheckPermissionBody(postCheckPermissionBody PostCheckPermissionBody) PermissionApiApiPostCheckPermissionRequest {
	r.postCheckPermissionBody = &postCheckPermissionBody
	return r
}

func (r PermissionApiApiPostCheckPermissionRequest) Execute() (*CheckPermissionResult, *http.Response, error) {
	return r.ApiService.PostCheckPermissionExecute(r)
}

/*
 * PostCheckPermission Check a permission
 * To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return PermissionApiApiPostCheckPermissionRequest
 */
func (a *PermissionApiService) PostCheckPermission(ctx context.Context) PermissionApiApiPostCheckPermissionRequest {
	return PermissionApiApiPostCheckPermissionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return CheckPermissionResult
 */
func (a *PermissionApiService) PostCheckPermissionExecute(r PermissionApiApiPostCheckPermissionRequest) (*CheckPermissionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *CheckPermissionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.PostCheckPermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check/openapi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxDepth != nil {
		localVarQueryParams.Add("max-depth", parameterToString(*r.maxDepth, ""))
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
	localVarPostBody = r.postCheckPermissionBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type PermissionApiApiPostCheckPermissionOrErrorRequest struct {
	ctx                            context.Context
	ApiService                     PermissionApi
	maxDepth                       *int64
	postCheckPermissionOrErrorBody *PostCheckPermissionOrErrorBody
}

func (r PermissionApiApiPostCheckPermissionOrErrorRequest) MaxDepth(maxDepth int64) PermissionApiApiPostCheckPermissionOrErrorRequest {
	r.maxDepth = &maxDepth
	return r
}
func (r PermissionApiApiPostCheckPermissionOrErrorRequest) PostCheckPermissionOrErrorBody(postCheckPermissionOrErrorBody PostCheckPermissionOrErrorBody) PermissionApiApiPostCheckPermissionOrErrorRequest {
	r.postCheckPermissionOrErrorBody = &postCheckPermissionOrErrorBody
	return r
}

func (r PermissionApiApiPostCheckPermissionOrErrorRequest) Execute() (*CheckPermissionResult, *http.Response, error) {
	return r.ApiService.PostCheckPermissionOrErrorExecute(r)
}

/*
 * PostCheckPermissionOrError Check a permission
 * To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return PermissionApiApiPostCheckPermissionOrErrorRequest
 */
func (a *PermissionApiService) PostCheckPermissionOrError(ctx context.Context) PermissionApiApiPostCheckPermissionOrErrorRequest {
	return PermissionApiApiPostCheckPermissionOrErrorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return CheckPermissionResult
 */
func (a *PermissionApiService) PostCheckPermissionOrErrorExecute(r PermissionApiApiPostCheckPermissionOrErrorRequest) (*CheckPermissionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *CheckPermissionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.PostCheckPermissionOrError")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxDepth != nil {
		localVarQueryParams.Add("max-depth", parameterToString(*r.maxDepth, ""))
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
	localVarPostBody = r.postCheckPermissionOrErrorBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v CheckPermissionResult
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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
