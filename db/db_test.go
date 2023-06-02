package db

import "testing"

func TestDB(t *testing.T) {
	d := MakeDB[Row]()

	d.Set(Row{
		id:   17,
		Val:  6969,
		Val1: 1000,
	})
	// d.Set(Row{
	// 	id:  69,
	// 	Val: false,
	// })

	ok, val := d.Get(17)
	if !ok {
		t.Fatalf("Not ok when should be ok")
	}
	if val.Val != 6969 {
		t.Fatalf("Incorrect value returned")
	}

	d.Dump("C:\\Users\\wdlea\\go\\src\\github.com\\wdlea\\SaveSystem\\db\\balls.txt")
}

type Row struct {
	id   uint64
	Val  uint64
	Val1 uint64
}

func (r Row) GetID() uint64 {
	return r.id
}
