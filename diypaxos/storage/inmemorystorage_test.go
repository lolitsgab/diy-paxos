package storage

import (
	"sync"
	"testing"
)

type FakeKvStoreServer struct {
	InMemoryStorage
}

func (s *FakeKvStoreServer) setStore(store *sync.Map) {
	s.store = store
}

func toSyncStore(m map[string]Value) *sync.Map {
	sm := &sync.Map{}
	for k, v := range m {
		sm.Store(k, v)
	}
	return sm
}

func getSize(m *sync.Map) int {
	length := 0
	m.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	return length
}

func TestInsert(t *testing.T) {
	for _, tc := range []struct {
		name    string
		key     string
		val     int
		store   *sync.Map
		wantErr bool
	}{
		{
			name:  "Insert new",
			key:   "new",
			val:   1,
			store: toSyncStore(map[string]Value{}),
		}, {
			name:    "Insert existing",
			key:     "exist",
			val:     1,
			store:   toSyncStore(map[string]Value{"exist": Value{Value: 1}}),
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
		store   *sync.Map
		wantErr bool
	}{
		{
			name:    "Get DNE",
			key:     "new",
			wantErr: true,
			store:   toSyncStore(map[string]Value{}),
		}, {
			name:  "Get Existing",
			key:   "exist",
			val:   1,
			store: toSyncStore(map[string]Value{"exist": Value{Value: 1}}),
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
		store   *sync.Map
		wantErr bool
	}{
		{
			name:    "Remove DNE",
			key:     "dne",
			store:   toSyncStore(map[string]Value{}),
			wantErr: true,
		}, {
			name:  "Remove Existing",
			key:   "exist",
			store: toSyncStore(map[string]Value{"exist": Value{Value: 1}}),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			storage := &FakeKvStoreServer{}
			storage.setStore(tc.store)
			oldSize := getSize(tc.store)
			if err := storage.Remove(tc.key); (err != nil && !tc.wantErr) || (err == nil && tc.wantErr) {
				t.Errorf("got error %v; wanted err? %v", err.Error(), tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if want, got := oldSize-1, getSize(tc.store); want != got {
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
		store   *sync.Map
		wantErr bool
	}{
		{
			name:    "Update DNE",
			key:     "dne",
			store:   toSyncStore(map[string]Value{}),
			wantErr: true,
		}, {
			name:  "Update Existing",
			key:   "exist",
			new:   2,
			store: toSyncStore(map[string]Value{"exist": Value{Value: 1}}),
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
			if val, _ := tc.store.Load(tc.key); val.(Value).Value != int32(tc.new) {
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
		store   *sync.Map
		wantErr bool
	}{
		{
			name:  "Update DNE",
			key:   "dne",
			store: toSyncStore(map[string]Value{}),
		}, {
			name:  "Update Existing",
			key:   "exist",
			new:   2,
			store: toSyncStore(map[string]Value{"exist": {Value: 1, Meta: Metadata{Version: 1}}}),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			storage := &FakeKvStoreServer{}
			storage.setStore(tc.store)
			oldVersion := int32(0)
			v, ok := tc.store.Load(tc.key)
			if ok {
				oldVersion = v.(Value).Meta.Version
			}
			if err := storage.Upsert(tc.key, int32(tc.new)); (err != nil && !tc.wantErr) || (err == nil && tc.wantErr) {
				t.Errorf("got error %v; wanted err? %v", err.Error(), tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if val, _ := tc.store.Load(tc.key); val.(Value).Value != int32(tc.new) {
				t.Errorf("got value %v, want %v", val, tc.new)
			} else if gotVer, wantVer := val.(Value).Meta.Version, oldVersion+1; gotVer != wantVer {
				t.Errorf("got version %v, want %v", gotVer, wantVer)
			}
		})
	}
}
