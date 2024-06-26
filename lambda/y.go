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

package lambda

type (
	// RecFct is the function called recursively
	RecFct[T, R any] func(T) R

	internalCombinator[T, R any] func(internalCombinator[T, R]) RecFct[T, R]
)

// Y is the Y-combinator based on https://dreamsongs.com/Files/WhyOfY.pdf
func Y[TRFRM ~func(RecFct[T, R]) RecFct[T, R], T, R any](f TRFRM) RecFct[T, R] {
	g := func(h internalCombinator[T, R]) RecFct[T, R] {
		return func(t T) R {
			return f(h(h))(t)
		}
	}
	return g(g)
}
