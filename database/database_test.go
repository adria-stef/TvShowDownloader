package database_test

import (
	"os"

	. "github.com/adria-stef/TvShowDownloader/database"
	"github.com/boltdb/bolt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var db *bolt.DB

var _ = BeforeSuite(func() {
	db = GetDB("./bolt.db")
})

var _ = AfterSuite(func() {
	db.Close()

	err := os.Remove("./bolt.db")
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("Database", func() {

	key := []byte("key")
	value := []byte("value")
	updatedValue := []byte("updatedValue")

	Describe("StoreData", func() {
		Context("Store new key", func() {

			It("should succeed", func() {
				StoreData(db, key, value)
				Expect(GetValue(db, key)).To(Equal(value))
			})
		})

		Context("Udate existing key", func() {
			It("should be updated", func() {

				StoreData(db, key, value)
				StoreData(db, key, updatedValue)

				Expect(GetValue(db, key)).To(Equal(updatedValue))
			})
		})
	})
})
