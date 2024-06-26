// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"time"
)

type ACLFailedEntry struct {
	ErrorMessage *string
	Name         *string
	Type         *string
}

type PathHierarchyListSegment struct {
	// REQUIRED
	PathItems    []*PathItemInternal `xml:"Blob"`
	PathPrefixes []*PathPrefix       `xml:"PathPrefix"`
}

// PathItemInternal - An Azure Storage blob
type PathItemInternal struct {
	// REQUIRED
	Deleted *bool `xml:"Deleted"`

	// REQUIRED
	Name *string `xml:"Name"`

	// REQUIRED; Properties of a blob
	Properties *PathPropertiesInternal `xml:"Properties"`

	// REQUIRED
	Snapshot         *string `xml:"Snapshot"`
	DeletionID       *string `xml:"DeletionId"`
	IsCurrentVersion *bool   `xml:"IsCurrentVersion"`
	VersionID        *string `xml:"VersionId"`
}

type PathPrefix struct {
	// REQUIRED
	Name *string `xml:"Name"`
}

// PathPropertiesInternal - Properties of a blob
type PathPropertiesInternal struct {
	// REQUIRED
	ETag *azcore.ETag `xml:"Etag"`

	// REQUIRED
	LastModified         *time.Time `xml:"Last-Modified"`
	AccessTierChangeTime *time.Time `xml:"AccessTierChangeTime"`
	AccessTierInferred   *bool      `xml:"AccessTierInferred"`
	BlobSequenceNumber   *int64     `xml:"x-ms-blob-sequence-number"`
	CacheControl         *string    `xml:"Cache-Control"`
	ContentDisposition   *string    `xml:"Content-Disposition"`
	ContentEncoding      *string    `xml:"Content-Encoding"`
	ContentLanguage      *string    `xml:"Content-Language"`

	// Size in bytes
	ContentLength             *int64     `xml:"Content-Length"`
	ContentMD5                []byte     `xml:"Content-MD5"`
	ContentType               *string    `xml:"Content-Type"`
	CopyCompletionTime        *time.Time `xml:"CopyCompletionTime"`
	CopyID                    *string    `xml:"CopyId"`
	CopyProgress              *string    `xml:"CopyProgress"`
	CopySource                *string    `xml:"CopySource"`
	CopyStatusDescription     *string    `xml:"CopyStatusDescription"`
	CreationTime              *time.Time `xml:"Creation-Time"`
	CustomerProvidedKeySHA256 *string    `xml:"CustomerProvidedKeySha256"`
	DeleteTime                *time.Time `xml:"DeleteTime"`
	DeletedTime               *time.Time `xml:"DeletedTime"`
	DestinationSnapshot       *string    `xml:"DestinationSnapshot"`

	// The name of the encryption scope under which the blob is encrypted.
	EncryptionScope        *string    `xml:"EncryptionScope"`
	ExpiresOn              *time.Time `xml:"Expiry-Time"`
	IncrementalCopy        *bool      `xml:"IncrementalCopy"`
	IsSealed               *bool      `xml:"Sealed"`
	LastAccessedOn         *time.Time `xml:"LastAccessTime"`
	RemainingRetentionDays *int32     `xml:"RemainingRetentionDays"`
	ServerEncrypted        *bool      `xml:"ServerEncrypted"`
	TagCount               *int32     `xml:"TagCount"`
}

type FileSystem struct {
	ETag         *string
	LastModified *string
	Name         *string
}

type FileSystemList struct {
	Filesystems []*FileSystem
}

// ListPathsHierarchySegmentResponse - An enumeration of blobs
type ListPathsHierarchySegmentResponse struct {
	// REQUIRED
	FileSystemName *string `xml:"ContainerName,attr"`

	// REQUIRED
	Segment *PathHierarchyListSegment `xml:"Blobs"`

	// REQUIRED
	ServiceEndpoint *string `xml:"ServiceEndpoint,attr"`
	Delimiter       *string `xml:"Delimiter"`
	Marker          *string `xml:"Marker"`
	MaxResults      *int32  `xml:"MaxResults"`
	NextMarker      *string `xml:"NextMarker"`
	Prefix          *string `xml:"Prefix"`
}

type Path struct {
	ContentLength     *int64
	CreationTime      *string
	ETag              *string
	EncryptionContext *string

	// The name of the encryption scope under which the blob is encrypted.
	EncryptionScope *string
	ExpiryTime      *string
	Group           *string
	IsDirectory     *bool
	LastModified    *string
	Name            *string
	Owner           *string
	Permissions     *string
}

type PathList struct {
	Paths []*Path
}

type SetAccessControlRecursiveResponse struct {
	DirectoriesSuccessful *int32
	FailedEntries         []*ACLFailedEntry
	FailureCount          *int32
	FilesSuccessful       *int32
}

type StorageError struct {
	// The service error response object.
	Error *StorageErrorError
}

// StorageErrorError - The service error response object.
type StorageErrorError struct {
	// The service error code.
	Code *string

	// The service error message.
	Message *string
}
