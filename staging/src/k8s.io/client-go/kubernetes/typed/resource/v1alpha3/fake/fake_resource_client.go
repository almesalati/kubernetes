/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha3 "k8s.io/client-go/kubernetes/typed/resource/v1alpha3"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeResourceV1alpha3 struct {
	*testing.Fake
}

func (c *FakeResourceV1alpha3) DeviceClasses() v1alpha3.DeviceClassInterface {
	return newFakeDeviceClasses(c)
}

func (c *FakeResourceV1alpha3) DeviceTaintRules() v1alpha3.DeviceTaintRuleInterface {
	return newFakeDeviceTaintRules(c)
}

func (c *FakeResourceV1alpha3) ResourceClaims(namespace string) v1alpha3.ResourceClaimInterface {
	return newFakeResourceClaims(c, namespace)
}

func (c *FakeResourceV1alpha3) ResourceClaimTemplates(namespace string) v1alpha3.ResourceClaimTemplateInterface {
	return newFakeResourceClaimTemplates(c, namespace)
}

func (c *FakeResourceV1alpha3) ResourceSlices() v1alpha3.ResourceSliceInterface {
	return newFakeResourceSlices(c)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeResourceV1alpha3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
