package contract

type DataManager interface {
	RepoManager
	SetIndexes() error
	Close() error
}
