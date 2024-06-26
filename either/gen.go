// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-08-11 11:37:35.9608486 +0200 CEST m=+0.009664501

package either

import (
	A "github.com/IBM/fp-go/internal/apply"
	T "github.com/IBM/fp-go/tuple"
)

// Eitherize0 converts a function with 0 parameters returning a tuple into a function with 0 parameters returning an Either
// The inverse function is [Uneitherize0]
func Eitherize0[F ~func() (R, error), R any](f F) func() Either[error, R] {
	return func() Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f()
		})
	}
}

// Uneitherize0 converts a function with 0 parameters returning an Either into a function with 0 parameters returning a tuple
// The inverse function is [Eitherize0]
func Uneitherize0[F ~func() Either[error, R], R any](f F) func() (R, error) {
	return func() (R, error) {
		return UnwrapError(f())
	}
}

// Eitherize1 converts a function with 1 parameters returning a tuple into a function with 1 parameters returning an Either
// The inverse function is [Uneitherize1]
func Eitherize1[F ~func(T0) (R, error), T0, R any](f F) func(T0) Either[error, R] {
	return func(t0 T0) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0)
		})
	}
}

// Uneitherize1 converts a function with 1 parameters returning an Either into a function with 1 parameters returning a tuple
// The inverse function is [Eitherize1]
func Uneitherize1[F ~func(T0) Either[error, R], T0, R any](f F) func(T0) (R, error) {
	return func(t0 T0) (R, error) {
		return UnwrapError(f(t0))
	}
}

// SequenceT1 converts 1 parameters of [Either[E, T]] into a [Either[E, Tuple1]].
func SequenceT1[E, T1 any](t1 Either[E, T1]) Either[E, T.Tuple1[T1]] {
	return A.SequenceT1(
		Map[E, T1, T.Tuple1[T1]],
		t1,
	)
}

// SequenceTuple1 converts a [Tuple1] of [Either[E, T]] into an [Either[E, Tuple1]].
func SequenceTuple1[E, T1 any](t T.Tuple1[Either[E, T1]]) Either[E, T.Tuple1[T1]] {
	return A.SequenceTuple1(
		Map[E, T1, T.Tuple1[T1]],
		t,
	)
}

// TraverseTuple1 converts a [Tuple1] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple1]].
func TraverseTuple1[F1 ~func(A1) Either[E, T1], E, A1, T1 any](f1 F1) func(T.Tuple1[A1]) Either[E, T.Tuple1[T1]] {
	return func(t T.Tuple1[A1]) Either[E, T.Tuple1[T1]] {
		return A.TraverseTuple1(
			Map[E, T1, T.Tuple1[T1]],
			f1,
			t,
		)
	}
}

// Eitherize2 converts a function with 2 parameters returning a tuple into a function with 2 parameters returning an Either
// The inverse function is [Uneitherize2]
func Eitherize2[F ~func(T0, T1) (R, error), T0, T1, R any](f F) func(T0, T1) Either[error, R] {
	return func(t0 T0, t1 T1) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0, t1)
		})
	}
}

// Uneitherize2 converts a function with 2 parameters returning an Either into a function with 2 parameters returning a tuple
// The inverse function is [Eitherize2]
func Uneitherize2[F ~func(T0, T1) Either[error, R], T0, T1, R any](f F) func(T0, T1) (R, error) {
	return func(t0 T0, t1 T1) (R, error) {
		return UnwrapError(f(t0, t1))
	}
}

// SequenceT2 converts 2 parameters of [Either[E, T]] into a [Either[E, Tuple2]].
func SequenceT2[E, T1, T2 any](t1 Either[E, T1], t2 Either[E, T2]) Either[E, T.Tuple2[T1, T2]] {
	return A.SequenceT2(
		Map[E, T1, func(T2) T.Tuple2[T1, T2]],
		Ap[T.Tuple2[T1, T2], E, T2],
		t1,
		t2,
	)
}

