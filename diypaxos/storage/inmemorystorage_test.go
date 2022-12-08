package storage

import "testing"

type FakeKvStoreServer struct {
	InMemoryStorage
}

func (s *FakeKvStoreServer) setStore(store map[string]Value) {
	s.store = store
}

func TestInsert(t *testing.T) {
	for _, tc := range []struct {
		name    string
		key     string
		val     int
		store   map[string]Value
		wantErr bool
	}{
		{
			name:  "Insert new",
			key:   "new",
			val:   1,
			store: map[string]Value{},
		}, {
			name:    "Insert existing",
			key:     "exist",
			val:     1,
			store:   map[string]Value{"exist": Value{Value: 1}},
			wantErr: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			storage := &FakeKvStoreServer{}
			storage.setStore(tc.store)
			if err := storage.Insert(tc.key, int32(tc.val)); (err != nil && !tc.wantErr) || (err == nil && tc.wantErr) {
				t.Errorf("got error %v; wanted err? %v", err.Error(), tc.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	for _, tc := range []struct {
		name    string
		key     string
		val     int
		store   map[string]Value
		wantErr bool
	}{
		{
			name:    "Get DNE",
			key:     "new",
			wantErr: true,
			store:   map[string]Value{},
		}, {
			name:  "Get Existing",
			key:   "exist",
			val:   1,
			store: map[string]Value{"exist": Value{Value: 1}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			storage := &FakeKvStoreServer{}
			storage.setStore(tc.store)
			if got, err := storage.Get(tc.key); (err != nil && !tc.wantErr) || (err == nil && tc.wantErr) {
				t.Errorf("got error %v; wanted err? %v", err.Error(), tc.wantErr)
			} else if err == nil && got != int32(tc.val) {
				t.Errorf("got %v, expected %v", got, tc.val)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	for _, tc := range []struct {
		name    string
		key     string
		store   map[string]Value
		wantErr bool
	}{
		{
			name:    "Remove DNE",
			key:     "dne",
			store:   map[string]Value{},
			wantErr: true,
		}, {
			name:  "Remove Existing",
			key:   "exist",
			store: map[string]Value{"exist": Value{Value: 1}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			storage := &FakeKvStoreServer{}
			storage.setStore(tc.store)
			oldSize := len(tc.store)
			if err := storage.Remove(tc.key); (err != nil && !tc.wantErr) || (err == nil && tc.wantErr) {
				t.Errorf("got error %v; wanted err? %v", err.Error(), tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if want, got := oldSize-1, len(tc.store); want != got {
				t.Errorf("len(store)=%v, wanted %v", want, got)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	for _, tc := range []struct {
		name    string
		key     string
		new     int
		store   map[string]Value
		wantErr bool
	}{
		{
			name:    "Update DNE",
			key:     "dne",
			store:   map[string]Value{},
			wantErr: true,
		}, {
			name:  "Update Existing",
			key:   "exist",
			new:   2,
			store: map[string]Value{"exist": Value{Value: 1}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			storage := &FakeKvStoreServer{}
			storage.setStore(tc.store)
			if err := storage.Update(tc.key, int32(tc.new)); (err != nil && !tc.wantErr) || (err == nil && tc.wantErr) {
				t.Errorf("got error %v; wanted err? %v", err.Error(), tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if val, _ := tc.store[tc.key]; val.Value != int32(tc.new) {
				t.Errorf("got %v, want %v", val, tc.new)
			}
		})
	}
}

func TestUpsert(t *testing.T) {
	for _, tc := range []struct {
		name    string
		key     string
		new     int
		store   map[string]Value
		wantErr bool
	}{
		{
			name:  "Update DNE",
			key:   "dne",
			store: map[string]Value{},
		}, {
			name:  "Update Existing",
			key:   "exist",
			new:   2,
			store: map[string]Value{"exist": {Value: 1}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			storage := &FakeKvStoreServer{}
			storage.setStore(tc.store)
			if err := storage.Upsert(tc.key, int32(tc.new)); (err != nil && !tc.wantErr) || (err == nil && tc.wantErr) {
				t.Errorf("got error %v; wanted err? %v", err.Error(), tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if val, _ := tc.store[tc.key]; val.Value != int32(tc.new) {
				t.Errorf("got %v, want %v", val, tc.new)
			}
		})
	}
}
