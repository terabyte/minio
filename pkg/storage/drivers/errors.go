/*
 * Minimalist Object Storage, (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package drivers

import "fmt"

// BackendError - generic disk backend error
type BackendError struct {
	Path string
}

// BackendCorrupted - path has corrupted data
type BackendCorrupted BackendError

// APINotImplemented - generic API not implemented error
type APINotImplemented struct {
	API string
}

// GenericBucketError - generic bucket error
type GenericBucketError struct {
	Bucket string
}

// GenericObjectError - generic object error
type GenericObjectError struct {
	Bucket string
	Object string
}

// ImplementationError - generic implementation error
type ImplementationError struct {
	Bucket string
	Object string
	Err    error
}

// DigestError - Generic Md5 error
type DigestError struct {
	Bucket string
	Key    string
	Md5    string
}

/// ACL related errors

// InvalidACL - acl invalid
type InvalidACL struct {
	ACL string
}

func (e InvalidACL) Error() string {
	return "Requested ACL is " + e.ACL + " invalid"
}

/// Bucket related errors

// BucketNameInvalid - bucketname provided is invalid
type BucketNameInvalid GenericBucketError

// BucketExists - bucket already exists
type BucketExists GenericBucketError

// BucketNotFound - requested bucket not found
type BucketNotFound GenericBucketError

// TooManyBuckets - total buckets exceeded
type TooManyBuckets GenericBucketError

/// Object related errors

// ObjectNotFound - requested object not found
type ObjectNotFound GenericObjectError

// ObjectExists - object already exists
type ObjectExists GenericObjectError

// EntityTooLarge - object size exceeds maximum limit
type EntityTooLarge struct {
	GenericObjectError
	Size      string
	TotalSize string
}

// ObjectNameInvalid - object name provided is invalid
type ObjectNameInvalid GenericObjectError

// BadDigest - md5 mismatch from data received
type BadDigest DigestError

// InvalidDigest - md5 in request header invalid
type InvalidDigest DigestError

// Return string an error formatted as the given text
func (e ImplementationError) Error() string {
	error := ""
	if e.Bucket != "" {
		error = error + "Bucket: " + e.Bucket + " "
	}
	if e.Object != "" {
		error = error + "Object: " + e.Object + " "
	}
	error = error + "Error: " + e.Err.Error()
	return error
}

// EmbedError - wrapper function for error object
func EmbedError(bucket, object string, err error) ImplementationError {
	return ImplementationError{
		Bucket: bucket,
		Object: object,
		Err:    err,
	}
}

// Return string an error formatted as the given text
func (e ObjectNotFound) Error() string {
	return "Object not Found: " + e.Bucket + "#" + e.Object
}

// Return string an error formatted as the given text
func (e APINotImplemented) Error() string {
	return "Api not implemented: " + e.API
}

// Return string an error formatted as the given text
func (e ObjectExists) Error() string {
	return "Object exists: " + e.Bucket + "#" + e.Object
}

// Return string an error formatted as the given text
func (e BucketNameInvalid) Error() string {
	return "Bucket name invalid: " + e.Bucket
}

// Return string an error formatted as the given text
func (e BucketExists) Error() string {
	return "Bucket exists: " + e.Bucket
}

// Return string an error formatted as the given text
func (e TooManyBuckets) Error() string {
	return "Bucket limit exceeded beyond 100, cannot create bucket: " + e.Bucket
}

// Return string an error formatted as the given text
func (e BucketNotFound) Error() string {
	return "Bucket not Found: " + e.Bucket
}

// Return string an error formatted as the given text
func (e ObjectNameInvalid) Error() string {
	return "Object name invalid: " + e.Bucket + "#" + e.Object
}

// Return string an error formatted as the given text
func (e EntityTooLarge) Error() string {
	return e.Bucket + "#" + e.Object + "with " + e.Size + "reached maximum allowed size limit " + e.TotalSize
}

// Return string an error formatted as the given text
func (e BackendCorrupted) Error() string {
	return "Backend corrupted: " + e.Path
}

// Return string an error formatted as the given text
func (e BadDigest) Error() string {
	return "Md5 provided " + e.Md5 + " mismatches for: " + e.Bucket + "#" + e.Key
}

// Return string an error formatted as the given text
func (e InvalidDigest) Error() string {
	return "Md5 provided " + e.Md5 + " is invalid"
}

// OperationNotPermitted - operation not permitted
type OperationNotPermitted struct {
	Op     string
	Reason string
}

func (e OperationNotPermitted) Error() string {
	return "Operation " + e.Op + " not permitted for reason: " + e.Reason
}

// InvalidRange - invalid range
type InvalidRange struct {
	Start  int64
	Length int64
}

func (e InvalidRange) Error() string {
	return fmt.Sprintf("Invalid range start:%d length:%d", e.Start, e.Length)
}