// SequenceTuple2 converts a [Tuple2] of [Either[E, T]] into an [Either[E, Tuple2]].
func SequenceTuple2[E, T1, T2 any](t T.Tuple2[Either[E, T1], Either[E, T2]]) Either[E, T.Tuple2[T1, T2]] {
	return A.SequenceTuple2(
		Map[E, T1, func(T2) T.Tuple2[T1, T2]],
		Ap[T.Tuple2[T1, T2], E, T2],
		t,
	)
}

// TraverseTuple2 converts a [Tuple2] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple2]].
func TraverseTuple2[F1 ~func(A1) Either[E, T1], F2 ~func(A2) Either[E, T2], E, A1, T1, A2, T2 any](f1 F1, f2 F2) func(T.Tuple2[A1, A2]) Either[E, T.Tuple2[T1, T2]] {
	return func(t T.Tuple2[A1, A2]) Either[E, T.Tuple2[T1, T2]] {
		return A.TraverseTuple2(
			Map[E, T1, func(T2) T.Tuple2[T1, T2]],
			Ap[T.Tuple2[T1, T2], E, T2],
			f1,
			f2,
			t,
		)
	}
}

// Eitherize3 converts a function with 3 parameters returning a tuple into a function with 3 parameters returning an Either
// The inverse function is [Uneitherize3]
func Eitherize3[F ~func(T0, T1, T2) (R, error), T0, T1, T2, R any](f F) func(T0, T1, T2) Either[error, R] {
	return func(t0 T0, t1 T1, t2 T2) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0, t1, t2)
		})
	}
}

// Uneitherize3 converts a function with 3 parameters returning an Either into a function with 3 parameters returning a tuple
// The inverse function is [Eitherize3]
func Uneitherize3[F ~func(T0, T1, T2) Either[error, R], T0, T1, T2, R any](f F) func(T0, T1, T2) (R, error) {
	return func(t0 T0, t1 T1, t2 T2) (R, error) {
		return UnwrapError(f(t0, t1, t2))
	}
}

// SequenceT3 converts 3 parameters of [Either[E, T]] into a [Either[E, Tuple3]].
func SequenceT3[E, T1, T2, T3 any](t1 Either[E, T1], t2 Either[E, T2], t3 Either[E, T3]) Either[E, T.Tuple3[T1, T2, T3]] {
	return A.SequenceT3(
		Map[E, T1, func(T2) func(T3) T.Tuple3[T1, T2, T3]],
		Ap[func(T3) T.Tuple3[T1, T2, T3], E, T2],
		Ap[T.Tuple3[T1, T2, T3], E, T3],
		t1,
		t2,
		t3,
	)
}

// SequenceTuple3 converts a [Tuple3] of [Either[E, T]] into an [Either[E, Tuple3]].
func SequenceTuple3[E, T1, T2, T3 any](t T.Tuple3[Either[E, T1], Either[E, T2], Either[E, T3]]) Either[E, T.Tuple3[T1, T2, T3]] {
	return A.SequenceTuple3(
		Map[E, T1, func(T2) func(T3) T.Tuple3[T1, T2, T3]],
		Ap[func(T3) T.Tuple3[T1, T2, T3], E, T2],
		Ap[T.Tuple3[T1, T2, T3], E, T3],
		t,
	)
}

// TraverseTuple3 converts a [Tuple3] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple3]].
func TraverseTuple3[F1 ~func(A1) Either[E, T1], F2 ~func(A2) Either[E, T2], F3 ~func(A3) Either[E, T3], E, A1, T1, A2, T2, A3, T3 any](f1 F1, f2 F2, f3 F3) func(T.Tuple3[A1, A2, A3]) Either[E, T.Tuple3[T1, T2, T3]] {
	return func(t T.Tuple3[A1, A2, A3]) Either[E, T.Tuple3[T1, T2, T3]] {
		return A.TraverseTuple3(
			Map[E, T1, func(T2) func(T3) T.Tuple3[T1, T2, T3]],
			Ap[func(T3) T.Tuple3[T1, T2, T3], E, T2],
			Ap[T.Tuple3[T1, T2, T3], E, T3],
			f1,
			f2,
			f3,
			t,
		)
	}
}

