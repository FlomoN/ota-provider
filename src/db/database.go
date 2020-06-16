package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Watcher is a Repository/Release that should be monitored
type Watcher struct {
	Repo      string
	Device    string
	ETag      string
	ReleaseID int
	Version   string
}

// Store saves the data object and stores additional information like github login and the list of monitored Releases
type Store struct {
	Init     bool
	GhName   string
	GhToken  string
	MQTTHost string
	Watch    []Watcher
}

// Load is loading the Store from json file
func Load() *Store {

	var store Store

	pwd, _ := os.Getwd()

	f, err := ioutil.ReadFile(filepath.Join(pwd, "./data/store.json"))
	if os.IsNotExist(err) {
		os.Mkdir(filepath.Join(pwd, "./data"), 0777)
		_f, _ := os.Create(filepath.Join(pwd, "./data/store.json"))
		store = Store{false, "", "", "", []Watcher{}}
		x, _ := json.Marshal(store)
		_f.Write(x)
		_f.Close()
	} else {
		err = json.Unmarshal(f, &store)
	}

	return &store
}

// Add is adding an Element to the Watchlist
func (s *Store) Add(w Watcher) {
	s.Watch = append(s.Watch, w)
	s.saveData()
}

// SetCreds sets the github user data and takes Store out of uninitialized state
func (s *Store) SetCreds(user string, token string) {
	s.GhName = user
	s.GhToken = token
	s.Init = true
	s.saveData()
}

func (s *Store) saveData() {
	pwd, _ := os.Getwd()
	x, _ := json.Marshal(s)
	ioutil.WriteFile(filepath.Join(pwd, "./data/store.json"), x, 0777)
}

//UpdateEIV updates etag, releaseid and version of the Watcher
func (w *Watcher) UpdateEIV(s *Store, etag string, id int, version string) {
	w.ETag = etag
	w.ReleaseID = id
	w.Version = version
	s.saveData()
}
