package security

import "golang.org/x/crypto/bcrypt"

func BcryptHashPassword(password string)(resultHash string, err error){
	resultHashByte, err :=bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil{
		return resultHash, err
	}
	resultHash = string(resultHashByte)
	return resultHash, nil
}

func BcryptCheckPasswordHash(password string, hash string) (bool){
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}
