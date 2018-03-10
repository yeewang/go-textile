package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"sync"

	"github.com/textileio/mill-go/repo"
)

var SettingsNotSetError error = errors.New("Settings not set")

type SettingsDB struct {
	modelStore
}

func NewConfigurationStore(db *sql.DB, lock *sync.Mutex) repo.ConfigurationStore {
	return &SettingsDB{modelStore{db, lock}}
}

func (s *SettingsDB) Put(settings repo.SettingsData) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(&settings, "", "    ")
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert or replace into config(key, value) values(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec("settings", string(b))
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *SettingsDB) Get() (repo.SettingsData, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	settings := repo.SettingsData{}
	stmt, err := s.db.Prepare("select value from config where key=?")
	if err != nil {
		return settings, err
	}
	defer stmt.Close()
	var settingsBytes []byte
	err = stmt.QueryRow("settings").Scan(&settingsBytes)
	if err != nil {
		return settings, SettingsNotSetError
	}
	err = json.Unmarshal(settingsBytes, &settings)
	if err != nil {
		return settings, err
	}
	return settings, nil
}

func (s *SettingsDB) Update(settings repo.SettingsData) error {
	current, err := s.Get()
	if err != nil {
		return errors.New("Not Found")
	}
	if settings.Version == nil {
		settings.Version = current.Version
	}
	err = s.Put(settings)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes all settings from the database. It's a destructive action that should be used with care.
func (s *SettingsDB) Delete() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	stmt, err := s.db.Prepare("delete from config where key = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec("settings")

	return err
}
