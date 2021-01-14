// Copyright (c) 2021 Doc.ai and/or its affiliates.
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
package etcd_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/networkservicemesh/api/pkg/api/registry"
	"github.com/networkservicemesh/sdk/pkg/registry/core/adapters"
	"github.com/stretchr/testify/require"

	"github.com/networkservicemesh/sdk-k8s/pkg/registry/etcd"
	"github.com/networkservicemesh/sdk-k8s/pkg/tools/k8s/client/clientset/versioned/fake"
)

func Test_NSE_Reregister(t *testing.T) {
	s := etcd.NewNetworkServiceEndpointRegistryServer(context.Background(), "default", fake.NewSimpleClientset())
	_, err := s.Register(context.Background(), &registry.NetworkServiceEndpoint{Name: "nse-1"})
	require.NoError(t, err)
	_, err = s.Register(context.Background(), &registry.NetworkServiceEndpoint{Name: "nse-1", NetworkServiceNames: []string{"ns-1"}})
	require.NoError(t, err)
}

func Test_NSE_Find(t *testing.T) {
	s := etcd.NewNetworkServiceEndpointRegistryServer(context.Background(), "default", fake.NewSimpleClientset())
	_, err := s.Register(context.Background(), &registry.NetworkServiceEndpoint{Name: "nse-1", NetworkServiceNames: []string{"ns-1"}})
	require.NoError(t, err)
	_, err = s.Register(context.Background(), &registry.NetworkServiceEndpoint{Name: "nse-2", NetworkServiceNames: []string{"ns-1"}})
	require.NoError(t, err)

	client := adapters.NetworkServiceEndpointServerToClient(s)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := client.Find(ctx, &registry.NetworkServiceEndpointQuery{Watch: true, NetworkServiceEndpoint: &registry.NetworkServiceEndpoint{NetworkServiceNames: []string{"ns-1"}}})
	require.NoError(t, err)

	ch := registry.ReadNetworkServiceEndpointChannel(stream)

	nse1 := <-ch
	nse2 := <-ch
	go func() {
		time.Sleep(time.Second / 2)
		_, _ = s.Register(context.Background(), &registry.NetworkServiceEndpoint{Name: "nse-3", NetworkServiceNames: []string{"ns-1"}})

	}()
	nse3 := <-ch

	fmt.Println(nse1)
	fmt.Println(nse2)
	fmt.Println(nse3)
}
