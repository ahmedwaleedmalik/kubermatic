/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package seedsync

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"k8c.io/kubermatic/v2/pkg/controller/util/predicate"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/provider"

	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	// ControllerName is the name of this very controller.
	ControllerName = "seed-sync-controller"

	// ManagedByLabel is the label used to identify the resources
	// created by this controller.
	ManagedByLabel = "app.kubernetes.io/managed-by"

	// CleanupFinalizer is put on Seed CRs to facilitate proper
	// cleanup when a Seed is deleted.
	CleanupFinalizer = "kubermatic.io/cleanup-seed-sync"
)

// Add creates a new Seed-Sync controller and sets up Watches
func Add(
	ctx context.Context,
	mgr manager.Manager,
	numWorkers int,
	log *zap.SugaredLogger,
	namespace string,
	seedKubeconfigGetter provider.SeedKubeconfigGetter,
) error {
	reconciler := &Reconciler{
		Client:           mgr.GetClient(),
		recorder:         mgr.GetEventRecorderFor(ControllerName),
		log:              log.Named(ControllerName),
		seedClientGetter: provider.SeedClientGetterFactory(seedKubeconfigGetter),
	}

	ctrlOptions := controller.Options{Reconciler: reconciler, MaxConcurrentReconciles: numWorkers}
	c, err := controller.New(ControllerName, mgr, ctrlOptions)
	if err != nil {
		return err
	}

	// watch all seeds in the given namespace
	if err := c.Watch(&source.Kind{Type: &kubermaticv1.Seed{}}, &handler.EnqueueRequestForObject{}, predicate.ByNamespace(namespace)); err != nil {
		return fmt.Errorf("failed to create watcher: %v", err)
	}

	return nil
}