// Eitherize4 converts a function with 4 parameters returning a tuple into a function with 4 parameters returning an Either
// The inverse function is [Uneitherize4]
func Eitherize4[F ~func(T0, T1, T2, T3) (R, error), T0, T1, T2, T3, R any](f F) func(T0, T1, T2, T3) Either[error, R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0, t1, t2, t3)
		})
	}
}

// Uneitherize4 converts a function with 4 parameters returning an Either into a function with 4 parameters returning a tuple
// The inverse function is [Eitherize4]
func Uneitherize4[F ~func(T0, T1, T2, T3) Either[error, R], T0, T1, T2, T3, R any](f F) func(T0, T1, T2, T3) (R, error) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3) (R, error) {
		return UnwrapError(f(t0, t1, t2, t3))
	}
}

// SequenceT4 converts 4 parameters of [Either[E, T]] into a [Either[E, Tuple4]].
func SequenceT4[E, T1, T2, T3, T4 any](t1 Either[E, T1], t2 Either[E, T2], t3 Either[E, T3], t4 Either[E, T4]) Either[E, T.Tuple4[T1, T2, T3, T4]] {
	return A.SequenceT4(
		Map[E, T1, func(T2) func(T3) func(T4) T.Tuple4[T1, T2, T3, T4]],
		Ap[func(T3) func(T4) T.Tuple4[T1, T2, T3, T4], E, T2],
		Ap[func(T4) T.Tuple4[T1, T2, T3, T4], E, T3],
		Ap[T.Tuple4[T1, T2, T3, T4], E, T4],
		t1,
		t2,
		t3,
		t4,
	)
}

// SequenceTuple4 converts a [Tuple4] of [Either[E, T]] into an [Either[E, Tuple4]].
func SequenceTuple4[E, T1, T2, T3, T4 any](t T.Tuple4[Either[E, T1], Either[E, T2], Either[E, T3], Either[E, T4]]) Either[E, T.Tuple4[T1, T2, T3, T4]] {
	return A.SequenceTuple4(
		Map[E, T1, func(T2) func(T3) func(T4) T.Tuple4[T1, T2, T3, T4]],
		Ap[func(T3) func(T4) T.Tuple4[T1, T2, T3, T4], E, T2],
		Ap[func(T4) T.Tuple4[T1, T2, T3, T4], E, T3],
		Ap[T.Tuple4[T1, T2, T3, T4], E, T4],
		t,
	)
}

// TraverseTuple4 converts a [Tuple4] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple4]].
func TraverseTuple4[F1 ~func(A1) Either[E, T1], F2 ~func(A2) Either[E, T2], F3 ~func(A3) Either[E, T3], F4 ~func(A4) Either[E, T4], E, A1, T1, A2, T2, A3, T3, A4, T4 any](f1 F1, f2 F2, f3 F3, f4 F4) func(T.Tuple4[A1, A2, A3, A4]) Either[E, T.Tuple4[T1, T2, T3, T4]] {
	return func(t T.Tuple4[A1, A2, A3, A4]) Either[E, T.Tuple4[T1, T2, T3, T4]] {
		return A.TraverseTuple4(
			Map[E, T1, func(T2) func(T3) func(T4) T.Tuple4[T1, T2, T3, T4]],
			Ap[func(T3) func(T4) T.Tuple4[T1, T2, T3, T4], E, T2],
			Ap[func(T4) T.Tuple4[T1, T2, T3, T4], E, T3],
			Ap[T.Tuple4[T1, T2, T3, T4], E, T4],
			f1,
			f2,
			f3,
			f4,
			t,
		)
	}
}

// Eitherize5 converts a function with 5 parameters returning a tuple into a function with 5 parameters returning an Either
// The inverse function is [Uneitherize5]
func Eitherize5[F ~func(T0, T1, T2, T3, T4) (R, error), T0, T1, T2, T3, T4, R any](f F) func(T0, T1, T2, T3, T4) Either[error, R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0, t1, t2, t3, t4)
		})
	}
}

