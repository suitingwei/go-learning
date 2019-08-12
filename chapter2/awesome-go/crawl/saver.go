package crawl

type Saver interface {
	//save the awesome go project into the storage, whether the database or the filesystem
	Save(ag *AwesomeGoData) error
}

type DatabaseSaver struct {
}

func (saver *DatabaseSaver) Save(ag *AwesomeGoData) error {
	return nil
}

type FileSaver struct {
	path string //path to save the data
}

func (saver *FileSaver) Save(ag *AwesomeGoData) error {

	return nil
}

func NewFileSaver(savePath string) Saver {

	saver := &FileSaver{
		path: savePath,
	}
	return saver
}
