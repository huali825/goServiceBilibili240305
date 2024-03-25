package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	fmt.Println("这是一个测试")
	//pc, _, _, _ := runtime.Caller(1)
	//funcName := runtime.FuncForPC(pc).Name()
	//fmt.Println(funcName)

	password := "qsgctys711!@#"
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(encrypted)

	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("两个密码比较成功了")
	}
	assert.NoError(t, err)
}