// Uneitherize5 converts a function with 5 parameters returning an Either into a function with 5 parameters returning a tuple
// The inverse function is [Eitherize5]
func Uneitherize5[F ~func(T0, T1, T2, T3, T4) Either[error, R], T0, T1, T2, T3, T4, R any](f F) func(T0, T1, T2, T3, T4) (R, error) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4) (R, error) {
		return UnwrapError(f(t0, t1, t2, t3, t4))
	}
}

// SequenceT5 converts 5 parameters of [Either[E, T]] into a [Either[E, Tuple5]].
func SequenceT5[E, T1, T2, T3, T4, T5 any](t1 Either[E, T1], t2 Either[E, T2], t3 Either[E, T3], t4 Either[E, T4], t5 Either[E, T5]) Either[E, T.Tuple5[T1, T2, T3, T4, T5]] {
	return A.SequenceT5(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5]],
		Ap[func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], E, T2],
		Ap[func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], E, T3],
		Ap[func(T5) T.Tuple5[T1, T2, T3, T4, T5], E, T4],
		Ap[T.Tuple5[T1, T2, T3, T4, T5], E, T5],
		t1,
		t2,
		t3,
		t4,
		t5,
	)
}

// SequenceTuple5 converts a [Tuple5] of [Either[E, T]] into an [Either[E, Tuple5]].
func SequenceTuple5[E, T1, T2, T3, T4, T5 any](t T.Tuple5[Either[E, T1], Either[E, T2], Either[E, T3], Either[E, T4], Either[E, T5]]) Either[E, T.Tuple5[T1, T2, T3, T4, T5]] {
	return A.SequenceTuple5(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5]],
		Ap[func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], E, T2],
		Ap[func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], E, T3],
		Ap[func(T5) T.Tuple5[T1, T2, T3, T4, T5], E, T4],
		Ap[T.Tuple5[T1, T2, T3, T4, T5], E, T5],
		t,
	)
}

// TraverseTuple5 converts a [Tuple5] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple5]].
func TraverseTuple5[F1 ~func(A1) Either[E, T1], F2 ~func(A2) Either[E, T2], F3 ~func(A3) Either[E, T3], F4 ~func(A4) Either[E, T4], F5 ~func(A5) Either[E, T5], E, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5) func(T.Tuple5[A1, A2, A3, A4, A5]) Either[E, T.Tuple5[T1, T2, T3, T4, T5]] {
	return func(t T.Tuple5[A1, A2, A3, A4, A5]) Either[E, T.Tuple5[T1, T2, T3, T4, T5]] {
		return A.TraverseTuple5(
			Map[E, T1, func(T2) func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5]],
			Ap[func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], E, T2],
			Ap[func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], E, T3],
			Ap[func(T5) T.Tuple5[T1, T2, T3, T4, T5], E, T4],
			Ap[T.Tuple5[T1, T2, T3, T4, T5], E, T5],
			f1,
			f2,
			f3,
			f4,
			f5,
			t,
		)
	}
}

// Eitherize6 converts a function with 6 parameters returning a tuple into a function with 6 parameters returning an Either
// The inverse function is [Uneitherize6]
func Eitherize6[F ~func(T0, T1, T2, T3, T4, T5) (R, error), T0, T1, T2, T3, T4, T5, R any](f F) func(T0, T1, T2, T3, T4, T5) Either[error, R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0, t1, t2, t3, t4, t5)
		})
	}
}

// Uneitherize6 converts a function with 6 parameters returning an Either into a function with 6 parameters returning a tuple
// The inverse function is [Eitherize6]
func Uneitherize6[F ~func(T0, T1, T2, T3, T4, T5) Either[error, R], T0, T1, T2, T3, T4, T5, R any](f F) func(T0, T1, T2, T3, T4, T5) (R, error) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) (R, error) {
		return UnwrapError(f(t0, t1, t2, t3, t4, t5))
	}
}

