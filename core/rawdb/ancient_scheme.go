// Copyright 2022 The go-ethereum Authors
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

package rawdb

import "path/filepath"

// The list of table names of chain freezer.
const (
	// chainFreezerHeaderTable indicates the name of the freezer header table.
	chainFreezerHeaderTable = "headers"

	// chainFreezerHashTable indicates the name of the freezer canonical hash table.
	chainFreezerHashTable = "hashes"

	// chainFreezerBodiesTable indicates the name of the freezer block body table.
	chainFreezerBodiesTable = "bodies"

	// chainFreezerReceiptTable indicates the name of the freezer receipts table.
	chainFreezerReceiptTable = "receipts"

	// chainFreezerDifficultyTable indicates the name of the freezer total difficulty table.
	chainFreezerDifficultyTable = "diffs"
)

// chainFreezerNoSnappy configures whether compression is disabled for the ancient-tables.
// Hashes and difficulties don't compress well.
var chainFreezerNoSnappy = map[string]bool{
	chainFreezerHeaderTable:     false,
	chainFreezerHashTable:       true,
	chainFreezerBodiesTable:     false,
	chainFreezerReceiptTable:    false,
	chainFreezerDifficultyTable: true,
}

const (
	// trieHistoryTableSize defines the maximum size of freezer data files.
	trieHistoryTableSize = 2 * 1000 * 1000 * 1000

	// trieHistoryTable indicates the name of the trie history table.
	trieHistoryTable = "history"
)

// trieHistoryFreezerNoSnappy configures whether compression is disabled for the ancient
// trie histories
var trieHistoryFreezerNoSnappy = map[string]bool{
	trieHistoryTable: false,
}

// The list of identifiers of ancient stores.
var (
	chainFreezerName       = "chain"       // the folder name of chain segment ancient store.
	trieHistoryFreezerName = "triehistory" // the folder name of trie history ancient store.
)

// freezers the collections of all builtin freezers.
var freezers = []string{chainFreezerName, trieHistoryFreezerName}

// NewTrieHistoryFreezer initializes the freezer for trie histories.
func NewTrieHistoryFreezer(ancientDir string, readOnly bool) (*ResettableFreezer, error) {
	return NewResettableFreezer(filepath.Join(ancientDir, trieHistoryFreezerName), "eth/db/triehistory", readOnly, trieHistoryTableSize, trieHistoryFreezerNoSnappy)
}
