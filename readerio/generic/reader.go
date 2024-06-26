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

package generic

import (
	F "github.com/IBM/fp-go/function"
	FR "github.com/IBM/fp-go/internal/fromreader"
	"github.com/IBM/fp-go/internal/readert"
	IO "github.com/IBM/fp-go/io/generic"
	R "github.com/IBM/fp-go/reader/generic"
)

func FromIO[GEA ~func(E) GIOA, GIOA ~func() A, E, A any](t GIOA) GEA {
	return R.Of[GEA, E](t)
}

func FromReader[GA ~func(E) A, GEA ~func(E) GIOA, GIOA ~func() A, E, A any](r GA) GEA {
	return readert.MonadFromReader[GA, GEA](IO.Of[GIOA, A], r)
}

func MonadMap[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](fa GEA, f func(A) B) GEB {
	return readert.MonadMap[GEA, GEB](IO.MonadMap[GIOA, GIOB, A, B], fa, f)
}

func Map[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](f func(A) B) func(GEA) GEB {
	return F.Bind2nd(MonadMap[GEA, GEB, GIOA, GIOB, E, A, B], f)
}

func MonadChain[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](ma GEA, f func(A) GEB) GEB {
	return readert.MonadChain(IO.MonadChain[GIOA, GIOB, A, B], ma, f)
}

func Chain[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](f func(A) GEB) func(GEA) GEB {
	return F.Bind2nd(MonadChain[GEA, GEB, GIOA, GIOB, E, A, B], f)
}

func Of[GEA ~func(E) GIOA, GIOA ~func() A, E, A any](a A) GEA {
	return readert.MonadOf[GEA](IO.Of[GIOA, A], a)
}

func MonadAp[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fab GEFAB, fa GEA) GEB {
	return readert.MonadAp[GEA, GEB, GEFAB, E, A](IO.MonadAp[GIOA, GIOB, GIOFAB, A, B], fab, fa)
}

func Ap[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fa GEA) func(GEFAB) GEB {
	return F.Bind2nd(MonadAp[GEA, GEB, GEFAB, GIOA, GIOB, GIOFAB, E, A, B], fa)
}

func MonadApSeq[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fab GEFAB, fa GEA) GEB {
	return readert.MonadAp[GEA, GEB, GEFAB, E, A](IO.MonadApSeq[GIOA, GIOB, GIOFAB, A, B], fab, fa)
}

func ApSeq[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fa GEA) func(GEFAB) GEB {
	return F.Bind2nd(MonadApSeq[GEA, GEB, GEFAB, GIOA, GIOB, GIOFAB, E, A, B], fa)
}

func MonadApPar[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fab GEFAB, fa GEA) GEB {
	return readert.MonadAp[GEA, GEB, GEFAB, E, A](IO.MonadApPar[GIOA, GIOB, GIOFAB, A, B], fab, fa)
}

func ApPar[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fa GEA) func(GEFAB) GEB {
	return F.Bind2nd(MonadApPar[GEA, GEB, GEFAB, GIOA, GIOB, GIOFAB, E, A, B], fa)
}

func Ask[GEE ~func(E) GIOE, GIOE ~func() E, E any]() GEE {
	return FR.Ask(FromReader[func(E) E, GEE, GIOE, E, E])()
}

func Asks[GA ~func(E) A, GEA ~func(E) GIOA, GIOA ~func() A, E, A any](r GA) GEA {
	return FR.Asks(FromReader[GA, GEA, GIOA, E, A])(r)
}

func MonadChainIOK[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](ma GEA, f func(A) GIOB) GEB {
	return MonadChain(ma, F.Flow2(f, FromIO[GEB]))
}

func ChainIOK[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](f func(A) GIOB) func(GEA) GEB {
	return F.Bind2nd(MonadChainIOK[GEA, GEB, GIOA, GIOB, E, A, B], f)
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[GEA ~func(E) GA, GA ~func() A, E, A any](gen func() GEA) GEA {
	return func(e E) GA {
		return func() A {
			return gen()(e)()
		}
	}
}
