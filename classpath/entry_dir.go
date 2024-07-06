package classpath

import (
	"os"
	"path/filepath"
)

/**
 * @author DingZhenJie
 * @date 2024/7/3
 * @desc: todo
 ** If the project has helped you, please make sure to give me a star
 * */

type DirEntry struct {
	//存放绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	//将给定的路径转换为绝对路径。具体来说，它会解析 path，并返回该路径的绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// todo:怎么获取对应的各个className
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := os.ReadFile(fileName)
	return data, self, err
}
func (self *DirEntry) String() string {
	return self.absDir
}
