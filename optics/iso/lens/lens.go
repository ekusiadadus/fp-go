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

package lens

import (
	F "github.com/IBM/fp-go/function"
	I "github.com/IBM/fp-go/optics/iso"
	L "github.com/IBM/fp-go/optics/lens"
)

// IsoAsLens converts an `Iso` to a `Lens`
func IsoAsLens[S, A any](sa I.Iso[S, A]) L.Lens[S, A] {
	return L.MakeLensCurried(sa.Get, F.Flow2(sa.ReverseGet, F.Constant1[S, S]))
}

// IsoAsLensRef converts an `Iso` to a `Lens`
func IsoAsLensRef[S, A any](sa I.Iso[*S, A]) L.Lens[*S, A] {
	return L.MakeLensRefCurried(sa.Get, F.Flow2(sa.ReverseGet, F.Constant1[*S, *S]))
}
