// Copyright 2015 CoreOS, Inc. All rights reserved.
// Copyright 2014 Dropbox, Inc. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license,
// which can be found in the LICENSE file.

package sqlbuilder

import (
	"bytes"
)

type SubqueryClause interface {
	SerializeSql(d Dialect, out *bytes.Buffer) error
}

type Clause interface {
	SerializeSql(d Dialect, out *bytes.Buffer) error
}

// A clause that can be used in order by
type OrderByClause interface {
	Clause
	isOrderByClauseInterface
}

// An expression
type Expression interface {
	Clause
	isExpressionInterface
}

type BoolExpression interface {
	Clause
	isBoolExpressionInterface
}

// A clause that is selectable.
type Projection interface {
	Clause
	isProjectionInterface
	SerializeSqlForColumnList(includeTableName bool, d Dialect, out *bytes.Buffer) error
}

//
// Boiler plates ...
//

type isOrderByClauseInterface interface {
	isOrderByClauseType()
}

type isOrderByClause struct {
}

func (o *isOrderByClause) isOrderByClauseType() {
}

type isExpressionInterface interface {
	isExpressionType()
}

type isExpression struct {
	isOrderByClause // can always use expression in order by.
}

func (e *isExpression) isExpressionType() {
}

type isBoolExpressionInterface interface {
	isExpressionInterface
	isBoolExpressionType()
}

type isBoolExpression struct {
}

func (e *isBoolExpression) isBoolExpressionType() {
}

type isProjectionInterface interface {
	isProjectionType()
}

type isProjection struct {
}

func (p *isProjection) isProjectionType() {
}
