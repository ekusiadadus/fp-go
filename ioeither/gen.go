// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-08-11 11:38:09.4187123 +0200 CEST m=+0.099620101

package ioeither

import (
	G "github.com/IBM/fp-go/ioeither/generic"
	T "github.com/IBM/fp-go/tuple"
)

// Eitherize0 converts a function with 1 parameters returning a tuple into a function with 0 parameters returning a [IOEither[error, R]]
func Eitherize0[F ~func() (R, error), R any](f F) func() IOEither[error, R] {
	return G.Eitherize0[IOEither[error, R]](f)
}

// Uneitherize0 converts a function with 1 parameters returning a tuple into a function with 0 parameters returning a [IOEither[error, R]]
func Uneitherize0[F ~func() IOEither[error, R], R any](f F) func() (R, error) {
	return G.Uneitherize0[IOEither[error, R]](f)
}

// Eitherize1 converts a function with 2 parameters returning a tuple into a function with 1 parameters returning a [IOEither[error, R]]
func Eitherize1[F ~func(T1) (R, error), T1, R any](f F) func(T1) IOEither[error, R] {
	return G.Eitherize1[IOEither[error, R]](f)
}

// Uneitherize1 converts a function with 2 parameters returning a tuple into a function with 1 parameters returning a [IOEither[error, R]]
func Uneitherize1[F ~func(T1) IOEither[error, R], T1, R any](f F) func(T1) (R, error) {
	return G.Uneitherize1[IOEither[error, R]](f)
}

// SequenceT1 converts 1 [IOEither[E, T]] into a [IOEither[E, T.Tuple1[T1]]]
func SequenceT1[E, T1 any](
	t1 IOEither[E, T1],
) IOEither[E, T.Tuple1[T1]] {
	return G.SequenceT1[
		IOEither[E, T.Tuple1[T1]],
		IOEither[E, T1],
	](t1)
}

// SequenceTuple1 converts a [T.Tuple1[IOEither[E, T]]] into a [IOEither[E, T.Tuple1[T1]]]
func SequenceTuple1[E, T1 any](t T.Tuple1[IOEither[E, T1]]) IOEither[E, T.Tuple1[T1]] {
	return G.SequenceTuple1[
		IOEither[E, T.Tuple1[T1]],
		IOEither[E, T1],
	](t)
}

// TraverseTuple1 converts a [T.Tuple1[IOEither[E, T]]] into a [IOEither[E, T.Tuple1[T1]]]
func TraverseTuple1[F1 ~func(A1) IOEither[E, T1], E, A1, T1 any](f1 F1) func(T.Tuple1[A1]) IOEither[E, T.Tuple1[T1]] {
	return G.TraverseTuple1[IOEither[E, T.Tuple1[T1]]](f1)
}

// Eitherize2 converts a function with 3 parameters returning a tuple into a function with 2 parameters returning a [IOEither[error, R]]
func Eitherize2[F ~func(T1, T2) (R, error), T1, T2, R any](f F) func(T1, T2) IOEither[error, R] {
	return G.Eitherize2[IOEither[error, R]](f)
}

// Uneitherize2 converts a function with 3 parameters returning a tuple into a function with 2 parameters returning a [IOEither[error, R]]
func Uneitherize2[F ~func(T1, T2) IOEither[error, R], T1, T2, R any](f F) func(T1, T2) (R, error) {
	return G.Uneitherize2[IOEither[error, R]](f)
}

// SequenceT2 converts 2 [IOEither[E, T]] into a [IOEither[E, T.Tuple2[T1, T2]]]
func SequenceT2[E, T1, T2 any](
	t1 IOEither[E, T1],
	t2 IOEither[E, T2],
) IOEither[E, T.Tuple2[T1, T2]] {
	return G.SequenceT2[
		IOEither[E, T.Tuple2[T1, T2]],
		IOEither[E, T1],
		IOEither[E, T2],
	](t1, t2)
}

