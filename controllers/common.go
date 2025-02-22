/*
Copyright 2023 The Nephio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

const (
	ConfigMapVersionAnnotation = "workload.nephio.org/configMapVersion"

	// TODO(jbelamaric): Update to use ImageConfig spec.ImagePaths["upf"]
	AMFImage = "docker.io/nephio/free5gc-amf:3.2.1"
	SMFImage = "docker.io/nephio/free5gc-smf:3.2.1"
	// UPFImage = "docker.io/nephio/free5gc-upf:3.2.1"
	UPFImage = "prod.harbor.keysight.digital/loadcore/upf:nephio-r1"
)
