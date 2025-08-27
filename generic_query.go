package gen

import (
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type QueryObject[T any] interface {
	DB() *gorm.DB
	ReplaceDB(db *gorm.DB) T
	Clone(db *gorm.DB) T
}

type QueryBase[T QueryObject[T]] struct {
	Obj T
}

func (q *QueryBase[T]) Available() bool { return q.Obj.DB() != nil }

func (q *QueryBase[T]) ReadDB() T {
	return q.ReplaceDB(q.Obj.DB().Clauses(dbresolver.Read))
}

func (q *QueryBase[T]) WriteDB() T {
	return q.ReplaceDB(q.Obj.DB().Clauses(dbresolver.Write))
}

func (q *QueryBase[T]) ReplaceDB(db *gorm.DB) T {
	return q.Obj.Clone(db)
}

func (q *QueryBase[T]) Transaction(fc func(tx T) error, opts ...*sql.TxOptions) error {
	return q.Obj.DB().Transaction(func(tx *gorm.DB) error { return fc(q.Obj.ReplaceDB(tx)) }, opts...)
}

// WithTx0 执行数据库操作，支持外部传入事务或自动创建事务（泛型版本）
func WithTx0[Q QueryObject[Q]](tx Q, fn func(tx Q) error, opts ...*sql.TxOptions) error {
	// 如果已经处于事务中，直接使用现有事务
	if _, ok := tx.DB().InstanceGet("gorm:transaction"); ok {
		return fn(tx)
	}

	// 否则开启新事务
	return tx.DB().Transaction(func(txDB *gorm.DB) error {
		txQuery := tx.ReplaceDB(txDB)
		return fn(txQuery)
	}, opts...)
}

func WithTx[T any, Q QueryObject[Q]](q Q, fn func(tx Q) (T, error), opts ...*sql.TxOptions) (T, error) {
	var zero T
	// 检查是否已在事务中
	if _, ok := q.DB().InstanceGet("gorm:transaction"); ok {
		return fn(q)
	}

	// 开启新事务
	var result T
	err := q.DB().Transaction(func(tx *gorm.DB) error {
		txQuery := q.ReplaceDB(tx)
		var innerErr error
		result, innerErr = fn(txQuery)
		return innerErr
	}, opts...)

	if err != nil {
		return zero, err
	}
	return result, nil
}
