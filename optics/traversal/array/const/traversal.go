// Copyright (c) 2023 IBM Corp.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identity

import (
	C "github.com/IBM/fp-go/constant"
	M "github.com/IBM/fp-go/monoid"
	AR "github.com/IBM/fp-go/optics/traversal/array/generic/const"
	G "github.com/IBM/fp-go/optics/traversal/generic"
)

// FromArray returns a traversal from an array for the identity monad
func FromArray[E, A any](m M.Monoid[E]) G.Traversal[[]A, A, C.Const[E, []A], C.Const[E, A]] {
	return AR.FromArray[[]A, E, A](m)
}
