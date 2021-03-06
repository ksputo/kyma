// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/helm-broker/pkg/apis/addons/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterAddonsConfigurationLister helps list ClusterAddonsConfigurations.
type ClusterAddonsConfigurationLister interface {
	// List lists all ClusterAddonsConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterAddonsConfiguration, err error)
	// Get retrieves the ClusterAddonsConfiguration from the index for a given name.
	Get(name string) (*v1alpha1.ClusterAddonsConfiguration, error)
	ClusterAddonsConfigurationListerExpansion
}

// clusterAddonsConfigurationLister implements the ClusterAddonsConfigurationLister interface.
type clusterAddonsConfigurationLister struct {
	indexer cache.Indexer
}

// NewClusterAddonsConfigurationLister returns a new ClusterAddonsConfigurationLister.
func NewClusterAddonsConfigurationLister(indexer cache.Indexer) ClusterAddonsConfigurationLister {
	return &clusterAddonsConfigurationLister{indexer: indexer}
}

// List lists all ClusterAddonsConfigurations in the indexer.
func (s *clusterAddonsConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterAddonsConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterAddonsConfiguration))
	})
	return ret, err
}

// Get retrieves the ClusterAddonsConfiguration from the index for a given name.
func (s *clusterAddonsConfigurationLister) Get(name string) (*v1alpha1.ClusterAddonsConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusteraddonsconfiguration"), name)
	}
	return obj.(*v1alpha1.ClusterAddonsConfiguration), nil
}