// SequenceT6 converts 6 parameters of [Either[E, T]] into a [Either[E, Tuple6]].
func SequenceT6[E, T1, T2, T3, T4, T5, T6 any](t1 Either[E, T1], t2 Either[E, T2], t3 Either[E, T3], t4 Either[E, T4], t5 Either[E, T5], t6 Either[E, T6]) Either[E, T.Tuple6[T1, T2, T3, T4, T5, T6]] {
	return A.SequenceT6(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6]],
		Ap[func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T2],
		Ap[func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T3],
		Ap[func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T4],
		Ap[func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T5],
		Ap[T.Tuple6[T1, T2, T3, T4, T5, T6], E, T6],
		t1,
		t2,
		t3,
		t4,
		t5,
		t6,
	)
}

// SequenceTuple6 converts a [Tuple6] of [Either[E, T]] into an [Either[E, Tuple6]].
func SequenceTuple6[E, T1, T2, T3, T4, T5, T6 any](t T.Tuple6[Either[E, T1], Either[E, T2], Either[E, T3], Either[E, T4], Either[E, T5], Either[E, T6]]) Either[E, T.Tuple6[T1, T2, T3, T4, T5, T6]] {
	return A.SequenceTuple6(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6]],
		Ap[func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T2],
		Ap[func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T3],
		Ap[func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T4],
		Ap[func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T5],
		Ap[T.Tuple6[T1, T2, T3, T4, T5, T6], E, T6],
		t,
	)
}

// TraverseTuple6 converts a [Tuple6] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple6]].
func TraverseTuple6[F1 ~func(A1) Either[E, T1], F2 ~func(A2) Either[E, T2], F3 ~func(A3) Either[E, T3], F4 ~func(A4) Either[E, T4], F5 ~func(A5) Either[E, T5], F6 ~func(A6) Either[E, T6], E, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6) func(T.Tuple6[A1, A2, A3, A4, A5, A6]) Either[E, T.Tuple6[T1, T2, T3, T4, T5, T6]] {
	return func(t T.Tuple6[A1, A2, A3, A4, A5, A6]) Either[E, T.Tuple6[T1, T2, T3, T4, T5, T6]] {
		return A.TraverseTuple6(
			Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6]],
			Ap[func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T2],
			Ap[func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T3],
			Ap[func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T4],
			Ap[func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], E, T5],
			Ap[T.Tuple6[T1, T2, T3, T4, T5, T6], E, T6],
			f1,
			f2,
			f3,
			f4,
			f5,
			f6,
			t,
		)
	}
}

// Eitherize7 converts a function with 7 parameters returning a tuple into a function with 7 parameters returning an Either
// The inverse function is [Uneitherize7]
func Eitherize7[F ~func(T0, T1, T2, T3, T4, T5, T6) (R, error), T0, T1, T2, T3, T4, T5, T6, R any](f F) func(T0, T1, T2, T3, T4, T5, T6) Either[error, R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0, t1, t2, t3, t4, t5, t6)
		})
	}
}

// Uneitherize7 converts a function with 7 parameters returning an Either into a function with 7 parameters returning a tuple
// The inverse function is [Eitherize7]
func Uneitherize7[F ~func(T0, T1, T2, T3, T4, T5, T6) Either[error, R], T0, T1, T2, T3, T4, T5, T6, R any](f F) func(T0, T1, T2, T3, T4, T5, T6) (R, error) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) (R, error) {
		return UnwrapError(f(t0, t1, t2, t3, t4, t5, t6))
	}
}

// SequenceT7 converts 7 parameters of [Either[E, T]] into a [Either[E, Tuple7]].
func SequenceT7[E, T1, T2, T3, T4, T5, T6, T7 any](t1 Either[E, T1], t2 Either[E, T2], t3 Either[E, T3], t4 Either[E, T4], t5 Either[E, T5], t6 Either[E, T6], t7 Either[E, T7]) Either[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	return A.SequenceT7(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
		Ap[func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T2],
		Ap[func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T3],
		Ap[func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T4],
		Ap[func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T5],
		Ap[func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T6],
		Ap[T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T7],
		t1,
		t2,
		t3,
		t4,
		t5,
		t6,
		t7,
	)
}

