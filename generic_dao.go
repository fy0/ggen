package gen

import (
	"context"

	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

// GenericDao 是一个泛型类，T 代表带有 field 的 DO 类型（如 systemRoleDO），R 代表原始的 DO 类型（如 do.SystemRoleDO）
type GenericDao[T any, R any] struct {
	DO
	Fields T
}

// Debug 启用调试模式
func (g GenericDao[T, R]) Debug() *GenericDao[T, R] {
	return g.withDO(g.DO.Debug())
}

// WithContext 设置上下文
func (g GenericDao[T, R]) WithContext(ctx context.Context) *GenericDao[T, R] {
	return g.withDO(g.DO.WithContext(ctx))
}

// ReadDB 使用读数据库
func (g GenericDao[T, R]) ReadDB() *GenericDao[T, R] {
	return g.withDO(g.DO.Clauses(dbresolver.Read))
}

// WriteDB 使用写数据库
func (g GenericDao[T, R]) WriteDB() *GenericDao[T, R] {
	return g.withDO(g.DO.Clauses(dbresolver.Write))
}

// Session 设置会话配置
func (g GenericDao[T, R]) Session(config *gorm.Session) *GenericDao[T, R] {
	return g.withDO(g.DO.Session(config))
}

// Clauses 添加子句
func (g GenericDao[T, R]) Clauses(conds ...clause.Expression) *GenericDao[T, R] {
	return g.withDO(g.DO.Clauses(conds...))
}

// Returning 设置返回值
func (g GenericDao[T, R]) Returning(value interface{}, columns ...string) *GenericDao[T, R] {
	return g.withDO(g.DO.Returning(value, columns...))
}

// Not 添加 NOT 条件
func (g GenericDao[T, R]) Not(conds ...Condition) *GenericDao[T, R] {
	return g.withDO(g.DO.Not(conds...))
}

// Or 添加 OR 条件
func (g GenericDao[T, R]) Or(conds ...Condition) *GenericDao[T, R] {
	return g.withDO(g.DO.Or(conds...))
}

// Select 选择字段
func (g GenericDao[T, R]) Select(conds ...field.Expr) *GenericDao[T, R] {
	return g.withDO(g.DO.Select(conds...))
}

// Where 添加 WHERE 条件
func (g GenericDao[T, R]) Where(conds ...Condition) *GenericDao[T, R] {
	return g.withDO(g.DO.Where(conds...))
}

// Order 添加排序
func (g GenericDao[T, R]) Order(conds ...field.Expr) *GenericDao[T, R] {
	return g.withDO(g.DO.Order(conds...))
}

// Distinct 去重
func (g GenericDao[T, R]) Distinct(cols ...field.Expr) *GenericDao[T, R] {
	return g.withDO(g.DO.Distinct(cols...))
}

// Omit 忽略字段
func (g GenericDao[T, R]) Omit(cols ...field.Expr) *GenericDao[T, R] {
	return g.withDO(g.DO.Omit(cols...))
}

// Join 内连接
func (g GenericDao[T, R]) Join(table schema.Tabler, on ...field.Expr) *GenericDao[T, R] {
	return g.withDO(g.DO.Join(table, on...))
}

// LeftJoin 左连接
func (g GenericDao[T, R]) LeftJoin(table schema.Tabler, on ...field.Expr) *GenericDao[T, R] {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

// RightJoin 右连接
func (g GenericDao[T, R]) RightJoin(table schema.Tabler, on ...field.Expr) *GenericDao[T, R] {
	return g.withDO(g.DO.RightJoin(table, on...))
}

// Group 分组
func (g GenericDao[T, R]) Group(cols ...field.Expr) *GenericDao[T, R] {
	return g.withDO(g.DO.Group(cols...))
}

// Having 添加 HAVING 条件
func (g GenericDao[T, R]) Having(conds ...Condition) *GenericDao[T, R] {
	return g.withDO(g.DO.Having(conds...))
}

// Limit 限制数量
func (g GenericDao[T, R]) Limit(limit int) *GenericDao[T, R] {
	return g.withDO(g.DO.Limit(limit))
}

// Offset 偏移量
func (g GenericDao[T, R]) Offset(offset int) *GenericDao[T, R] {
	return g.withDO(g.DO.Offset(offset))
}

// Scopes 应用作用域
func (g GenericDao[T, R]) Scopes(funcs ...func(Dao) Dao) *GenericDao[T, R] {
	return g.withDO(g.DO.Scopes(funcs...))
}

// Unscoped 包含软删除的记录
func (g GenericDao[T, R]) Unscoped() *GenericDao[T, R] {
	return g.withDO(g.DO.Unscoped())
}

// Create 创建记录
func (g GenericDao[T, R]) Create(values ...*R) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

// CreateInBatches 批量创建记录
func (g GenericDao[T, R]) CreateInBatches(values []*R, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save 保存记录（相当于 UPSERT）
func (g GenericDao[T, R]) Save(values ...*R) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

// First 查询第一条记录
func (g GenericDao[T, R]) First() (*R, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*R), nil
	}
}

// Take 查询一条记录
func (g GenericDao[T, R]) Take() (*R, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*R), nil
	}
}

// Last 查询最后一条记录
func (g GenericDao[T, R]) Last() (*R, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*R), nil
	}
}

// Find 查询多条记录
func (g GenericDao[T, R]) Find() ([]*R, error) {
	result, err := g.DO.Find()
	return result.([]*R), err
}

// FindInBatch 批量查询记录
func (g GenericDao[T, R]) FindInBatch(batchSize int, fc func(tx Dao, batch int) error) (results []*R, err error) {
	buf := make([]*R, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

// FindInBatches 批量查询记录到指定切片
func (g GenericDao[T, R]) FindInBatches(result *[]*R, batchSize int, fc func(tx Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

// Attrs 设置属性（仅在记录不存在时使用）
func (g GenericDao[T, R]) Attrs(attrs ...field.AssignExpr) *GenericDao[T, R] {
	return g.withDO(g.DO.Attrs(attrs...))
}

// Assign 分配属性
func (g GenericDao[T, R]) Assign(attrs ...field.AssignExpr) *GenericDao[T, R] {
	return g.withDO(g.DO.Assign(attrs...))
}

// Joins 连接关联
func (g GenericDao[T, R]) Joins(fields ...field.RelationField) *GenericDao[T, R] {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

// Preload 预加载关联
func (g GenericDao[T, R]) Preload(fields ...field.RelationField) *GenericDao[T, R] {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

// FirstOrInit 查询第一条记录或初始化
func (g GenericDao[T, R]) FirstOrInit() (*R, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*R), nil
	}
}

// FirstOrCreate 查询第一条记录或创建
func (g GenericDao[T, R]) FirstOrCreate() (*R, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*R), nil
	}
}

// FindByPage 分页查询
func (g GenericDao[T, R]) FindByPage(offset int, limit int) (result []*R, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

// ScanByPage 分页扫描到指定结构
func (g GenericDao[T, R]) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

// Scan 扫描到指定结构
func (g GenericDao[T, R]) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

// Delete 删除记录
func (g GenericDao[T, R]) Delete(models ...*R) (result ResultInfo, err error) {
	return g.DO.Delete(models)
}

// Count 统计记录数
func (g GenericDao[T, R]) Count() (int64, error) {
	return g.DO.Count()
}

// withDO 内部方法，用于链式调用
func (g *GenericDao[T, R]) withDO(do Dao) *GenericDao[T, R] {
	g.DO = *do.(*DO)
	return g
}