// SequenceTuple2 converts a [T.Tuple2[IOEither[E, T]]] into a [IOEither[E, T.Tuple2[T1, T2]]]
func SequenceTuple2[E, T1, T2 any](t T.Tuple2[IOEither[E, T1], IOEither[E, T2]]) IOEither[E, T.Tuple2[T1, T2]] {
	return G.SequenceTuple2[
		IOEither[E, T.Tuple2[T1, T2]],
		IOEither[E, T1],
		IOEither[E, T2],
	](t)
}

// TraverseTuple2 converts a [T.Tuple2[IOEither[E, T]]] into a [IOEither[E, T.Tuple2[T1, T2]]]
func TraverseTuple2[F1 ~func(A1) IOEither[E, T1], F2 ~func(A2) IOEither[E, T2], E, A1, A2, T1, T2 any](f1 F1, f2 F2) func(T.Tuple2[A1, A2]) IOEither[E, T.Tuple2[T1, T2]] {
	return G.TraverseTuple2[IOEither[E, T.Tuple2[T1, T2]]](f1, f2)
}

// Eitherize3 converts a function with 4 parameters returning a tuple into a function with 3 parameters returning a [IOEither[error, R]]
func Eitherize3[F ~func(T1, T2, T3) (R, error), T1, T2, T3, R any](f F) func(T1, T2, T3) IOEither[error, R] {
	return G.Eitherize3[IOEither[error, R]](f)
}

// Uneitherize3 converts a function with 4 parameters returning a tuple into a function with 3 parameters returning a [IOEither[error, R]]
func Uneitherize3[F ~func(T1, T2, T3) IOEither[error, R], T1, T2, T3, R any](f F) func(T1, T2, T3) (R, error) {
	return G.Uneitherize3[IOEither[error, R]](f)
}

