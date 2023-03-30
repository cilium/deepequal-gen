/*
SPDX-License-Identifier: Apache-2.0
Copyright 2016 The Kubernetes Authors.
*/

package embedded

type StructWithEmbeddedPointerStructs struct {
	OtherStructPointer
}

type OtherStructPointer struct {
	BoolField   *bool
	IntField    *int
	StringField *string
	FloatField  *float64
	Self        *OtherStructPointer
}
