package monolith

import (
	"github.com/golang/glog"
	"github.com/syndtr/goleveldb/leveldb"
	"net/url"
)

var db *leveldb.DB

func Init(dbFilePath string) error {
	dbl, err := leveldb.OpenFile(dbFilePath, nil)
	db = dbl
	return err
}

func Close() {
	db.Close()
}

// Obtains redirection target, merges pathes and redirects.
// Returns:
// location and nil if found,
// redirect key and ErrKeyNotFound if not found,
// nil and ErrUnexpected if other error happens.
func ResolveLocation(key string) (*string, error) {
	data, err := db.Get([]byte(key), nil)

	if err == leveldb.ErrNotFound {
		glog.V(1).Info("Not found: ", key)
		return &key, ErrKeyNotFound
	} else if err != nil {
		glog.Error(err)
		return nil, NewErrUnexpected(err)
	}

	location := string(data)
	glog.V(2).Info("Found: ", location)
	return &location, nil
}

func Add(key string, url *url.URL) error {
	return db.Put([]byte(key), []byte(url.String()), nil)
}
