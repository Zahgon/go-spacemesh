// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package util

import (
	"reflect"
)

// UnmarshalFixedJSON decodes the input as a string with 0x prefix. The length of out
// determines the required input length. This function is commonly used to implement the
// UnmarshalJSON method for fixed-size types.
func UnmarshalFixedJSON(typ reflect.Type, input, out []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// check for quoted string

// UnmarshalFixedText decodes the input as a string with 0x prefix. The length of out
// determines the required input length. This function is commonly used to implement the
// UnmarshalText method for fixed-size types.
func UnmarshalFixedText(typename string, input, out []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Pre-verify syntax before modifying out.

func bytesHave0xPrefix(input []byte) bool { _ = "STUB: not implemented"; return false }

func checkText(input []byte, wantPrefix bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// empty strings are allowed

func wrapTypeError(err error, typ reflect.Type) error { _ = "STUB: not implemented"; return nil }

const badNibble = ^uint64(0)

func decodeNibble(in byte) uint64 { _ = "STUB: not implemented"; return 0 }

func Base64Encode(src []byte) []byte { _ = "STUB: not implemented"; return nil }

func Base64Decode(dst, src []byte) error { _ = "STUB: not implemented"; return nil }
