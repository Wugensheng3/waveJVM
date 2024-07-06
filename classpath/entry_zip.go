package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

/**
 * @author DingZhenJie
 * @date 2024/7/3
 * @desc: 解析压缩包内的class文件相关的方法以及struct,注java的压缩包即为jar包,zip文件也行
 ** If the project has helped you, please make sure to give me a star
 * */

type ZipEntry struct {
	absPath string
	r       *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath, nil}
}
func (self *ZipEntry) openJar() error {
	r, err := zip.OpenReader(self.absPath)
	if err == nil {
		self.r = r
	}
	return err
}
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {

	if self.r == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}
	for _, f := range self.r.File {
		if f.Name == className { //优化把对应的r,都缓存起来,防止同一个包内加载多次class文件
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := io.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}

	}
	return nil, nil, errors.New("Class not found: " + className)
}
func (self *ZipEntry) String() string {
	return self.absPath
}
