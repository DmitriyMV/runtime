// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package typed

import (
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta/spec"
)

// DeepCopyable requires a spec to have DeepCopy method which will be used during Resource copy.
type DeepCopyable[T any] interface {
	DeepCopy() T
}

// ResourceDefinition is a phantom type which acts as info supplier for Resource String and ResourceDefinition
// methods. It intantianed only during String and ResourceDefinition calls, so it should never contain any data which
// survives those calls. Any empty struct{} will do.
type ResourceDefinition[T any] interface {
	ResourceDefinition(md resource.Metadata, spec T) spec.ResourceDefinitionSpec
}

// Resource provides a generic base implementation for resource.Resource.
type Resource[T DeepCopyable[T], RD ResourceDefinition[T]] struct {
	spec T
	md   resource.Metadata
}

// Metadata implements Resource.
func (t *Resource[T, RD]) Metadata() *resource.Metadata {
	return &t.md
}

// Spec implements resource.Resource.
func (t *Resource[T, RD]) Spec() interface{} {
	return t.spec
}

// TypedSpec returns a pointer to spec field.
func (t *Resource[T, RD]) TypedSpec() *T {
	return &t.spec
}

// DeepCopy returns a deep copy of Resource.
func (t *Resource[T, RD]) DeepCopy() resource.Resource { //nolint:ireturn
	return &Resource[T, RD]{t.spec.DeepCopy(), t.md}
}

// ResourceDefinition implements spec.ResourceDefinitionProvider interface.
func (t *Resource[T, RD]) ResourceDefinition() spec.ResourceDefinitionSpec {
	var zero RD

	return zero.ResourceDefinition(t.md, t.spec)
}

// NewResource initializes and returns a new instance of Resource with typed spec field.
func NewResource[T DeepCopyable[T], RD ResourceDefinition[T]](md resource.Metadata, spec T) *Resource[T, RD] {
	result := Resource[T, RD]{md: md, spec: spec}
	result.md.BumpVersion()

	return &result
}