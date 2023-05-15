package gencrypt

import "golang.org/x/crypto/bcrypt"

// EncryptPassword
//
//	@Description: 加密密码
//	@param password
//	@return string
//	@return error
func EncryptPassword(password string) (string, error) {
	// 加密密码，bcrypt.DefaultCost 代表使用默认加密成本
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 如果有错误则返回异常，加密后的空字符串返回为空字符串，因为加密失败
		return "", err
	} else {
		// 返回加密后的密码和空异常
		return string(encryptPassword), nil
	}
}

// EqualsPassword
//
//	@Description: 匹配密码
//	@param password 输入的密码
//	@param encryptPassword 加密密码
//	@return bool
func EqualsPassword(password, encryptPassword string) bool {
	// 使用 bcrypt 当中的 CompareHashAndPassword 对比密码是否正确，第一个参数为加密后的密码，第二个参数为未加密的密码
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	// 对比密码是否正确会返回一个异常 nil 就证明密码正确
	return err == nil
}
