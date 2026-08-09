/*
SPDX-License-Identifier: Apache-2.0
Copyright 2016 The Kubernetes Authors.
*/

// +deepequal-gen=package

// This is a test package.
package arrays

// MAC is a named array of a primitive type, e.g. a hardware address.
type MAC [6]byte

// Inner is a comparable struct used as an array element type.
type Inner struct {
	Int32  int32
	String string
}

type Ttest struct {
	Bytes   [4]byte
	Strings [2]string
	Named   MAC
	Structs [2]Inner
}
