package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// 重定向标准输出以捕获输出
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 执行main函数
	main()

	// 关闭写入端并恢复标准输出
	w.Close()
	os.Stdout = old

	// 读取捕获的输出
	var buf [1024]byte
	n, _ := r.Read(buf[:])
	output := string(buf[:n])

	// 检查输出是否为“Hello, World!\n”
	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}

	// 检查文件是否成功打开
	_, err := os.Open("test")
	if err != nil {
		t.Errorf("failed to open file: %v", err)
	}
}
