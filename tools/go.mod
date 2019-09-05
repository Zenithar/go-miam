module go.zenithar.org/miam/tools

go 1.13

replace github.com/envoyproxy/protoc-gen-validate => github.com/Zenithar/protoc-gen-validate v0.2.0-java.0.20190808132234-82cadd0ebcfe

replace github.com/uber/prototool => github.com/Zenithar/prototool v1.8.2-0.20190905183603-d4acf041d295

require (
	github.com/envoyproxy/protoc-gen-validate v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.7.0
	github.com/frapposelli/wwhrd v0.2.4
	github.com/gogo/protobuf v1.3.0
	github.com/golang/mock v1.3.1
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365 // indirect
	github.com/izumin5210/gex v0.5.1 // indirect
	github.com/magefile/mage v1.8.0
	github.com/uber/prototool v0.0.0-00010101000000-000000000000
	go.zenithar.org/protoc-gen-cobra v0.0.3
	gotest.tools/gotestsum v0.3.5
	mvdan.cc/gofumpt v0.0.0-20190729090447-96300e3d49fb
)
