package resource

import (
	. "github.com/coral"
	"github.com/coral/db"

	. "github.com/cell/constant"
)

// @author yangyang
// @review
// 批量获取用户当前资源数量
// @param userId <用户id>
// @param productIds <要获取的productIds>
// @return <info, err>
// 		err  =   0 success
// 				-1 db error
//		info [{
//			product_id:<资源id>,
//			amount:<当前数量>,
//		}, ...]
func GetAmountBatch(
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
		`SELECT product_id, amount FROM resource
		WHERE user_id = ? AND product_id IN (`+productIdsWhere+`)`,
		params...)
	return infos, 0
}

// @author yangyang
// @review
// 修改用户系统资源数量
// @param userId <用户id>
// @param productId <资源id>
// @param amount <资源数量>
// @return err
// 		err  =   0 success
// 				-1 db error
// 				 1 资源不足
// 				 2 资源不存在或者类型不对
func UpdateAmount(userId, productId, amount int) (int, int) {
	conn := db.UseDB(DEF_DB)
	products := conn.Select(`
		SELECT id FROM product WHERE id = ? AND type = ? LIMIT 1`,
		productId, PRODUCT_TYPE_RESOURCE)
	if len(products) < 1 {
		return 0, RTInt(2)
	}
	infos := conn.Select(
		`SELECT id, amount FROM resource
		WHERE user_id = ? AND product_id = ? LIMIT 1`,
		userId, productId)
	if len(infos) < 1 {
		if amount < 0 {
			return amount, RTInt(1)
		}
		ret := conn.Insert(
			`INSERT INTO resource (user_id, product_id, amount) VALUES (?,?,?)`,
			userId, productId, amount)
		if ret <= 0 {
			return 0, RTInt(-1)
		}
		return amount, 0
	} else {
		tarAmount := Int(infos[0]["amount"]) + amount
		if tarAmount < 0 {
			return tarAmount, RTInt(1)
		}
		ret := conn.Update(
			`UPDATE resource SET amount = ?
			WHERE id = ? AND amount = ? LIMIT 1`,
			tarAmount, infos[0]["id"], infos[0]["amount"])
		if ret != 1 {
			return 0, RTInt(-1)
		}
		return tarAmount, 0
	}
}
