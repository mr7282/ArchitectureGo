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
	MyDB := &mapDB{
		itemsTable: &itemsTable{
			items: map[int32]*models.Item{
				1: {
					ID: 1,
					Name: "TestItem_1",
					Price: 10.0,
				},
				2: {
					ID: 2,
					Name: "TestItem_2",
					Price: 20.0,
				},
				3: {
					ID: 3,
					Name: "TestItem_3",
					Price: 30.0,
				},
			},
			maxID: 3,
		},
	}

	newItem := &models.Item{
		ID: 4,
		Name: "Want Test 4",
		Price: 40.0,
	}

	want, err := MyDB.CreateItem(newItem)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, newItem) {
		t.Fatalf("want %v get %v", newItem, want)
	}
	if MyDB.itemsTable.maxID != 4 {
		t.Fatalf("want 4 get %v", MyDB.itemsTable.maxID)
	}
}

func TestGetItem(t *testing.T) {
	MyDB := &mapDB{
		itemsTable: &itemsTable{
			items: map[int32]*models.Item{
				1: {
					ID: 1,
					Name: "TestItem_1",
					Price: 10.0,
				},
				2: {
					ID: 2,
					Name: "TestItem_2",
					Price: 20.0,
				},
				3: {
					ID: 3,
					Name: "TestItem_3",
					Price: 30.0,
				},
			},
			maxID: 3,
		},
	}

	getTest, err := MyDB.GetItem(3)
	if err != nil{
		t.Fatal(err)
	}
	if !reflect.DeepEqual(getTest, MyDB.itemsTable.items[3]){
		t.Fatalf("want %v get %v", MyDB.itemsTable.items[3], getTest)
	}
}

func TestDeletItem(t *testing.T) {
	MyDB := &mapDB{
		itemsTable: &itemsTable{
			items: map[int32]*models.Item{
				1: {
					ID: 1,
					Name: "TestItem_1",
					Price: 10.0,
				},
				2: {
					ID: 2,
					Name: "TestItem_2",
					Price: 20.0,
				},
				3: {
					ID: 3,
					Name: "TestItem_3",
					Price: 30.0,
				},
			},
			maxID: 3,
		},
	}

	MyDB.DeleteItem(2)

	if len(MyDB.itemsTable.items) != 2 {
		t.Fatalf("Function DeletItem is not work")
	}
}

func TestUpdateItem(t *testing.T) {
	MyDB := &mapDB{
		itemsTable: &itemsTable{
			items: map[int32]*models.Item{
				1: {
					ID: 1,
					Name: "TestItem_1",
					Price: 10.0,
				},
				2: {
					ID: 2,
					Name: "TestItem_2",
					Price: 20.0,
				},
				3: {
					ID: 3,
					Name: "TestItem_3",
					Price: 30.0,
				},
			},
			maxID: 3,
		},
	}

	want := &models.Item{
		ID: 2,
		Name: "Update_Test_Ok",
		Price: 0.2,
	}

	_, err := MyDB.UpdateItem(want)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, MyDB.itemsTable.items[2]) {
		t.Fatalf("want %v got %v", want, MyDB.itemsTable.items[2])
	}
}
