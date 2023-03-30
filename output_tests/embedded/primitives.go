/*
SPDX-License-Identifier: Apache-2.0
Copyright 2016 The Kubernetes Authors.
*/

package embedded

type StructPrimitives struct {
	BoolField   bool
	IntField    int
	StringField string
	FloatField  float64
	OtherStructPrimitives
}

type OtherStructPrimitives struct {
	BoolField   bool
	IntField    int
	StringField string
	FloatField  float64
}
