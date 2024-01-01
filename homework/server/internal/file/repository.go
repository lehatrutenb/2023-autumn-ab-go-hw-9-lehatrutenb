package file

type Repository interface {
	AddFile(File) error
	GetFileByName(string) (*File, error)
	GetFileNames() ([]string, error)
}
