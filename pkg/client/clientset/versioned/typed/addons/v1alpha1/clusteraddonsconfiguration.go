/*
Copyright 2020 The Helm Broker Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kyma-project/helm-broker/pkg/apis/addons/v1alpha1"
	scheme "github.com/kyma-project/helm-broker/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterAddonsConfigurationsGetter has a method to return a ClusterAddonsConfigurationInterface.
// A group's client should implement this interface.
type ClusterAddonsConfigurationsGetter interface {
	ClusterAddonsConfigurations() ClusterAddonsConfigurationInterface
}

// ClusterAddonsConfigurationInterface has methods to work with ClusterAddonsConfiguration resources.
type ClusterAddonsConfigurationInterface interface {
	Create(ctx context.Context, clusterAddonsConfiguration *v1alpha1.ClusterAddonsConfiguration, opts v1.CreateOptions) (*v1alpha1.ClusterAddonsConfiguration, error)
	Update(ctx context.Context, clusterAddonsConfiguration *v1alpha1.ClusterAddonsConfiguration, opts v1.UpdateOptions) (*v1alpha1.ClusterAddonsConfiguration, error)
	UpdateStatus(ctx context.Context, clusterAddonsConfiguration *v1alpha1.ClusterAddonsConfiguration, opts v1.UpdateOptions) (*v1alpha1.ClusterAddonsConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterAddonsConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterAddonsConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterAddonsConfiguration, err error)
	ClusterAddonsConfigurationExpansion
}

// clusterAddonsConfigurations implements ClusterAddonsConfigurationInterface
type clusterAddonsConfigurations struct {
	client rest.Interface
}

// newClusterAddonsConfigurations returns a ClusterAddonsConfigurations
func newClusterAddonsConfigurations(c *AddonsV1alpha1Client) *clusterAddonsConfigurations {
	return &clusterAddonsConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterAddonsConfiguration, and returns the corresponding clusterAddonsConfiguration object, and an error if there is any.
func (c *clusterAddonsConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterAddonsConfiguration, err error) {
	result = &v1alpha1.ClusterAddonsConfiguration{}
	err = c.client.Get().
		Resource("clusteraddonsconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterAddonsConfigurations that match those selectors.
func (c *clusterAddonsConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterAddonsConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterAddonsConfigurationList{}
	err = c.client.Get().
		Resource("clusteraddonsconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterAddonsConfigurations.
func (c *clusterAddonsConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusteraddonsconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterAddonsConfiguration and creates it.  Returns the server's representation of the clusterAddonsConfiguration, and an error, if there is any.
func (c *clusterAddonsConfigurations) Create(ctx context.Context, clusterAddonsConfiguration *v1alpha1.ClusterAddonsConfiguration, opts v1.CreateOptions) (result *v1alpha1.ClusterAddonsConfiguration, err error) {
	result = &v1alpha1.ClusterAddonsConfiguration{}
	err = c.client.Post().
		Resource("clusteraddonsconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAddonsConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterAddonsConfiguration and updates it. Returns the server's representation of the clusterAddonsConfiguration, and an error, if there is any.
func (c *clusterAddonsConfigurations) Update(ctx context.Context, clusterAddonsConfiguration *v1alpha1.ClusterAddonsConfiguration, opts v1.UpdateOptions) (result *v1alpha1.ClusterAddonsConfiguration, err error) {
	result = &v1alpha1.ClusterAddonsConfiguration{}
	err = c.client.Put().
		Resource("clusteraddonsconfigurations").
		Name(clusterAddonsConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAddonsConfiguration).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterAddonsConfigurations) UpdateStatus(ctx context.Context, clusterAddonsConfiguration *v1alpha1.ClusterAddonsConfiguration, opts v1.UpdateOptions) (result *v1alpha1.ClusterAddonsConfiguration, err error) {
	result = &v1alpha1.ClusterAddonsConfiguration{}
	err = c.client.Put().
		Resource("clusteraddonsconfigurations").
		Name(clusterAddonsConfiguration.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAddonsConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterAddonsConfiguration and deletes it. Returns an error if one occurs.
func (c *clusterAddonsConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusteraddonsconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterAddonsConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusteraddonsconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterAddonsConfiguration.
func (c *clusterAddonsConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterAddonsConfiguration, err error) {
	result = &v1alpha1.ClusterAddonsConfiguration{}
	err = c.client.Patch(pt).
		Resource("clusteraddonsconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
