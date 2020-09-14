/*
Copyright 2020 Google LLC

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

package v1

import (
	"context"
	"time"

	v1 "github.com/google/knative-gcp/pkg/apis/events/v1"
	scheme "github.com/google/knative-gcp/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudSchedulerSourcesGetter has a method to return a CloudSchedulerSourceInterface.
// A group's client should implement this interface.
type CloudSchedulerSourcesGetter interface {
	CloudSchedulerSources(namespace string) CloudSchedulerSourceInterface
}

// CloudSchedulerSourceInterface has methods to work with CloudSchedulerSource resources.
type CloudSchedulerSourceInterface interface {
	Create(ctx context.Context, cloudSchedulerSource *v1.CloudSchedulerSource, opts metav1.CreateOptions) (*v1.CloudSchedulerSource, error)
	Update(ctx context.Context, cloudSchedulerSource *v1.CloudSchedulerSource, opts metav1.UpdateOptions) (*v1.CloudSchedulerSource, error)
	UpdateStatus(ctx context.Context, cloudSchedulerSource *v1.CloudSchedulerSource, opts metav1.UpdateOptions) (*v1.CloudSchedulerSource, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CloudSchedulerSource, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CloudSchedulerSourceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CloudSchedulerSource, err error)
	CloudSchedulerSourceExpansion
}

// cloudSchedulerSources implements CloudSchedulerSourceInterface
type cloudSchedulerSources struct {
	client rest.Interface
	ns     string
}

// newCloudSchedulerSources returns a CloudSchedulerSources
func newCloudSchedulerSources(c *EventsV1Client, namespace string) *cloudSchedulerSources {
	return &cloudSchedulerSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudSchedulerSource, and returns the corresponding cloudSchedulerSource object, and an error if there is any.
func (c *cloudSchedulerSources) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CloudSchedulerSource, err error) {
	result = &v1.CloudSchedulerSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudSchedulerSources that match those selectors.
func (c *cloudSchedulerSources) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CloudSchedulerSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CloudSchedulerSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudSchedulerSources.
func (c *cloudSchedulerSources) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cloudSchedulerSource and creates it.  Returns the server's representation of the cloudSchedulerSource, and an error, if there is any.
func (c *cloudSchedulerSources) Create(ctx context.Context, cloudSchedulerSource *v1.CloudSchedulerSource, opts metav1.CreateOptions) (result *v1.CloudSchedulerSource, err error) {
	result = &v1.CloudSchedulerSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudSchedulerSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cloudSchedulerSource and updates it. Returns the server's representation of the cloudSchedulerSource, and an error, if there is any.
func (c *cloudSchedulerSources) Update(ctx context.Context, cloudSchedulerSource *v1.CloudSchedulerSource, opts metav1.UpdateOptions) (result *v1.CloudSchedulerSource, err error) {
	result = &v1.CloudSchedulerSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Name(cloudSchedulerSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudSchedulerSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cloudSchedulerSources) UpdateStatus(ctx context.Context, cloudSchedulerSource *v1.CloudSchedulerSource, opts metav1.UpdateOptions) (result *v1.CloudSchedulerSource, err error) {
	result = &v1.CloudSchedulerSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Name(cloudSchedulerSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudSchedulerSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cloudSchedulerSource and deletes it. Returns an error if one occurs.
func (c *cloudSchedulerSources) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudSchedulerSources) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cloudSchedulerSource.
func (c *cloudSchedulerSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CloudSchedulerSource, err error) {
	result = &v1.CloudSchedulerSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
