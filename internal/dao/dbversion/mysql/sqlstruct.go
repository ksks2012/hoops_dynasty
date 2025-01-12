// Code generated by go-literal-code-gen. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
)

// ** SQL schema external filter

const metaKeyDistributedLockMetaSchemaRev = "distributed-lock-meta.schema"

const currentDistributedLockMetaSchemaRev = 1

type schemaRevision struct {
	// TODO: unknown translation mode 0 for symbol [DistributedLockMeta]
}

func (rev *schemaRevision) IsUpToDate() bool {
	return true
}

type schemaManager struct {
	referenceTableName string
	ctx                context.Context
	conn               *sql.DB
}

func (m *schemaManager) FetchSchemaRevision() (schemaRev *schemaRevision, err error) {
	schemaRev = &schemaRevision{}
	return schemaRev, nil
}

// upgrade routine for symbol not generated: DistributedLockMeta

// ** Generated code for 1 table entries
