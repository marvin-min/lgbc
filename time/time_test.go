package cos

import "testing"

func Test_CosName(t *testing.T)  {
	 i := CosName("hello")
	if i == "Cos hello" {
		t.Log("测试通过")
	}else {
		t.Error("测试失败")
	}
}