// SequenceTuple7 converts a [Tuple7] of [Either[E, T]] into an [Either[E, Tuple7]].
func SequenceTuple7[E, T1, T2, T3, T4, T5, T6, T7 any](t T.Tuple7[Either[E, T1], Either[E, T2], Either[E, T3], Either[E, T4], Either[E, T5], Either[E, T6], Either[E, T7]]) Either[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	return A.SequenceTuple7(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
		Ap[func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T2],
		Ap[func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T3],
		Ap[func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T4],
		Ap[func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T5],
		Ap[func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T6],
		Ap[T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T7],
		t,
	)
}

// TraverseTuple7 converts a [Tuple7] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple7]].
func TraverseTuple7[F1 ~func(A1) Either[E, T1], F2 ~func(A2) Either[E, T2], F3 ~func(A3) Either[E, T3], F4 ~func(A4) Either[E, T4], F5 ~func(A5) Either[E, T5], F6 ~func(A6) Either[E, T6], F7 ~func(A7) Either[E, T7], E, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6, A7, T7 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7) func(T.Tuple7[A1, A2, A3, A4, A5, A6, A7]) Either[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	return func(t T.Tuple7[A1, A2, A3, A4, A5, A6, A7]) Either[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
		return A.TraverseTuple7(
			Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
			Ap[func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T2],
			Ap[func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T3],
			Ap[func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T4],
			Ap[func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T5],
			Ap[func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T6],
			Ap[T.Tuple7[T1, T2, T3, T4, T5, T6, T7], E, T7],
			f1,
			f2,
			f3,
			f4,
			f5,
			f6,
			f7,
			t,
		)
	}
}

// Eitherize8 converts a function with 8 parameters returning a tuple into a function with 8 parameters returning an Either
// The inverse function is [Uneitherize8]
func Eitherize8[F ~func(T0, T1, T2, T3, T4, T5, T6, T7) (R, error), T0, T1, T2, T3, T4, T5, T6, T7, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7) Either[error, R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0, t1, t2, t3, t4, t5, t6, t7)
		})
	}
}

// Uneitherize8 converts a function with 8 parameters returning an Either into a function with 8 parameters returning a tuple
// The inverse function is [Eitherize8]
func Uneitherize8[F ~func(T0, T1, T2, T3, T4, T5, T6, T7) Either[error, R], T0, T1, T2, T3, T4, T5, T6, T7, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7) (R, error) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) (R, error) {
		return UnwrapError(f(t0, t1, t2, t3, t4, t5, t6, t7))
	}
}

// SequenceT8 converts 8 parameters of [Either[E, T]] into a [Either[E, Tuple8]].
func SequenceT8[E, T1, T2, T3, T4, T5, T6, T7, T8 any](t1 Either[E, T1], t2 Either[E, T2], t3 Either[E, T3], t4 Either[E, T4], t5 Either[E, T5], t6 Either[E, T6], t7 Either[E, T7], t8 Either[E, T8]) Either[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	return A.SequenceT8(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
		Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T2],
		Ap[func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T3],
		Ap[func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T4],
		Ap[func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T5],
		Ap[func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T6],
		Ap[func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T7],
		Ap[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T8],
		t1,
		t2,
		t3,
		t4,
		t5,
		t6,
		t7,
		t8,
	)
}

// SequenceTuple8 converts a [Tuple8] of [Either[E, T]] into an [Either[E, Tuple8]].
func SequenceTuple8[E, T1, T2, T3, T4, T5, T6, T7, T8 any](t T.Tuple8[Either[E, T1], Either[E, T2], Either[E, T3], Either[E, T4], Either[E, T5], Either[E, T6], Either[E, T7], Either[E, T8]]) Either[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	return A.SequenceTuple8(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
		Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T2],
		Ap[func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T3],
		Ap[func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T4],
		Ap[func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T5],
		Ap[func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T6],
		Ap[func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T7],
		Ap[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T8],
		t,
	)
}