// SequenceT3 converts 3 [IOEither[E, T]] into a [IOEither[E, T.Tuple3[T1, T2, T3]]]
func SequenceT3[E, T1, T2, T3 any](
	t1 IOEither[E, T1],
	t2 IOEither[E, T2],
	t3 IOEither[E, T3],
) IOEither[E, T.Tuple3[T1, T2, T3]] {
	return G.SequenceT3[
		IOEither[E, T.Tuple3[T1, T2, T3]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
	](t1, t2, t3)
}

// SequenceTuple3 converts a [T.Tuple3[IOEither[E, T]]] into a [IOEither[E, T.Tuple3[T1, T2, T3]]]
func SequenceTuple3[E, T1, T2, T3 any](t T.Tuple3[IOEither[E, T1], IOEither[E, T2], IOEither[E, T3]]) IOEither[E, T.Tuple3[T1, T2, T3]] {
	return G.SequenceTuple3[
		IOEither[E, T.Tuple3[T1, T2, T3]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
	](t)
}

// TraverseTuple3 converts a [T.Tuple3[IOEither[E, T]]] into a [IOEither[E, T.Tuple3[T1, T2, T3]]]
func TraverseTuple3[F1 ~func(A1) IOEither[E, T1], F2 ~func(A2) IOEither[E, T2], F3 ~func(A3) IOEither[E, T3], E, A1, A2, A3, T1, T2, T3 any](f1 F1, f2 F2, f3 F3) func(T.Tuple3[A1, A2, A3]) IOEither[E, T.Tuple3[T1, T2, T3]] {
	return G.TraverseTuple3[IOEither[E, T.Tuple3[T1, T2, T3]]](f1, f2, f3)
}

// Eitherize4 converts a function with 5 parameters returning a tuple into a function with 4 parameters returning a [IOEither[error, R]]
func Eitherize4[F ~func(T1, T2, T3, T4) (R, error), T1, T2, T3, T4, R any](f F) func(T1, T2, T3, T4) IOEither[error, R] {
	return G.Eitherize4[IOEither[error, R]](f)
}

// Uneitherize4 converts a function with 5 parameters returning a tuple into a function with 4 parameters returning a [IOEither[error, R]]
func Uneitherize4[F ~func(T1, T2, T3, T4) IOEither[error, R], T1, T2, T3, T4, R any](f F) func(T1, T2, T3, T4) (R, error) {
	return G.Uneitherize4[IOEither[error, R]](f)
}

// SequenceT4 converts 4 [IOEither[E, T]] into a [IOEither[E, T.Tuple4[T1, T2, T3, T4]]]
func SequenceT4[E, T1, T2, T3, T4 any](
	t1 IOEither[E, T1],
	t2 IOEither[E, T2],
	t3 IOEither[E, T3],
	t4 IOEither[E, T4],
) IOEither[E, T.Tuple4[T1, T2, T3, T4]] {
	return G.SequenceT4[
		IOEither[E, T.Tuple4[T1, T2, T3, T4]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
	](t1, t2, t3, t4)
}

// SequenceTuple4 converts a [T.Tuple4[IOEither[E, T]]] into a [IOEither[E, T.Tuple4[T1, T2, T3, T4]]]
func SequenceTuple4[E, T1, T2, T3, T4 any](t T.Tuple4[IOEither[E, T1], IOEither[E, T2], IOEither[E, T3], IOEither[E, T4]]) IOEither[E, T.Tuple4[T1, T2, T3, T4]] {
	return G.SequenceTuple4[
		IOEither[E, T.Tuple4[T1, T2, T3, T4]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
	](t)
}

// TraverseTuple4 converts a [T.Tuple4[IOEither[E, T]]] into a [IOEither[E, T.Tuple4[T1, T2, T3, T4]]]
func TraverseTuple4[F1 ~func(A1) IOEither[E, T1], F2 ~func(A2) IOEither[E, T2], F3 ~func(A3) IOEither[E, T3], F4 ~func(A4) IOEither[E, T4], E, A1, A2, A3, A4, T1, T2, T3, T4 any](f1 F1, f2 F2, f3 F3, f4 F4) func(T.Tuple4[A1, A2, A3, A4]) IOEither[E, T.Tuple4[T1, T2, T3, T4]] {
	return G.TraverseTuple4[IOEither[E, T.Tuple4[T1, T2, T3, T4]]](f1, f2, f3, f4)
}

// Eitherize5 converts a function with 6 parameters returning a tuple into a function with 5 parameters returning a [IOEither[error, R]]
func Eitherize5[F ~func(T1, T2, T3, T4, T5) (R, error), T1, T2, T3, T4, T5, R any](f F) func(T1, T2, T3, T4, T5) IOEither[error, R] {
	return G.Eitherize5[IOEither[error, R]](f)
}

// Uneitherize5 converts a function with 6 parameters returning a tuple into a function with 5 parameters returning a [IOEither[error, R]]
func Uneitherize5[F ~func(T1, T2, T3, T4, T5) IOEither[error, R], T1, T2, T3, T4, T5, R any](f F) func(T1, T2, T3, T4, T5) (R, error) {
	return G.Uneitherize5[IOEither[error, R]](f)
}

// SequenceT5 converts 5 [IOEither[E, T]] into a [IOEither[E, T.Tuple5[T1, T2, T3, T4, T5]]]
func SequenceT5[E, T1, T2, T3, T4, T5 any](
	t1 IOEither[E, T1],
	t2 IOEither[E, T2],
	t3 IOEither[E, T3],
	t4 IOEither[E, T4],
	t5 IOEither[E, T5],
) IOEither[E, T.Tuple5[T1, T2, T3, T4, T5]] {
	return G.SequenceT5[
		IOEither[E, T.Tuple5[T1, T2, T3, T4, T5]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
	](t1, t2, t3, t4, t5)
}

// SequenceTuple5 converts a [T.Tuple5[IOEither[E, T]]] into a [IOEither[E, T.Tuple5[T1, T2, T3, T4, T5]]]
func SequenceTuple5[E, T1, T2, T3, T4, T5 any](t T.Tuple5[IOEither[E, T1], IOEither[E, T2], IOEither[E, T3], IOEither[E, T4], IOEither[E, T5]]) IOEither[E, T.Tuple5[T1, T2, T3, T4, T5]] {
	return G.SequenceTuple5[
		IOEither[E, T.Tuple5[T1, T2, T3, T4, T5]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
	](t)
}

// TraverseTuple5 converts a [T.Tuple5[IOEither[E, T]]] into a [IOEither[E, T.Tuple5[T1, T2, T3, T4, T5]]]
func TraverseTuple5[F1 ~func(A1) IOEither[E, T1], F2 ~func(A2) IOEither[E, T2], F3 ~func(A3) IOEither[E, T3], F4 ~func(A4) IOEither[E, T4], F5 ~func(A5) IOEither[E, T5], E, A1, A2, A3, A4, A5, T1, T2, T3, T4, T5 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5) func(T.Tuple5[A1, A2, A3, A4, A5]) IOEither[E, T.Tuple5[T1, T2, T3, T4, T5]] {
	return G.TraverseTuple5[IOEither[E, T.Tuple5[T1, T2, T3, T4, T5]]](f1, f2, f3, f4, f5)
}

// Eitherize6 converts a function with 7 parameters returning a tuple into a function with 6 parameters returning a [IOEither[error, R]]
func Eitherize6[F ~func(T1, T2, T3, T4, T5, T6) (R, error), T1, T2, T3, T4, T5, T6, R any](f F) func(T1, T2, T3, T4, T5, T6) IOEither[error, R] {
	return G.Eitherize6[IOEither[error, R]](f)
}

// Uneitherize6 converts a function with 7 parameters returning a tuple into a function with 6 parameters returning a [IOEither[error, R]]
func Uneitherize6[F ~func(T1, T2, T3, T4, T5, T6) IOEither[error, R], T1, T2, T3, T4, T5, T6, R any](f F) func(T1, T2, T3, T4, T5, T6) (R, error) {
	return G.Uneitherize6[IOEither[error, R]](f)
}

// SequenceT6 converts 6 [IOEither[E, T]] into a [IOEither[E, T.Tuple6[T1, T2, T3, T4, T5, T6]]]
func SequenceT6[E, T1, T2, T3, T4, T5, T6 any](
	t1 IOEither[E, T1],
	t2 IOEither[E, T2],
	t3 IOEither[E, T3],
	t4 IOEither[E, T4],
	t5 IOEither[E, T5],
	t6 IOEither[E, T6],
) IOEither[E, T.Tuple6[T1, T2, T3, T4, T5, T6]] {
	return G.SequenceT6[
		IOEither[E, T.Tuple6[T1, T2, T3, T4, T5, T6]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
	](t1, t2, t3, t4, t5, t6)
}

// SequenceTuple6 converts a [T.Tuple6[IOEither[E, T]]] into a [IOEither[E, T.Tuple6[T1, T2, T3, T4, T5, T6]]]
func SequenceTuple6[E, T1, T2, T3, T4, T5, T6 any](t T.Tuple6[IOEither[E, T1], IOEither[E, T2], IOEither[E, T3], IOEither[E, T4], IOEither[E, T5], IOEither[E, T6]]) IOEither[E, T.Tuple6[T1, T2, T3, T4, T5, T6]] {
	return G.SequenceTuple6[
		IOEither[E, T.Tuple6[T1, T2, T3, T4, T5, T6]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
	](t)
}

// TraverseTuple6 converts a [T.Tuple6[IOEither[E, T]]] into a [IOEither[E, T.Tuple6[T1, T2, T3, T4, T5, T6]]]
func TraverseTuple6[F1 ~func(A1) IOEither[E, T1], F2 ~func(A2) IOEither[E, T2], F3 ~func(A3) IOEither[E, T3], F4 ~func(A4) IOEither[E, T4], F5 ~func(A5) IOEither[E, T5], F6 ~func(A6) IOEither[E, T6], E, A1, A2, A3, A4, A5, A6, T1, T2, T3, T4, T5, T6 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6) func(T.Tuple6[A1, A2, A3, A4, A5, A6]) IOEither[E, T.Tuple6[T1, T2, T3, T4, T5, T6]] {
	return G.TraverseTuple6[IOEither[E, T.Tuple6[T1, T2, T3, T4, T5, T6]]](f1, f2, f3, f4, f5, f6)
}

// Eitherize7 converts a function with 8 parameters returning a tuple into a function with 7 parameters returning a [IOEither[error, R]]
func Eitherize7[F ~func(T1, T2, T3, T4, T5, T6, T7) (R, error), T1, T2, T3, T4, T5, T6, T7, R any](f F) func(T1, T2, T3, T4, T5, T6, T7) IOEither[error, R] {
	return G.Eitherize7[IOEither[error, R]](f)
}

// Uneitherize7 converts a function with 8 parameters returning a tuple into a function with 7 parameters returning a [IOEither[error, R]]
func Uneitherize7[F ~func(T1, T2, T3, T4, T5, T6, T7) IOEither[error, R], T1, T2, T3, T4, T5, T6, T7, R any](f F) func(T1, T2, T3, T4, T5, T6, T7) (R, error) {
	return G.Uneitherize7[IOEither[error, R]](f)
}

// SequenceT7 converts 7 [IOEither[E, T]] into a [IOEither[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]]]
func SequenceT7[E, T1, T2, T3, T4, T5, T6, T7 any](
	t1 IOEither[E, T1],
	t2 IOEither[E, T2],
	t3 IOEither[E, T3],
	t4 IOEither[E, T4],
	t5 IOEither[E, T5],
	t6 IOEither[E, T6],
	t7 IOEither[E, T7],
) IOEither[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	return G.SequenceT7[
		IOEither[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
		IOEither[E, T7],
	](t1, t2, t3, t4, t5, t6, t7)
}

// SequenceTuple7 converts a [T.Tuple7[IOEither[E, T]]] into a [IOEither[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]]]
func SequenceTuple7[E, T1, T2, T3, T4, T5, T6, T7 any](t T.Tuple7[IOEither[E, T1], IOEither[E, T2], IOEither[E, T3], IOEither[E, T4], IOEither[E, T5], IOEither[E, T6], IOEither[E, T7]]) IOEither[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	return G.SequenceTuple7[
		IOEither[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
		IOEither[E, T7],
	](t)
}

// TraverseTuple7 converts a [T.Tuple7[IOEither[E, T]]] into a [IOEither[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]]]
func TraverseTuple7[F1 ~func(A1) IOEither[E, T1], F2 ~func(A2) IOEither[E, T2], F3 ~func(A3) IOEither[E, T3], F4 ~func(A4) IOEither[E, T4], F5 ~func(A5) IOEither[E, T5], F6 ~func(A6) IOEither[E, T6], F7 ~func(A7) IOEither[E, T7], E, A1, A2, A3, A4, A5, A6, A7, T1, T2, T3, T4, T5, T6, T7 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7) func(T.Tuple7[A1, A2, A3, A4, A5, A6, A7]) IOEither[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	return G.TraverseTuple7[IOEither[E, T.Tuple7[T1, T2, T3, T4, T5, T6, T7]]](f1, f2, f3, f4, f5, f6, f7)
}

// Eitherize8 converts a function with 9 parameters returning a tuple into a function with 8 parameters returning a [IOEither[error, R]]
func Eitherize8[F ~func(T1, T2, T3, T4, T5, T6, T7, T8) (R, error), T1, T2, T3, T4, T5, T6, T7, T8, R any](f F) func(T1, T2, T3, T4, T5, T6, T7, T8) IOEither[error, R] {
	return G.Eitherize8[IOEither[error, R]](f)
}

// Uneitherize8 converts a function with 9 parameters returning a tuple into a function with 8 parameters returning a [IOEither[error, R]]
func Uneitherize8[F ~func(T1, T2, T3, T4, T5, T6, T7, T8) IOEither[error, R], T1, T2, T3, T4, T5, T6, T7, T8, R any](f F) func(T1, T2, T3, T4, T5, T6, T7, T8) (R, error) {
	return G.Uneitherize8[IOEither[error, R]](f)
}

// SequenceT8 converts 8 [IOEither[E, T]] into a [IOEither[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]]
func SequenceT8[E, T1, T2, T3, T4, T5, T6, T7, T8 any](
	t1 IOEither[E, T1],
	t2 IOEither[E, T2],
	t3 IOEither[E, T3],
	t4 IOEither[E, T4],
	t5 IOEither[E, T5],
	t6 IOEither[E, T6],
	t7 IOEither[E, T7],
	t8 IOEither[E, T8],
) IOEither[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	return G.SequenceT8[
		IOEither[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
		IOEither[E, T7],
		IOEither[E, T8],
	](t1, t2, t3, t4, t5, t6, t7, t8)
}

// SequenceTuple8 converts a [T.Tuple8[IOEither[E, T]]] into a [IOEither[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]]
func SequenceTuple8[E, T1, T2, T3, T4, T5, T6, T7, T8 any](t T.Tuple8[IOEither[E, T1], IOEither[E, T2], IOEither[E, T3], IOEither[E, T4], IOEither[E, T5], IOEither[E, T6], IOEither[E, T7], IOEither[E, T8]]) IOEither[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	return G.SequenceTuple8[
		IOEither[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
		IOEither[E, T7],
		IOEither[E, T8],
	](t)
}

// TraverseTuple8 converts a [T.Tuple8[IOEither[E, T]]] into a [IOEither[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]]
func TraverseTuple8[F1 ~func(A1) IOEither[E, T1], F2 ~func(A2) IOEither[E, T2], F3 ~func(A3) IOEither[E, T3], F4 ~func(A4) IOEither[E, T4], F5 ~func(A5) IOEither[E, T5], F6 ~func(A6) IOEither[E, T6], F7 ~func(A7) IOEither[E, T7], F8 ~func(A8) IOEither[E, T8], E, A1, A2, A3, A4, A5, A6, A7, A8, T1, T2, T3, T4, T5, T6, T7, T8 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8) func(T.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) IOEither[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	return G.TraverseTuple8[IOEither[E, T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]](f1, f2, f3, f4, f5, f6, f7, f8)
}

// Eitherize9 converts a function with 10 parameters returning a tuple into a function with 9 parameters returning a [IOEither[error, R]]
func Eitherize9[F ~func(T1, T2, T3, T4, T5, T6, T7, T8, T9) (R, error), T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f F) func(T1, T2, T3, T4, T5, T6, T7, T8, T9) IOEither[error, R] {
	return G.Eitherize9[IOEither[error, R]](f)
}

// Uneitherize9 converts a function with 10 parameters returning a tuple into a function with 9 parameters returning a [IOEither[error, R]]
func Uneitherize9[F ~func(T1, T2, T3, T4, T5, T6, T7, T8, T9) IOEither[error, R], T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f F) func(T1, T2, T3, T4, T5, T6, T7, T8, T9) (R, error) {
	return G.Uneitherize9[IOEither[error, R]](f)
}

// SequenceT9 converts 9 [IOEither[E, T]] into a [IOEither[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]]
func SequenceT9[E, T1, T2, T3, T4, T5, T6, T7, T8, T9 any](
	t1 IOEither[E, T1],
	t2 IOEither[E, T2],
	t3 IOEither[E, T3],
	t4 IOEither[E, T4],
	t5 IOEither[E, T5],
	t6 IOEither[E, T6],
	t7 IOEither[E, T7],
	t8 IOEither[E, T8],
	t9 IOEither[E, T9],
) IOEither[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	return G.SequenceT9[
		IOEither[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
		IOEither[E, T7],
		IOEither[E, T8],
		IOEither[E, T9],
	](t1, t2, t3, t4, t5, t6, t7, t8, t9)
}

// SequenceTuple9 converts a [T.Tuple9[IOEither[E, T]]] into a [IOEither[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]]
func SequenceTuple9[E, T1, T2, T3, T4, T5, T6, T7, T8, T9 any](t T.Tuple9[IOEither[E, T1], IOEither[E, T2], IOEither[E, T3], IOEither[E, T4], IOEither[E, T5], IOEither[E, T6], IOEither[E, T7], IOEither[E, T8], IOEither[E, T9]]) IOEither[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	return G.SequenceTuple9[
		IOEither[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
		IOEither[E, T7],
		IOEither[E, T8],
		IOEither[E, T9],
	](t)
}

// TraverseTuple9 converts a [T.Tuple9[IOEither[E, T]]] into a [IOEither[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]]
func TraverseTuple9[F1 ~func(A1) IOEither[E, T1], F2 ~func(A2) IOEither[E, T2], F3 ~func(A3) IOEither[E, T3], F4 ~func(A4) IOEither[E, T4], F5 ~func(A5) IOEither[E, T5], F6 ~func(A6) IOEither[E, T6], F7 ~func(A7) IOEither[E, T7], F8 ~func(A8) IOEither[E, T8], F9 ~func(A9) IOEither[E, T9], E, A1, A2, A3, A4, A5, A6, A7, A8, A9, T1, T2, T3, T4, T5, T6, T7, T8, T9 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9) func(T.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) IOEither[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	return G.TraverseTuple9[IOEither[E, T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]](f1, f2, f3, f4, f5, f6, f7, f8, f9)
}

// Eitherize10 converts a function with 11 parameters returning a tuple into a function with 10 parameters returning a [IOEither[error, R]]
func Eitherize10[F ~func(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) (R, error), T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R any](f F) func(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) IOEither[error, R] {
	return G.Eitherize10[IOEither[error, R]](f)
}

// Uneitherize10 converts a function with 11 parameters returning a tuple into a function with 10 parameters returning a [IOEither[error, R]]
func Uneitherize10[F ~func(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) IOEither[error, R], T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R any](f F) func(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) (R, error) {
	return G.Uneitherize10[IOEither[error, R]](f)
}

// SequenceT10 converts 10 [IOEither[E, T]] into a [IOEither[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]]
func SequenceT10[E, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](
	t1 IOEither[E, T1],
	t2 IOEither[E, T2],
	t3 IOEither[E, T3],
	t4 IOEither[E, T4],
	t5 IOEither[E, T5],
	t6 IOEither[E, T6],
	t7 IOEither[E, T7],
	t8 IOEither[E, T8],
	t9 IOEither[E, T9],
	t10 IOEither[E, T10],
) IOEither[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	return G.SequenceT10[
		IOEither[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
		IOEither[E, T7],
		IOEither[E, T8],
		IOEither[E, T9],
		IOEither[E, T10],
	](t1, t2, t3, t4, t5, t6, t7, t8, t9, t10)
}

// SequenceTuple10 converts a [T.Tuple10[IOEither[E, T]]] into a [IOEither[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]]
func SequenceTuple10[E, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](t T.Tuple10[IOEither[E, T1], IOEither[E, T2], IOEither[E, T3], IOEither[E, T4], IOEither[E, T5], IOEither[E, T6], IOEither[E, T7], IOEither[E, T8], IOEither[E, T9], IOEither[E, T10]]) IOEither[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	return G.SequenceTuple10[
		IOEither[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
		IOEither[E, T1],
		IOEither[E, T2],
		IOEither[E, T3],
		IOEither[E, T4],
		IOEither[E, T5],
		IOEither[E, T6],
		IOEither[E, T7],
		IOEither[E, T8],
		IOEither[E, T9],
		IOEither[E, T10],
	](t)
}

// TraverseTuple10 converts a [T.Tuple10[IOEither[E, T]]] into a [IOEither[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]]
func TraverseTuple10[F1 ~func(A1) IOEither[E, T1], F2 ~func(A2) IOEither[E, T2], F3 ~func(A3) IOEither[E, T3], F4 ~func(A4) IOEither[E, T4], F5 ~func(A5) IOEither[E, T5], F6 ~func(A6) IOEither[E, T6], F7 ~func(A7) IOEither[E, T7], F8 ~func(A8) IOEither[E, T8], F9 ~func(A9) IOEither[E, T9], F10 ~func(A10) IOEither[E, T10], E, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9, f10 F10) func(T.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) IOEither[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	return G.TraverseTuple10[IOEither[E, T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]](f1, f2, f3, f4, f5, f6, f7, f8, f9, f10)
}
