// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/console/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConsoleSampleLister helps list ConsoleSamples.
// All objects returned here must be treated as read-only.
type ConsoleSampleLister interface {
	// List lists all ConsoleSamples in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ConsoleSample, err error)
	// Get retrieves the ConsoleSample from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ConsoleSample, error)
	ConsoleSampleListerExpansion
}

// consoleSampleLister implements the ConsoleSampleLister interface.
type consoleSampleLister struct {
	indexer cache.Indexer
}

// NewConsoleSampleLister returns a new ConsoleSampleLister.
func NewConsoleSampleLister(indexer cache.Indexer) ConsoleSampleLister {
	return &consoleSampleLister{indexer: indexer}
}

// List lists all ConsoleSamples in the indexer.
func (s *consoleSampleLister) List(selector labels.Selector) (ret []*v1.ConsoleSample, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ConsoleSample))
	})
	return ret, err
}

// Get retrieves the ConsoleSample from the index for a given name.
func (s *consoleSampleLister) Get(name string) (*v1.ConsoleSample, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("consolesample"), name)
	}
	return obj.(*v1.ConsoleSample), nil
}