// TraverseTuple8 converts a [Tuple8] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple8]].
func TraverseTuple8[F1 ~func(A1) Either[E, T1], F2 ~func(A2) Either[E, T2], F3 ~func(A3) Either[E, T3], F4 ~func(A4) Either[E, T4], F5 ~func(A5) Either[E, T5], F6 ~func(A6) Either[E, T6], F7 ~func(A7) Either[E, T7], F8 ~func(A8) Either[E, T8], E, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6, A7, T7, A8, T8 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8) func(T.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) Either[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	return func(t T.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) Either[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
		return A.TraverseTuple8(
			Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
			Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T2],
			Ap[func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T3],
			Ap[func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T4],
			Ap[func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T5],
			Ap[func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T6],
			Ap[func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T7],
			Ap[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], E, T8],
			f1,
			f2,
			f3,
			f4,
			f5,
			f6,
			f7,
			f8,
			t,
		)
	}
}

// Eitherize9 converts a function with 9 parameters returning a tuple into a function with 9 parameters returning an Either
// The inverse function is [Uneitherize9]
func Eitherize9[F ~func(T0, T1, T2, T3, T4, T5, T6, T7, T8) (R, error), T0, T1, T2, T3, T4, T5, T6, T7, T8, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8) Either[error, R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0, t1, t2, t3, t4, t5, t6, t7, t8)
		})
	}
}

// Uneitherize9 converts a function with 9 parameters returning an Either into a function with 9 parameters returning a tuple
// The inverse function is [Eitherize9]
func Uneitherize9[F ~func(T0, T1, T2, T3, T4, T5, T6, T7, T8) Either[error, R], T0, T1, T2, T3, T4, T5, T6, T7, T8, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8) (R, error) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) (R, error) {
		return UnwrapError(f(t0, t1, t2, t3, t4, t5, t6, t7, t8))
	}
}

// SequenceT9 converts 9 parameters of [Either[E, T]] into a [Either[E, Tuple9]].
func SequenceT9[E, T1, T2, T3, T4, T5, T6, T7, T8, T9 any](t1 Either[E, T1], t2 Either[E, T2], t3 Either[E, T3], t4 Either[E, T4], t5 Either[E, T5], t6 Either[E, T6], t7 Either[E, T7], t8 Either[E, T8], t9 Either[E, T9]) Either[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	return A.SequenceT9(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
		Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T2],
		Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T3],
		Ap[func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T4],
		Ap[func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T5],
		Ap[func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T6],
		Ap[func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T7],
		Ap[func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T8],
		Ap[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T9],
		t1,
		t2,
		t3,
		t4,
		t5,
		t6,
		t7,
		t8,
		t9,
	)
}

// SequenceTuple9 converts a [Tuple9] of [Either[E, T]] into an [Either[E, Tuple9]].
func SequenceTuple9[E, T1, T2, T3, T4, T5, T6, T7, T8, T9 any](t T.Tuple9[Either[E, T1], Either[E, T2], Either[E, T3], Either[E, T4], Either[E, T5], Either[E, T6], Either[E, T7], Either[E, T8], Either[E, T9]]) Either[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	return A.SequenceTuple9(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
		Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T2],
		Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T3],
		Ap[func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T4],
		Ap[func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T5],
		Ap[func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T6],
		Ap[func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T7],
		Ap[func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T8],
		Ap[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T9],
		t,
	)
}

