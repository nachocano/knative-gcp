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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/google/knative-gcp/pkg/apis/events/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloudPubSubSourceLister helps list CloudPubSubSources.
type CloudPubSubSourceLister interface {
	// List lists all CloudPubSubSources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudPubSubSource, err error)
	// CloudPubSubSources returns an object that can list and get CloudPubSubSources.
	CloudPubSubSources(namespace string) CloudPubSubSourceNamespaceLister
	CloudPubSubSourceListerExpansion
}

// cloudPubSubSourceLister implements the CloudPubSubSourceLister interface.
type cloudPubSubSourceLister struct {
	indexer cache.Indexer
}

// NewCloudPubSubSourceLister returns a new CloudPubSubSourceLister.
func NewCloudPubSubSourceLister(indexer cache.Indexer) CloudPubSubSourceLister {
	return &cloudPubSubSourceLister{indexer: indexer}
}

// List lists all CloudPubSubSources in the indexer.
func (s *cloudPubSubSourceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudPubSubSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudPubSubSource))
	})
	return ret, err
}

// CloudPubSubSources returns an object that can list and get CloudPubSubSources.
func (s *cloudPubSubSourceLister) CloudPubSubSources(namespace string) CloudPubSubSourceNamespaceLister {
	return cloudPubSubSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudPubSubSourceNamespaceLister helps list and get CloudPubSubSources.
type CloudPubSubSourceNamespaceLister interface {
	// List lists all CloudPubSubSources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudPubSubSource, err error)
	// Get retrieves the CloudPubSubSource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudPubSubSource, error)
	CloudPubSubSourceNamespaceListerExpansion
}

// cloudPubSubSourceNamespaceLister implements the CloudPubSubSourceNamespaceLister
// interface.
type cloudPubSubSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudPubSubSources in the indexer for a given namespace.
func (s cloudPubSubSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudPubSubSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudPubSubSource))
	})
	return ret, err
}

// Get retrieves the CloudPubSubSource from the indexer for a given namespace and name.
func (s cloudPubSubSourceNamespaceLister) Get(name string) (*v1alpha1.CloudPubSubSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudpubsubsource"), name)
	}
	return obj.(*v1alpha1.CloudPubSubSource), nil
}
