package mappers

const (
	FindUserById rawSQL = 1 << iota
	UpdateUserPwd
)

var (
	//查询用户信息
	findUserByIdSQL = `
        SELECT 
    		user_id,account,password,role_name,employee_id,update_time,state
		FROM
    		sd_sys_user
		WHERE
			user_id = ?

    `
	// 更新密码信息
	updateUserPwdSQL = `
        UPDATE
            sd_sys_user
        SET
            password = ?,
            update_time = NOW()
        WHERE
            user_id = ?
    `
)

// 获取用户相关SQL
func FetchUserRawSQL(t rawSQL) string {
	switch t {
	case FindUserById:
		return findUserByIdSQL
	case UpdateUserPwd:
		return updateUserPwdSQL
	default:
		return ""
	}
}
