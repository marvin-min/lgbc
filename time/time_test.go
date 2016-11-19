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


func BenchmarkCosName(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i:= 0;i < b.N;i ++{
		CosName("啊大法师发沙发上发sf安师大发生发顺丰")
	}
}