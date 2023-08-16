// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kosmos.io/clusterlink/pkg/apis/clusterlink/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterNodeLister helps list ClusterNodes.
// All objects returned here must be treated as read-only.
type ClusterNodeLister interface {
	// List lists all ClusterNodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterNode, err error)
	// Get retrieves the ClusterNode from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterNode, error)
	ClusterNodeListerExpansion
}

// clusterNodeLister implements the ClusterNodeLister interface.
type clusterNodeLister struct {
	indexer cache.Indexer
}

// NewClusterNodeLister returns a new ClusterNodeLister.
func NewClusterNodeLister(indexer cache.Indexer) ClusterNodeLister {
	return &clusterNodeLister{indexer: indexer}
}

// List lists all ClusterNodes in the indexer.
func (s *clusterNodeLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterNode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterNode))
	})
	return ret, err
}

// Get retrieves the ClusterNode from the index for a given name.
func (s *clusterNodeLister) Get(name string) (*v1alpha1.ClusterNode, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusternode"), name)
	}
	return obj.(*v1alpha1.ClusterNode), nil
}