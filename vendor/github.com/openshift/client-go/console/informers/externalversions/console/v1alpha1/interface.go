// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openshift/client-go/console/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ConsolePlugins returns a ConsolePluginInformer.
	ConsolePlugins() ConsolePluginInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ConsolePlugins returns a ConsolePluginInformer.
func (v *version) ConsolePlugins() ConsolePluginInformer {
	return &consolePluginInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}