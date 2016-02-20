package database

import (
	"log"

	"github.com/boltdb/bolt"
)

var bucket = []byte("bucket")

//GetDB returns a bolt.db instance
func GetDB(dbFilePath string) *bolt.DB {
	db, err := bolt.Open(dbFilePath, 0644, nil)
	if err != nil {
		log.Fatalf("Error while opening DB %v", err)
	}
	return db
}

//GetValue returns value for a given key
func GetValue(db *bolt.DB, key []byte) []byte {
	var val []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)
		if bucket == nil {
			return nil
		}

		val = bucket.Get(key)
		return nil
	})
	if err != nil {
		log.Fatalf("Error while getting value %s from DB %v", key, err)
	}
	return val
}

//StoreData strores key-value pairs in DB
func StoreData(db *bolt.DB, key []byte, value []byte) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			log.Fatalf("Error creatign bucket in DB %v", err)
		}

		err = bucket.Put(key, value)
		if err != nil {
			log.Fatalf("Error while storing key-value pair [%s-%s] from DB %v", key, value, err)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error while updating DB with key-value pair [%s-%s]  %v", key, value, err)
	}
}
