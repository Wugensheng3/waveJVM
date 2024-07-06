package classpath

/**
 * @author DingZhenJie
 * @date 2024/7/3
 * @desc: 表示类路径的接口,底下还会有其他实现,如DirEntry等等,功能主要是为了加载.class文件
 ** If the project has helped you, please make sure to give me a star
 * */
import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	/**
	负责寻找和加载class文件
	todo: className:相对路径(为什么要相对路径?)
	*/
	readClass(className string) ([]byte, Entry, error)
	/**
	返回变量的字符串表示
	*/
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)

}
