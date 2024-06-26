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

package option

import (
	F "github.com/IBM/fp-go/function"
	RA "github.com/IBM/fp-go/internal/array"
)

// TraverseArray transforms an array
func TraverseArrayG[GA ~[]A, GB ~[]B, A, B any](f func(A) Option[B]) func(GA) Option[GB] {
	return RA.Traverse[GA](
		Of[GB],
		Map[GB, func(B) GB],
		Ap[GB, B],

		f,
	)
}

// TraverseArray transforms an array
func TraverseArray[A, B any](f func(A) Option[B]) func([]A) Option[[]B] {
	return TraverseArrayG[[]A, []B](f)
}

func SequenceArrayG[GA ~[]A, GOA ~[]Option[A], A any](ma GOA) Option[GA] {
	return TraverseArrayG[GOA, GA](F.Identity[Option[A]])(ma)
}

// SequenceArray converts a homogeneous sequence of either into an either of sequence
func SequenceArray[A any](ma []Option[A]) Option[[]A] {
	return SequenceArrayG[[]A](ma)
}
