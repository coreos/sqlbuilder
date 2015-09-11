// Copyright 2015 CoreOS, Inc. All rights reserved.
// Copyright 2014 Dropbox, Inc. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license,
// which can be found in the LICENSE file.

package sqlbuilder

var table1Col1 = IntColumn("col1", Nullable)
var table1Col2 = IntColumn("col2", Nullable)
var table1Col3 = IntColumn("col3", Nullable)
var table1Col4 = DateTimeColumn("col4", Nullable)
var table1 = NewTable(
	"table1",
	table1Col1,
	table1Col2,
	table1Col3,
	table1Col4)

var table2Col3 = IntColumn("col3", Nullable)
var table2Col4 = IntColumn("col4", Nullable)
var table2 = NewTable(
	"table2",
	table2Col3,
	table2Col4)

var table3Col1 = IntColumn("col1", Nullable)
var table3Col2 = IntColumn("col2", Nullable)
var table3 = NewTable(
	"table3",
	table3Col1,
	table3Col2)
