package db

import (
	"database/sql"
	"github.com/textileio/textile-go/repo"
	"sync"
	"testing"
	"time"
)

var storeReqDB repo.CafeStoreRequestStore

func init() {
	setupCafeStoreRequestDB()
}

func setupCafeStoreRequestDB() {
	conn, _ := sql.Open("sqlite3", ":memory:")
	initDatabaseTables(conn, "")
	storeReqDB = NewCafeStoreRequestStore(conn, new(sync.Mutex))
}

func TestCafeStoreRequestDB_Put(t *testing.T) {
	err := storeReqDB.Put(&repo.CafeStoreRequest{
		Id:     "abcde",
		CafeId: "boom",
		Date:   time.Now(),
	})
	if err != nil {
		t.Error(err)
	}
	stmt, err := storeReqDB.PrepareQuery("select id from storereqs where id=?")
	defer stmt.Close()
	var id string
	err = stmt.QueryRow("abcde").Scan(&id)
	if err != nil {
		t.Error(err)
	}
	if id != "abcde" {
		t.Errorf(`expected "abcde" got %s`, id)
	}
}

func TestCafeStoreRequestDB_List(t *testing.T) {
	setupCafeStoreRequestDB()
	err := storeReqDB.Put(&repo.CafeStoreRequest{
		Id:     "abcde",
		CafeId: "boom",
		Date:   time.Now(),
	})
	if err != nil {
		t.Error(err)
	}
	err = storeReqDB.Put(&repo.CafeStoreRequest{
		Id:     "abcdef",
		CafeId: "boom",
		Date:   time.Now().Add(time.Minute),
	})
	if err != nil {
		t.Error(err)
	}
	all := storeReqDB.List("", -1)
	if len(all) != 2 {
		t.Error("returned incorrect number of store requests")
		return
	}
	limited := storeReqDB.List("", 1)
	if len(limited) != 1 {
		t.Error("returned incorrect number of store requests")
		return
	}
	offset := storeReqDB.List(limited[0].Id, -1)
	if len(offset) != 1 {
		t.Error("returned incorrect number of store requests")
		return
	}
}

func TestCafeStoreRequestDB_Delete(t *testing.T) {
	err := storeReqDB.Delete("abcde")
	if err != nil {
		t.Error(err)
	}
	stmt, err := storeReqDB.PrepareQuery("select id from storereqs where id=?")
	defer stmt.Close()
	var id string
	if err := stmt.QueryRow("abcde").Scan(&id); err == nil {
		t.Error("delete failed")
	}
}

func TestCafeStoreRequestDB_DeleteByCafe(t *testing.T) {
	setupCafeStoreRequestDB()
	err := storeReqDB.Put(&repo.CafeStoreRequest{
		Id:     "xyz",
		CafeId: "boom",
		Date:   time.Now(),
	})
	if err != nil {
		t.Error(err)
	}
	err = storeReqDB.DeleteByCafe("boom")
	if err != nil {
		t.Error(err)
	}
	stmt, err := storeReqDB.PrepareQuery("select id from storereqs where id=?")
	defer stmt.Close()
	var id string
	if err := stmt.QueryRow("zyx").Scan(&id); err == nil {
		t.Error("delete by cafe failed")
	}
}
