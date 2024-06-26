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
)

// HKTA = HKT<A>
// HKTOA = HKT<Option<A>>
//
// Sequence converts an option of some higher kinded type into the higher kinded type of an option
func Sequence[A, HKTA, HKTOA any](
	_of func(Option[A]) HKTOA,
	_map func(HKTA, func(A) Option[A]) HKTOA,
) func(Option[HKTA]) HKTOA {
	return Fold(F.Nullary2(None[A], _of), F.Bind2nd(_map, Some[A]))
}
