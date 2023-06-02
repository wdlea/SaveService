package db

import (
	"encoding/binary"
	"os"
	"sync"

	binary_marshaler "github.com/kelindar/binary"
)

type IDBRow interface {
	GetID() uint64
}

type DBRow[RowType IDBRow] struct {
	Lock *sync.Mutex
	data RowType
}

type DB[RowType IDBRow] struct {
	entries map[uint64]*DBRow[RowType]
}

func MakeDB[RowType IDBRow]() *DB[RowType] {
	return &DB[RowType]{
		entries: make(map[uint64]*DBRow[RowType], 0),
	}
}

func (db *DB[RowType]) Set(value RowType) {
	id := value.GetID()

	entry, entry_present := db.entries[id]
	if !entry_present {
		entry = &DBRow[RowType]{
			Lock: &sync.Mutex{},
			data: value,
		}
		db.entries[id] = entry
	}

	entry.Lock.Lock()
	defer entry.Lock.Unlock()
	entry.data = value
}
func (db *DB[RowType]) Get(id uint64) (present bool, value RowType) {
	row, present := db.entries[id]
	if present {
		row.Lock.Lock()
		defer row.Lock.Unlock()

		value = row.data
	}
	return
}

func (d *DB[RowType]) Dump(file string) {
	append_stream := make(chan *[]byte, 64)
	defer close(append_stream)
	go writer(append_stream, file)

	wg := sync.WaitGroup{}

	for _, row := range d.entries {
		wg.Add(1)
		go row.DumpRow(append_stream, &wg)
	}

	wg.Wait()
}

func writer(stream chan *[]byte, file string) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err != nil {
		panic(err)
	}

	for {
		to_append, ok := <-stream
		if !ok {
			break
		}
		f.Write(*to_append)
	}
}

func (r *DBRow[RowType]) DumpRow(append_stream chan *[]byte, wg *sync.WaitGroup) {
	defer wg.Done()

	r.Lock.Lock()
	defer r.Lock.Unlock()

	id := r.data.GetID()

	//convert to bytes
	//sub-optimal at best

	line, err := binary_marshaler.Marshal(r.data)
	if err != nil {
		panic(err.Error())
	}

	idBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(idBytes, id)

	line = append(
		idBytes,
		line...,
	)

	append_stream <- &line
}
