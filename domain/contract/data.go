package contract

type DataManager interface {
	RepoManager
	Close() error
}
