// Copyright 2015 CoreOS, Inc. All rights reserved.
// Copyright 2014 Dropbox, Inc. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license,
// which can be found in the LICENSE file.

// A library for generating SQL programmatically.
//
// Known limitations for SELECT queries:
//  - does not currently support join table alias (and hence self join)
//  - does not support NATURAL joins and join USING
//
// Known limitation for INSERT statements:
//  - does not support "INSERT INTO SELECT"
//
// Known limitation for UPDATE statements:
//  - does not support update without a WHERE clause (since it is dangerous)
//  - does not support multi-table update
//
// Known limitation for DELETE statements:
//  - does not support delete without a WHERE clause (since it is dangerous)
//  - does not support multi-table delete
package sqlbuilder
