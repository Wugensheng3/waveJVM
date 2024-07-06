package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/**
 * @author DingZhenJie
 * @date 2024/7/3
 * @desc: 通配符对应的加载Entry
 ** If the project has helped you, please make sure to give me a star
 * */
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //意义是删掉*
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
