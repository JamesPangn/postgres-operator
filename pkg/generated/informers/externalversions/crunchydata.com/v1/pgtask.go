/*
 Copyright 2021 Crunchy Data Solutions, Inc.
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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	crunchydatacomv1 "github.com/crunchydata/postgres-operator/pkg/apis/crunchydata.com/v1"
	versioned "github.com/crunchydata/postgres-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/crunchydata/postgres-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/crunchydata/postgres-operator/pkg/generated/listers/crunchydata.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PgtaskInformer provides access to a shared informer and lister for
// Pgtasks.
type PgtaskInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PgtaskLister
}

type pgtaskInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPgtaskInformer constructs a new informer for Pgtask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPgtaskInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPgtaskInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPgtaskInformer constructs a new informer for Pgtask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPgtaskInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrunchydataV1().Pgtasks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrunchydataV1().Pgtasks(namespace).Watch(context.TODO(), options)
			},
		},
		&crunchydatacomv1.Pgtask{},
		resyncPeriod,
		indexers,
	)
}

func (f *pgtaskInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPgtaskInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pgtaskInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crunchydatacomv1.Pgtask{}, f.defaultInformer)
}

func (f *pgtaskInformer) Lister() v1.PgtaskLister {
	return v1.NewPgtaskLister(f.Informer().GetIndexer())
}
