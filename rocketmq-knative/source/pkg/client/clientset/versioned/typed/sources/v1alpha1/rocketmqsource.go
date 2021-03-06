/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

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
	"time"

	v1alpha1 "github.com/apache/rocketmq-externals/rocketmq-knative/source/pkg/apis/sources/v1alpha1"
	scheme "github.com/apache/rocketmq-externals/rocketmq-knative/source/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RocketMQSourcesGetter has a method to return a RocketMQSourceInterface.
// A group's client should implement this interface.
type RocketMQSourcesGetter interface {
	RocketMQSources(namespace string) RocketMQSourceInterface
}

// RocketMQSourceInterface has methods to work with RocketMQSource resources.
type RocketMQSourceInterface interface {
	Create(source *v1alpha1.RocketMQSource) (*v1alpha1.RocketMQSource, error)
	Update(*v1alpha1.RocketMQSource) (*v1alpha1.RocketMQSource, error)
	UpdateStatus(*v1alpha1.RocketMQSource) (*v1alpha1.RocketMQSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RocketMQSource, error)
	List(opts v1.ListOptions) (*v1alpha1.RocketMQSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RocketMQSource, err error)
	RocketMQSourceExpansion
}

type rocketmqSources struct {
	client rest.Interface
	ns     string
}

// newRocketMQSources returns a RocketMQSources
func newRocketMQSources(c *SourcesV1alpha1Client, namespace string) *rocketmqSources {
	return &rocketmqSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rocketmqSource, and returns the corresponding rocketmqSource object, and an error if there is any.
func (c *rocketmqSources) Get(name string, options v1.GetOptions) (result *v1alpha1.RocketMQSource, err error) {
	result = &v1alpha1.RocketMQSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rocketmqsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RocketMQSources that match those selectors.
func (c *rocketmqSources) List(opts v1.ListOptions) (result *v1alpha1.RocketMQSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RocketMQSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rocketmqsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch rocketmqSources a watch.Interface that watches the requested rocketmqSources.
func (c *rocketmqSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rocketmqsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a rocketmqSource and creates it.  Returns the server's representation of the rocketmqSources, and an error, if there is any.
func (c *rocketmqSources) Create(rocketmqSource *v1alpha1.RocketMQSource) (result *v1alpha1.RocketMQSource, err error) {
	result = &v1alpha1.RocketMQSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rocketmqsources").
		Body(rocketmqSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rocketmqSources and updates it. Returns the server's representation of the rocketmqSources, and an error, if there is any.
func (c *rocketmqSources) Update(rocketmqSource *v1alpha1.RocketMQSource) (result *v1alpha1.RocketMQSource, err error) {
	result = &v1alpha1.RocketMQSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rocketmqsources").
		Name(rocketmqSource.Name).
		Body(rocketmqSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *rocketmqSources) UpdateStatus(rocketmqSource *v1alpha1.RocketMQSource) (result *v1alpha1.RocketMQSource, err error) {
	result = &v1alpha1.RocketMQSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rocketmqsources").
		Name(rocketmqSource.Name).
		SubResource("status").
		Body(rocketmqSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the rocketmqsources and deletes it. Returns an error if one occurs.
func (c *rocketmqSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rocketmqsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rocketmqSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rocketmqsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched rocketmqSources.
func (c *rocketmqSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.rocketmqSources, err error) {
	result = &v1alpha1.rocketmqSources{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rocketmqsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
