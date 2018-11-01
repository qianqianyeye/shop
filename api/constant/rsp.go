/*
* 错误码定义
 */
package constant

var mRsp map[int]string

const (
	/*
	 * 系统
	 */
	RC_OK = 0 //&Result{0,}       // 成功

	/*
	* 业务级公用
	 */
	RC_SYS_ERR             = 1000000 // 系统错误
	RC_INTERFACE_NOT_FOUND = 1000001 // 接口未找到
	RC_TIMEOUT             = 1000002 // 请求超时
	RC_UNAUTHORIZED        = 1000003 // 未授权
	RC_PARM_ERR            = 1000004 // 参数错误

	/*
	* 管理员模块
	 */
	RC_ADMIN_NOT_FOUND = 2000001 // 管理员不存在
	RC_VALID_CODE_ERR  = 2000002 // 验证码错误
	RC_ADMIN_EXISTED   = 2000003 // 用户已存在
	RC_PASSWORD_ERR    = 2000004 // 管理员不存在
	RC_APPS_NOT_FOUND  = 2000005 // 未找到可用应用
	RC_PERMISSION_DENY = 2000006 // 没有权限

	RC_ORDER_NOT_FOUND  = 2100001 // 订单不存在
	RC_ORDER_REPAIR_ERR = 2100002 // 订单修复失败

	RC_CHANNEL_EXISTED = 2200003  //渠道已存在

	RC_UPLOAD_FORMAT = 3300001  //图片格式不正确
	RC_UPLOAD_FALL = 3300001  //图片格式不正确
)

func init() {
	mRsp = make(map[int]string, 0)
	mRsp[RC_OK] = "ok"
	mRsp[RC_SYS_ERR] = "sys err"
	mRsp[RC_INTERFACE_NOT_FOUND] = "interface not found"
	mRsp[RC_TIMEOUT] = "timeout"
	mRsp[RC_UNAUTHORIZED] = "unauthorized"
	mRsp[RC_PARM_ERR] = "param err"

	/*
	* 管理员模块
	 */
	mRsp[RC_ADMIN_NOT_FOUND] = "admin not found"
	mRsp[RC_VALID_CODE_ERR] = "valid-code err"
	mRsp[RC_ADMIN_EXISTED] = "admin existed"
	mRsp[RC_PASSWORD_ERR] = "password err"
	mRsp[RC_APPS_NOT_FOUND] = "apps not found"
	mRsp[RC_PERMISSION_DENY] = "permission deny"

	mRsp[RC_ORDER_NOT_FOUND] = "order not found"
	mRsp[RC_ORDER_REPAIR_ERR] = "order repair err"

	mRsp[RC_CHANNEL_EXISTED] = "channel existed"

	mRsp[RC_UPLOAD_FORMAT] = "check img format"
	mRsp[RC_UPLOAD_FALL] = "image upload fall"
}

// 获取错误码消息值
func M(code int) string {
	if v, ok := mRsp[code]; ok {
		return v
	}
	return "unknow"
}
