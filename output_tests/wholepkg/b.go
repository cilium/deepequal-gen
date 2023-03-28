/*
SPDX-License-Identifier: Apache-2.0
Copyright 2016 The Kubernetes Authors.
*/

package wholepkg

import "github.com/cilium/deepequal-gen/output_tests/otherpkg"

// Another type in another file.
type StructB struct {
	OtherWithPrimitivesOnly otherpkg.OtherStructWithPrimitivesOnly
	OtherWithPointers       otherpkg.OtherStructWithPointers
}
