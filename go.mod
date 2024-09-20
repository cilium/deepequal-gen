module github.com/cilium/deepequal-gen

go 1.23.0

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/google/gofuzz v1.2.0
	github.com/spf13/pflag v1.0.5
	k8s.io/gengo v0.0.0-20240911193312-2b36238f13e9
	k8s.io/klog/v2 v2.130.1
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/tools v0.25.0 // indirect
)

// This is to support type alias in go 1.23
// More details can be found in https://github.com/sayboras/gengo/pull/1
replace k8s.io/gengo => github.com/sayboras/gengo v0.0.0-20240920133050-8cd674ebf6a9
