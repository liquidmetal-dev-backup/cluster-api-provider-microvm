// Copyright 2024 Liquid Metal Authors. All Rights Reserved.
// SPDX-License-Identifier: MPL-2.0

package scope

// Scoper is the interface for a scope.
type Scoper interface {
	// Name returns the name of the resource.
	Name() string
	// Namespace returns the resources namespace.
	Namespace() string
	// ClusterName returns the name of the cluster.
	ClusterName() string

	// ControllerName returns the name of the controller that created the scope.
	ControllerName() string

	// Patch persists the resource and status.
	Patch() error
}
