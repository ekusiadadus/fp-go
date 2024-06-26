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

package file

import (
	"context"
	"fmt"

	RIO "github.com/IBM/fp-go/context/readerio"
	R "github.com/IBM/fp-go/context/readerioeither"
	"github.com/IBM/fp-go/errors"
	F "github.com/IBM/fp-go/function"
	IO "github.com/IBM/fp-go/io"
	J "github.com/IBM/fp-go/json"
)

type RecordType struct {
	Data string `json:"data"`
}

func getData(r RecordType) string {
	return r.Data
}

func ExampleReadFile() {

	data := F.Pipe4(
		ReadFile("./data/file.json"),
		R.ChainEitherK(J.Unmarshal[RecordType]),
		R.ChainFirstIOK(IO.Logf[RecordType]("Log: %v")),
		R.Map(getData),
		R.GetOrElse(F.Flow2(
			errors.ToString,
			RIO.Of[string],
		)),
	)

	result := data(context.Background())

	fmt.Println(result())

	// Output: Carsten
}
