/*
Copyright 2019 The Kubernetes Authors.

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

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	kops "k8s.io/kops/pkg/apis/kops"
	scheme "k8s.io/kops/pkg/client/clientset_generated/internalclientset/scheme"
)

// KeysetsGetter has a method to return a KeysetInterface.
// A group's client should implement this interface.
type KeysetsGetter interface {
	Keysets(namespace string) KeysetInterface
}

// KeysetInterface has methods to work with Keyset resources.
type KeysetInterface interface {
	Create(*kops.Keyset) (*kops.Keyset, error)
	Update(*kops.Keyset) (*kops.Keyset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*kops.Keyset, error)
	List(opts v1.ListOptions) (*kops.KeysetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kops.Keyset, err error)
	KeysetExpansion
}

// keysets implements KeysetInterface
type keysets struct {
	client rest.Interface
	ns     string
}

// newKeysets returns a Keysets
func newKeysets(c *KopsClient, namespace string) *keysets {
	return &keysets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the keyset, and returns the corresponding keyset object, and an error if there is any.
func (c *keysets) Get(name string, options v1.GetOptions) (result *kops.Keyset, err error) {
	result = &kops.Keyset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keysets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Keysets that match those selectors.
func (c *keysets) List(opts v1.ListOptions) (result *kops.KeysetList, err error) {
	result = &kops.KeysetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keysets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested keysets.
func (c *keysets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("keysets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a keyset and creates it.  Returns the server's representation of the keyset, and an error, if there is any.
func (c *keysets) Create(keyset *kops.Keyset) (result *kops.Keyset, err error) {
	result = &kops.Keyset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("keysets").
		Body(keyset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a keyset and updates it. Returns the server's representation of the keyset, and an error, if there is any.
func (c *keysets) Update(keyset *kops.Keyset) (result *kops.Keyset, err error) {
	result = &kops.Keyset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keysets").
		Name(keyset.Name).
		Body(keyset).
		Do().
		Into(result)
	return
}

// Delete takes name of the keyset and deletes it. Returns an error if one occurs.
func (c *keysets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keysets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *keysets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keysets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched keyset.
func (c *keysets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kops.Keyset, err error) {
	result = &kops.Keyset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("keysets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
