package repository

import (
	"reflect"
	"shop/models"
	"testing"
)

func TestNewMapDB(t *testing.T) {
	mapDB, ok := NewMapDB().(*mapDB)
	if !ok {
		t.Error("Can't open DB")
	}
	if mapDB.itemsTable == nil || mapDB.ordersTable == nil {
		t.Error("Can't open DB")
	}
}

func TestCreateItem(t *testing.T) {
	
}
