package day7

type content interface {
	name() string
	size() int
	addContent(con content)
	parentDir() *directory
}

type directory struct {
	dirName          string
	directoryContent []content
	parent           *directory
}

func (d *directory) name() string {
	return d.dirName
}

func (d *directory) size() int {
	var size int
	for _, c := range d.directoryContent {
		size += c.size()
	}
	return size
}

func (d *directory) addContent(con content) {
	d.directoryContent = append(d.directoryContent, con)
}

func (d *directory) parentDir() *directory {
	return d.parent
}

type file struct {
	fileName string
	fileSize int
	parent   *directory
}

func (f *file) name() string {
	return f.fileName
}

func (f *file) size() int {
	return f.fileSize
}

func (f *file) addContent(con content) {
}

func (f *file) parentDir() *directory {
	return f.parent
}
