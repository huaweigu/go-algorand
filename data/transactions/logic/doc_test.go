// Copyright (C) 2019-2020 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package logic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpDocs(t *testing.T) {
	opsSeen := make(map[string]bool, len(OpSpecs))
	for _, op := range OpSpecs {
		opsSeen[op.Name] = false
	}
	for _, od := range opDocList {
		_, exists := opsSeen[od.a]
		if !exists {
			t.Errorf("error: doc for op %#v that does not exist in OpSpecs", od.a)
		}
		opsSeen[od.a] = true
	}
	for op, seen := range opsSeen {
		if !seen {
			t.Errorf("error: doc for op %#v missing", op)
		}
	}
}

func TestOpGroupCoverage(t *testing.T) {
	opsSeen := make(map[string]bool, len(OpSpecs))
	for _, op := range OpSpecs {
		opsSeen[op.Name] = false
	}
	for _, og := range OpGroupList {
		for _, name := range og.Ops {
			_, exists := opsSeen[name]
			if !exists {
				t.Errorf("error: op %#v in group list but not in OpSpecs\n", name)
				continue
			}
			opsSeen[name] = true
		}
	}
	for name, seen := range opsSeen {
		if !seen {
			t.Errorf("warning: op %#v not in any group list\n", name)
		}
	}
}

func TestOpDoc(t *testing.T) {
	xd := OpDoc("txn")
	require.NotEmpty(t, xd)
	xd = OpDoc("NOT AN INSTRUCTION")
	require.Empty(t, xd)
}

func TestOpImmediateNote(t *testing.T) {
	xd := OpImmediateNote("txn")
	require.NotEmpty(t, xd)
	xd = OpImmediateNote("+")
	require.Empty(t, xd)
}

func TestOpDocExtra(t *testing.T) {
	xd := OpDocExtra("bnz")
	require.NotEmpty(t, xd)
	xd = OpDocExtra("+")
	require.Empty(t, xd)
}

func TestOpCost(t *testing.T) {
	c := OpCost("+")
	require.Equal(t, 1, c)
	c = OpCost("sha256")
	require.True(t, c > 1)
}

func TestOpSize(t *testing.T) {
	c := OpSize("+")
	require.Equal(t, 1, c)
	c = OpSize("intc")
	require.Equal(t, 2, c)
}

func TestTypeNameDescription(t *testing.T) {
	require.Equal(t, len(TxnTypeNames), len(typeEnumDescriptions))
	for i, a := range TxnTypeNames {
		b := TypeNameDescription(a)
		require.Equal(t, b, typeEnumDescriptions[i].b)
	}
	require.Equal(t, "invalid type name", TypeNameDescription("invalid type name"))
}
