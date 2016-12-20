package procuct

import (
	. "github.com/coral"
	"github.com/coral/db"

	. "github.com/cell/constant"
)

// @author yangyang
// @review
// 获取系统资源列表
// @param pageSize <分页大小>
// @param pageStart <分页开始>
// @return <info, err>
// 		err  =   0 success
// 				-1 db error
//		info [{
//			id:<资源id>,
//			type:<资源类型>,
//			name:<资源名称>,
//			additional:<资源描述>,
//			update_time:<资源最后更新时间>,
//			create_time:<资源首次添加时间>,
//		}, ...]
func List(pageSize, pageStart int) ([]map[string]interface{}, int) {
	conn := db.UseDB(DEF_DB)
	ret := conn.Select(
		`SELECT id, type, name, additional, update_time, create_time
		FROM product WHERE id > ? ORDER BY id LIMIT ?`, pageStart, pageSize)
	return ret, 0
}

// @author yangyang
// @review
// 新建系统资源
// @param productType <资源类型>
// @param name <资源名称>
// @param additional <资源描述>
// @return <productId, err>
// 		err  =   0 success
// 				-1 db error
func New(productType, name, additional string) (int, int) {
	conn := db.UseDB(DEF_DB)
	ret := conn.Insert(
		`INSERT INTO product (type, name, additional) VALUES (?,?,?)`,
		productType, name, additional)
	if ret > 0 {
		return Int(ret), 0
	}
	return 0, RTInt(-1)
}

// @author yangyang
// @review
// 修改系统资源
// @param productId<资源id>
// @param name <资源名称>
// @param additional <资源描述>
// @return err
// 		err  =   0 success
// 				-1 db error
func Update(productId int, name, additional string) int {
	conn := db.UseDB(DEF_DB)
	ret := conn.Update(
		`UPDATE product SET name = ?, additional = ?, update_time = NOW()
		WHERE id = ? LIMIT 1`,
		name, additional, productId)
	if ret != 1 {
		return RTInt(-1)
	}
	return 0
}