// TraverseTuple9 converts a [Tuple9] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple9]].
func TraverseTuple9[F1 ~func(A1) Either[E, T1], F2 ~func(A2) Either[E, T2], F3 ~func(A3) Either[E, T3], F4 ~func(A4) Either[E, T4], F5 ~func(A5) Either[E, T5], F6 ~func(A6) Either[E, T6], F7 ~func(A7) Either[E, T7], F8 ~func(A8) Either[E, T8], F9 ~func(A9) Either[E, T9], E, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6, A7, T7, A8, T8, A9, T9 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9) func(T.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) Either[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	return func(t T.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) Either[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
		return A.TraverseTuple9(
			Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
			Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T2],
			Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T3],
			Ap[func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T4],
			Ap[func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T5],
			Ap[func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T6],
			Ap[func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T7],
			Ap[func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T8],
			Ap[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], E, T9],
			f1,
			f2,
			f3,
			f4,
			f5,
			f6,
			f7,
			f8,
			f9,
			t,
		)
	}
}

// Eitherize10 converts a function with 10 parameters returning a tuple into a function with 10 parameters returning an Either
// The inverse function is [Uneitherize10]
func Eitherize10[F ~func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) (R, error), T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) Either[error, R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) Either[error, R] {
		return TryCatchError(func() (R, error) {
			return f(t0, t1, t2, t3, t4, t5, t6, t7, t8, t9)
		})
	}
}

// Uneitherize10 converts a function with 10 parameters returning an Either into a function with 10 parameters returning a tuple
// The inverse function is [Eitherize10]
func Uneitherize10[F ~func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) Either[error, R], T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) (R, error) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) (R, error) {
		return UnwrapError(f(t0, t1, t2, t3, t4, t5, t6, t7, t8, t9))
	}
}

// SequenceT10 converts 10 parameters of [Either[E, T]] into a [Either[E, Tuple10]].
func SequenceT10[E, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](t1 Either[E, T1], t2 Either[E, T2], t3 Either[E, T3], t4 Either[E, T4], t5 Either[E, T5], t6 Either[E, T6], t7 Either[E, T7], t8 Either[E, T8], t9 Either[E, T9], t10 Either[E, T10]) Either[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	return A.SequenceT10(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
		Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T2],
		Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T3],
		Ap[func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T4],
		Ap[func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T5],
		Ap[func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T6],
		Ap[func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T7],
		Ap[func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T8],
		Ap[func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T9],
		Ap[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T10],
		t1,
		t2,
		t3,
		t4,
		t5,
		t6,
		t7,
		t8,
		t9,
		t10,
	)
}

// SequenceTuple10 converts a [Tuple10] of [Either[E, T]] into an [Either[E, Tuple10]].
func SequenceTuple10[E, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](t T.Tuple10[Either[E, T1], Either[E, T2], Either[E, T3], Either[E, T4], Either[E, T5], Either[E, T6], Either[E, T7], Either[E, T8], Either[E, T9], Either[E, T10]]) Either[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	return A.SequenceTuple10(
		Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
		Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T2],
		Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T3],
		Ap[func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T4],
		Ap[func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T5],
		Ap[func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T6],
		Ap[func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T7],
		Ap[func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T8],
		Ap[func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T9],
		Ap[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T10],
		t,
	)
}

// TraverseTuple10 converts a [Tuple10] of [A] via transformation functions transforming [A] to [Either[E, A]] into a [Either[E, Tuple10]].
func TraverseTuple10[F1 ~func(A1) Either[E, T1], F2 ~func(A2) Either[E, T2], F3 ~func(A3) Either[E, T3], F4 ~func(A4) Either[E, T4], F5 ~func(A5) Either[E, T5], F6 ~func(A6) Either[E, T6], F7 ~func(A7) Either[E, T7], F8 ~func(A8) Either[E, T8], F9 ~func(A9) Either[E, T9], F10 ~func(A10) Either[E, T10], E, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6, A7, T7, A8, T8, A9, T9, A10, T10 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9, f10 F10) func(T.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) Either[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	return func(t T.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) Either[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
		return A.TraverseTuple10(
			Map[E, T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
			Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T2],
			Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T3],
			Ap[func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T4],
			Ap[func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T5],
			Ap[func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T6],
			Ap[func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T7],
			Ap[func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T8],
			Ap[func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T9],
			Ap[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], E, T10],
			f1,
			f2,
			f3,
			f4,
			f5,
			f6,
			f7,
			f8,
			f9,
			f10,
			t,
		)
	}
}
