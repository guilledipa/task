package db

import (
	"encoding/binary"
	"time"

	"github.com/coreos/bbolt"
)

var (
	taskBucket = []byte("tasks")
	db         *bbolt.DB
)

// Task is a user defined task to be saved in the Bolt DB.
type Task struct {
	Key   int
	Value string
}

// Init initializes a Bolt DB.
func Init(dbPath string) error {
	// var err error
	db, err := bbolt.Open(dbPath, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bbolt.Tx) error {
		// Setup the tasks bucket.
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// CreateTask creates a new task on the DB.
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bbolt.Tx) error {
		// Retrieve the task bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket(taskBucket)

		// Generate ID for the task.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id64, _ := b.NextSequence()
		id = int(id64)

		// Persist bytes to task bucket.
		return b.Put(itob(id), []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// AllTasks returns a list of all Tasks in the DB.
func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bbolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// DeleteTask deletes a task from the DB provided a key.
func DeleteTask(key int) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi returns an int of an 8-byte big endian representation of v.
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
