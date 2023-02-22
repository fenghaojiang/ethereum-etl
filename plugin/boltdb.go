package plugin

import "github.com/boltdb/bolt"

func BoltDB(dbName string) (*bolt.DB, error) {
	_db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		return nil, err
	}
	return _db, err
}
