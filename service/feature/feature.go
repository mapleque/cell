package feature

import (
	. "github.com/coral"
	"github.com/coral/db"

	. "github.com/tellus/constant"
)

// @author yangyang
// @review
// 批量获取用户当前资源状态
// @param userId <用户id>
// @param productIds <要获取的productIds>
// @return <info, err>
// 		err  =   0 success
// 				-1 db error
//		info [{
//			product_id:<资源id>,
//			     valid:<当前状态>,
//			     start:<开始时间>,
//			       end:<结束时间>,
//		}, ...]
func GetValidBatch(
	userId int, productIds []int) ([]map[string]interface{}, int) {
	productIdsWhere := ""
	params := []interface{}{userId}
	for _, productId := range productIds {
		if len(productIdsWhere) > 0 {
			productIdsWhere = productIdsWhere + ","
		}
		productIdsWhere = productIdsWhere + "?"
		params = append(params, productId)
	}
	conn := db.UseDB(DEF_DB)
	infos := conn.Select(
		`SELECT product_id, start_time AS start, end_time AS end,
		IF(end_time >= NOW() AND start_time <= NOW(), 1, 0) AS valid
		FROM feature
		WHERE user_id = ? AND product_id IN (`+productIdsWhere+`)`,
		params...)
	return infos, 0
}

// @author yangyang
// @review
// 修改用户系统资源数量
// @param userId <用户id>
// @param productId<资源id>
// @param start <开始时间, 可能为空，与end不能同时为空>
// @param end <结束时间，可能为空，与start不能同时为空>
// @return err
// 		err  =   0 success
// 				-1 db error
// 				 1 资源不存在或者类型不对
func Update(userId, productId int, start, end string) int {
	conn := db.UseDB(DEF_DB)
	products := conn.Select(`
		SELECT id FROM product WHERE id = ? AND type = ? LIMIT 1`,
		productId, PRODUCT_TYPE_FEATURE)
	if len(products) < 1 {
		return RTInt(1)
	}
	infos := conn.Select(
		`SELECT id FROM feature
		WHERE user_id = ? AND product_id = ? LIMIT 1`,
		userId, productId)
	if len(infos) < 1 {
		switch {
		case start == "":
			ret := conn.Insert(
				`INSERT INTO feature (user_id, product_id, start_time, end_time)
			VALUES (?,?,NOW,?)`,
				userId, productId, end)
			if ret <= 0 {
				return RTInt(-1)
			}
			return 0
		case end == "":
			ret := conn.Insert(
				`INSERT INTO feature (user_id, product_id, start_time, end_time)
			VALUES (?,?,?,NOW())`,
				userId, productId, start)
			if ret <= 0 {
				return RTInt(-1)
			}
			return 0
		default:
			ret := conn.Insert(
				`INSERT INTO feature (user_id, product_id, start_time, end_time)
			VALUES (?,?,?,?)`,
				userId, productId, start, end)
			if ret <= 0 {
				return RTInt(-1)
			}
			return 0
		}
	} else {
		switch {
		case start == "":
			ret := conn.Update(
				`UPDATE feature SET end_time = ?
			WHERE id = ? LIMIT 1`,
				end, infos[0]["id"])
			if ret != 1 {
				return RTInt(-1)
			}
			return 0
		case end == "":
			ret := conn.Update(
				`UPDATE feature SET start_time = ?
			WHERE id = ? LIMIT 1`,
				start, infos[0]["id"])
			if ret != 1 {
				return RTInt(-1)
			}
			return 0
		default:
			ret := conn.Update(
				`UPDATE feature SET start_time = ?, end_time = ?
			WHERE id = ? LIMIT 1`,
				start, end, infos[0]["id"])
			if ret != 1 {
				return RTInt(-1)
			}
			return 0
		}
	}
}
