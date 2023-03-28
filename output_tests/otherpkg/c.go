/*
SPDX-License-Identifier: Apache-2.0
Copyright 2016 The Kubernetes Authors.
*/

package otherpkg

type OtherStructWithPrimitivesOnly struct {
	BoolField bool
}

type OtherStructWithPointers struct {
	BoolField *bool
}
