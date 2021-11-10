// Copyright (c) 2020-2021 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s

//go:generate go get k8s.io/code-generator@v0.22.3
//go:generate bash -c "chmod +x $GOPATH/pkg/mod/k8s.io/code-generator@v0.22.3/generate-groups.sh"
//go:generate $GOPATH/pkg/mod/k8s.io/code-generator@v0.22.3/generate-groups.sh all github.com/networkservicemesh/sdk-k8s/pkg/tools/k8s/client github.com/networkservicemesh/sdk-k8s/pkg/tools/k8s/apis networkservicemesh.io:v1 --go-header-file ./../../../copyright.txt
//go:generate goimports -w -local github.com/networkservicemesh -d "./client"
//go:generate go mod tidy
