// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"net/http"
	"reflect"
	"strings"
)

type result struct {
	resp *http.Response
	err  error
}

type nonRetriableError struct {
	error
}

func (nonRetriableError) NonRetriable() {
	// marker method
}

func contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func getOptional[T any](v T) *T {
	if reflect.ValueOf(v).IsZero() {
		return nil
	}
	return &v
}

func parseOptional[T any](v string, parse func(v string) (T, error)) (*T, error) {
	if v == "" {
		return nil, nil
	}
	t, err := parse(v)
	if err != nil {
		return nil, err
	}
	return &t, err
}

func splitHelper(s, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}
