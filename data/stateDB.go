package data


type StateDatabase struct {
	dbPath string
}

// InitLocal Initialize state on 
func (db *StateDatabase) InitLocal(targetPath string) error {
	return nil
}