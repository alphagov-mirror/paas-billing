// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"database/sql"
	"io"
	"sync"

	cf "github.com/alphagov/paas-billing/cloudfoundry"
	"github.com/alphagov/paas-billing/db"
	"github.com/compose/gocomposeapi"
	_ "github.com/lib/pq"
)

type FakeSQLClient struct {
	InitSchemaStub        func() error
	initSchemaMutex       sync.RWMutex
	initSchemaArgsForCall []struct{}
	initSchemaReturns     struct {
		result1 error
	}
	initSchemaReturnsOnCall map[int]struct {
		result1 error
	}
	InsertUsageEventListStub        func(data *cf.UsageEventList, tableName string) error
	insertUsageEventListMutex       sync.RWMutex
	insertUsageEventListArgsForCall []struct {
		data      *cf.UsageEventList
		tableName string
	}
	insertUsageEventListReturns struct {
		result1 error
	}
	insertUsageEventListReturnsOnCall map[int]struct {
		result1 error
	}
	FetchLastGUIDStub        func(tableName string) (string, error)
	fetchLastGUIDMutex       sync.RWMutex
	fetchLastGUIDArgsForCall []struct {
		tableName string
	}
	fetchLastGUIDReturns struct {
		result1 string
		result2 error
	}
	fetchLastGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ExecStub        func(query string, args ...interface{}) (sql.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		query string
		args  []interface{}
	}
	execReturns struct {
		result1 sql.Result
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	PrepareStub        func(query string) (*sql.Stmt, error)
	prepareMutex       sync.RWMutex
	prepareArgsForCall []struct {
		query string
	}
	prepareReturns struct {
		result1 *sql.Stmt
		result2 error
	}
	prepareReturnsOnCall map[int]struct {
		result1 *sql.Stmt
		result2 error
	}
	QueryJSONStub        func(q string, args ...interface{}) io.Reader
	queryJSONMutex       sync.RWMutex
	queryJSONArgsForCall []struct {
		q    string
		args []interface{}
	}
	queryJSONReturns struct {
		result1 io.Reader
	}
	queryJSONReturnsOnCall map[int]struct {
		result1 io.Reader
	}
	QueryRowJSONStub        func(q string, args ...interface{}) io.Reader
	queryRowJSONMutex       sync.RWMutex
	queryRowJSONArgsForCall []struct {
		q    string
		args []interface{}
	}
	queryRowJSONReturns struct {
		result1 io.Reader
	}
	queryRowJSONReturnsOnCall map[int]struct {
		result1 io.Reader
	}
	UpdateViewsStub        func() error
	updateViewsMutex       sync.RWMutex
	updateViewsArgsForCall []struct{}
	updateViewsReturns     struct {
		result1 error
	}
	updateViewsReturnsOnCall map[int]struct {
		result1 error
	}
	BeginTxStub        func() (db.SQLClient, error)
	beginTxMutex       sync.RWMutex
	beginTxArgsForCall []struct{}
	beginTxReturns     struct {
		result1 db.SQLClient
		result2 error
	}
	beginTxReturnsOnCall map[int]struct {
		result1 db.SQLClient
		result2 error
	}
	RollbackStub        func() error
	rollbackMutex       sync.RWMutex
	rollbackArgsForCall []struct{}
	rollbackReturns     struct {
		result1 error
	}
	rollbackReturnsOnCall map[int]struct {
		result1 error
	}
	CommitStub        func() error
	commitMutex       sync.RWMutex
	commitArgsForCall []struct{}
	commitReturns     struct {
		result1 error
	}
	commitReturnsOnCall map[int]struct {
		result1 error
	}
	InsertComposeCursorStub        func(*string) error
	insertComposeCursorMutex       sync.RWMutex
	insertComposeCursorArgsForCall []struct {
		arg1 *string
	}
	insertComposeCursorReturns struct {
		result1 error
	}
	insertComposeCursorReturnsOnCall map[int]struct {
		result1 error
	}
	FetchComposeCursorStub        func() (*string, error)
	fetchComposeCursorMutex       sync.RWMutex
	fetchComposeCursorArgsForCall []struct{}
	fetchComposeCursorReturns     struct {
		result1 *string
		result2 error
	}
	fetchComposeCursorReturnsOnCall map[int]struct {
		result1 *string
		result2 error
	}
	InsertComposeLatestEventIDStub        func(string) error
	insertComposeLatestEventIDMutex       sync.RWMutex
	insertComposeLatestEventIDArgsForCall []struct {
		arg1 string
	}
	insertComposeLatestEventIDReturns struct {
		result1 error
	}
	insertComposeLatestEventIDReturnsOnCall map[int]struct {
		result1 error
	}
	FetchComposeLatestEventIDStub        func() (*string, error)
	fetchComposeLatestEventIDMutex       sync.RWMutex
	fetchComposeLatestEventIDArgsForCall []struct{}
	fetchComposeLatestEventIDReturns     struct {
		result1 *string
		result2 error
	}
	fetchComposeLatestEventIDReturnsOnCall map[int]struct {
		result1 *string
		result2 error
	}
	InsertComposeAuditEventsStub        func([]composeapi.AuditEvent) error
	insertComposeAuditEventsMutex       sync.RWMutex
	insertComposeAuditEventsArgsForCall []struct {
		arg1 []composeapi.AuditEvent
	}
	insertComposeAuditEventsReturns struct {
		result1 error
	}
	insertComposeAuditEventsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSQLClient) InitSchema() error {
	fake.initSchemaMutex.Lock()
	ret, specificReturn := fake.initSchemaReturnsOnCall[len(fake.initSchemaArgsForCall)]
	fake.initSchemaArgsForCall = append(fake.initSchemaArgsForCall, struct{}{})
	fake.recordInvocation("InitSchema", []interface{}{})
	fake.initSchemaMutex.Unlock()
	if fake.InitSchemaStub != nil {
		return fake.InitSchemaStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.initSchemaReturns.result1
}

func (fake *FakeSQLClient) InitSchemaCallCount() int {
	fake.initSchemaMutex.RLock()
	defer fake.initSchemaMutex.RUnlock()
	return len(fake.initSchemaArgsForCall)
}

func (fake *FakeSQLClient) InitSchemaReturns(result1 error) {
	fake.InitSchemaStub = nil
	fake.initSchemaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) InitSchemaReturnsOnCall(i int, result1 error) {
	fake.InitSchemaStub = nil
	if fake.initSchemaReturnsOnCall == nil {
		fake.initSchemaReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initSchemaReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) InsertUsageEventList(data *cf.UsageEventList, tableName string) error {
	fake.insertUsageEventListMutex.Lock()
	ret, specificReturn := fake.insertUsageEventListReturnsOnCall[len(fake.insertUsageEventListArgsForCall)]
	fake.insertUsageEventListArgsForCall = append(fake.insertUsageEventListArgsForCall, struct {
		data      *cf.UsageEventList
		tableName string
	}{data, tableName})
	fake.recordInvocation("InsertUsageEventList", []interface{}{data, tableName})
	fake.insertUsageEventListMutex.Unlock()
	if fake.InsertUsageEventListStub != nil {
		return fake.InsertUsageEventListStub(data, tableName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.insertUsageEventListReturns.result1
}

func (fake *FakeSQLClient) InsertUsageEventListCallCount() int {
	fake.insertUsageEventListMutex.RLock()
	defer fake.insertUsageEventListMutex.RUnlock()
	return len(fake.insertUsageEventListArgsForCall)
}

func (fake *FakeSQLClient) InsertUsageEventListArgsForCall(i int) (*cf.UsageEventList, string) {
	fake.insertUsageEventListMutex.RLock()
	defer fake.insertUsageEventListMutex.RUnlock()
	return fake.insertUsageEventListArgsForCall[i].data, fake.insertUsageEventListArgsForCall[i].tableName
}

func (fake *FakeSQLClient) InsertUsageEventListReturns(result1 error) {
	fake.InsertUsageEventListStub = nil
	fake.insertUsageEventListReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) InsertUsageEventListReturnsOnCall(i int, result1 error) {
	fake.InsertUsageEventListStub = nil
	if fake.insertUsageEventListReturnsOnCall == nil {
		fake.insertUsageEventListReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertUsageEventListReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) FetchLastGUID(tableName string) (string, error) {
	fake.fetchLastGUIDMutex.Lock()
	ret, specificReturn := fake.fetchLastGUIDReturnsOnCall[len(fake.fetchLastGUIDArgsForCall)]
	fake.fetchLastGUIDArgsForCall = append(fake.fetchLastGUIDArgsForCall, struct {
		tableName string
	}{tableName})
	fake.recordInvocation("FetchLastGUID", []interface{}{tableName})
	fake.fetchLastGUIDMutex.Unlock()
	if fake.FetchLastGUIDStub != nil {
		return fake.FetchLastGUIDStub(tableName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchLastGUIDReturns.result1, fake.fetchLastGUIDReturns.result2
}

func (fake *FakeSQLClient) FetchLastGUIDCallCount() int {
	fake.fetchLastGUIDMutex.RLock()
	defer fake.fetchLastGUIDMutex.RUnlock()
	return len(fake.fetchLastGUIDArgsForCall)
}

func (fake *FakeSQLClient) FetchLastGUIDArgsForCall(i int) string {
	fake.fetchLastGUIDMutex.RLock()
	defer fake.fetchLastGUIDMutex.RUnlock()
	return fake.fetchLastGUIDArgsForCall[i].tableName
}

func (fake *FakeSQLClient) FetchLastGUIDReturns(result1 string, result2 error) {
	fake.FetchLastGUIDStub = nil
	fake.fetchLastGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) FetchLastGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.FetchLastGUIDStub = nil
	if fake.fetchLastGUIDReturnsOnCall == nil {
		fake.fetchLastGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.fetchLastGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("Exec", []interface{}{query, args})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(query, args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.execReturns.result1, fake.execReturns.result2
}

func (fake *FakeSQLClient) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeSQLClient) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].query, fake.execArgsForCall[i].args
}

func (fake *FakeSQLClient) ExecReturns(result1 sql.Result, result2 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) ExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) Prepare(query string) (*sql.Stmt, error) {
	fake.prepareMutex.Lock()
	ret, specificReturn := fake.prepareReturnsOnCall[len(fake.prepareArgsForCall)]
	fake.prepareArgsForCall = append(fake.prepareArgsForCall, struct {
		query string
	}{query})
	fake.recordInvocation("Prepare", []interface{}{query})
	fake.prepareMutex.Unlock()
	if fake.PrepareStub != nil {
		return fake.PrepareStub(query)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.prepareReturns.result1, fake.prepareReturns.result2
}

func (fake *FakeSQLClient) PrepareCallCount() int {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return len(fake.prepareArgsForCall)
}

func (fake *FakeSQLClient) PrepareArgsForCall(i int) string {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return fake.prepareArgsForCall[i].query
}

func (fake *FakeSQLClient) PrepareReturns(result1 *sql.Stmt, result2 error) {
	fake.PrepareStub = nil
	fake.prepareReturns = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) PrepareReturnsOnCall(i int, result1 *sql.Stmt, result2 error) {
	fake.PrepareStub = nil
	if fake.prepareReturnsOnCall == nil {
		fake.prepareReturnsOnCall = make(map[int]struct {
			result1 *sql.Stmt
			result2 error
		})
	}
	fake.prepareReturnsOnCall[i] = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) QueryJSON(q string, args ...interface{}) io.Reader {
	fake.queryJSONMutex.Lock()
	ret, specificReturn := fake.queryJSONReturnsOnCall[len(fake.queryJSONArgsForCall)]
	fake.queryJSONArgsForCall = append(fake.queryJSONArgsForCall, struct {
		q    string
		args []interface{}
	}{q, args})
	fake.recordInvocation("QueryJSON", []interface{}{q, args})
	fake.queryJSONMutex.Unlock()
	if fake.QueryJSONStub != nil {
		return fake.QueryJSONStub(q, args...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.queryJSONReturns.result1
}

func (fake *FakeSQLClient) QueryJSONCallCount() int {
	fake.queryJSONMutex.RLock()
	defer fake.queryJSONMutex.RUnlock()
	return len(fake.queryJSONArgsForCall)
}

func (fake *FakeSQLClient) QueryJSONArgsForCall(i int) (string, []interface{}) {
	fake.queryJSONMutex.RLock()
	defer fake.queryJSONMutex.RUnlock()
	return fake.queryJSONArgsForCall[i].q, fake.queryJSONArgsForCall[i].args
}

func (fake *FakeSQLClient) QueryJSONReturns(result1 io.Reader) {
	fake.QueryJSONStub = nil
	fake.queryJSONReturns = struct {
		result1 io.Reader
	}{result1}
}

func (fake *FakeSQLClient) QueryJSONReturnsOnCall(i int, result1 io.Reader) {
	fake.QueryJSONStub = nil
	if fake.queryJSONReturnsOnCall == nil {
		fake.queryJSONReturnsOnCall = make(map[int]struct {
			result1 io.Reader
		})
	}
	fake.queryJSONReturnsOnCall[i] = struct {
		result1 io.Reader
	}{result1}
}

func (fake *FakeSQLClient) QueryRowJSON(q string, args ...interface{}) io.Reader {
	fake.queryRowJSONMutex.Lock()
	ret, specificReturn := fake.queryRowJSONReturnsOnCall[len(fake.queryRowJSONArgsForCall)]
	fake.queryRowJSONArgsForCall = append(fake.queryRowJSONArgsForCall, struct {
		q    string
		args []interface{}
	}{q, args})
	fake.recordInvocation("QueryRowJSON", []interface{}{q, args})
	fake.queryRowJSONMutex.Unlock()
	if fake.QueryRowJSONStub != nil {
		return fake.QueryRowJSONStub(q, args...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.queryRowJSONReturns.result1
}

func (fake *FakeSQLClient) QueryRowJSONCallCount() int {
	fake.queryRowJSONMutex.RLock()
	defer fake.queryRowJSONMutex.RUnlock()
	return len(fake.queryRowJSONArgsForCall)
}

func (fake *FakeSQLClient) QueryRowJSONArgsForCall(i int) (string, []interface{}) {
	fake.queryRowJSONMutex.RLock()
	defer fake.queryRowJSONMutex.RUnlock()
	return fake.queryRowJSONArgsForCall[i].q, fake.queryRowJSONArgsForCall[i].args
}

func (fake *FakeSQLClient) QueryRowJSONReturns(result1 io.Reader) {
	fake.QueryRowJSONStub = nil
	fake.queryRowJSONReturns = struct {
		result1 io.Reader
	}{result1}
}

func (fake *FakeSQLClient) QueryRowJSONReturnsOnCall(i int, result1 io.Reader) {
	fake.QueryRowJSONStub = nil
	if fake.queryRowJSONReturnsOnCall == nil {
		fake.queryRowJSONReturnsOnCall = make(map[int]struct {
			result1 io.Reader
		})
	}
	fake.queryRowJSONReturnsOnCall[i] = struct {
		result1 io.Reader
	}{result1}
}

func (fake *FakeSQLClient) UpdateViews() error {
	fake.updateViewsMutex.Lock()
	ret, specificReturn := fake.updateViewsReturnsOnCall[len(fake.updateViewsArgsForCall)]
	fake.updateViewsArgsForCall = append(fake.updateViewsArgsForCall, struct{}{})
	fake.recordInvocation("UpdateViews", []interface{}{})
	fake.updateViewsMutex.Unlock()
	if fake.UpdateViewsStub != nil {
		return fake.UpdateViewsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateViewsReturns.result1
}

func (fake *FakeSQLClient) UpdateViewsCallCount() int {
	fake.updateViewsMutex.RLock()
	defer fake.updateViewsMutex.RUnlock()
	return len(fake.updateViewsArgsForCall)
}

func (fake *FakeSQLClient) UpdateViewsReturns(result1 error) {
	fake.UpdateViewsStub = nil
	fake.updateViewsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) UpdateViewsReturnsOnCall(i int, result1 error) {
	fake.UpdateViewsStub = nil
	if fake.updateViewsReturnsOnCall == nil {
		fake.updateViewsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateViewsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) BeginTx() (db.SQLClient, error) {
	fake.beginTxMutex.Lock()
	ret, specificReturn := fake.beginTxReturnsOnCall[len(fake.beginTxArgsForCall)]
	fake.beginTxArgsForCall = append(fake.beginTxArgsForCall, struct{}{})
	fake.recordInvocation("BeginTx", []interface{}{})
	fake.beginTxMutex.Unlock()
	if fake.BeginTxStub != nil {
		return fake.BeginTxStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.beginTxReturns.result1, fake.beginTxReturns.result2
}

func (fake *FakeSQLClient) BeginTxCallCount() int {
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	return len(fake.beginTxArgsForCall)
}

func (fake *FakeSQLClient) BeginTxReturns(result1 db.SQLClient, result2 error) {
	fake.BeginTxStub = nil
	fake.beginTxReturns = struct {
		result1 db.SQLClient
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) BeginTxReturnsOnCall(i int, result1 db.SQLClient, result2 error) {
	fake.BeginTxStub = nil
	if fake.beginTxReturnsOnCall == nil {
		fake.beginTxReturnsOnCall = make(map[int]struct {
			result1 db.SQLClient
			result2 error
		})
	}
	fake.beginTxReturnsOnCall[i] = struct {
		result1 db.SQLClient
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) Rollback() error {
	fake.rollbackMutex.Lock()
	ret, specificReturn := fake.rollbackReturnsOnCall[len(fake.rollbackArgsForCall)]
	fake.rollbackArgsForCall = append(fake.rollbackArgsForCall, struct{}{})
	fake.recordInvocation("Rollback", []interface{}{})
	fake.rollbackMutex.Unlock()
	if fake.RollbackStub != nil {
		return fake.RollbackStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rollbackReturns.result1
}

func (fake *FakeSQLClient) RollbackCallCount() int {
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	return len(fake.rollbackArgsForCall)
}

func (fake *FakeSQLClient) RollbackReturns(result1 error) {
	fake.RollbackStub = nil
	fake.rollbackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) RollbackReturnsOnCall(i int, result1 error) {
	fake.RollbackStub = nil
	if fake.rollbackReturnsOnCall == nil {
		fake.rollbackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rollbackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) Commit() error {
	fake.commitMutex.Lock()
	ret, specificReturn := fake.commitReturnsOnCall[len(fake.commitArgsForCall)]
	fake.commitArgsForCall = append(fake.commitArgsForCall, struct{}{})
	fake.recordInvocation("Commit", []interface{}{})
	fake.commitMutex.Unlock()
	if fake.CommitStub != nil {
		return fake.CommitStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.commitReturns.result1
}

func (fake *FakeSQLClient) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *FakeSQLClient) CommitReturns(result1 error) {
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) CommitReturnsOnCall(i int, result1 error) {
	fake.CommitStub = nil
	if fake.commitReturnsOnCall == nil {
		fake.commitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.commitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) InsertComposeCursor(arg1 *string) error {
	fake.insertComposeCursorMutex.Lock()
	ret, specificReturn := fake.insertComposeCursorReturnsOnCall[len(fake.insertComposeCursorArgsForCall)]
	fake.insertComposeCursorArgsForCall = append(fake.insertComposeCursorArgsForCall, struct {
		arg1 *string
	}{arg1})
	fake.recordInvocation("InsertComposeCursor", []interface{}{arg1})
	fake.insertComposeCursorMutex.Unlock()
	if fake.InsertComposeCursorStub != nil {
		return fake.InsertComposeCursorStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.insertComposeCursorReturns.result1
}

func (fake *FakeSQLClient) InsertComposeCursorCallCount() int {
	fake.insertComposeCursorMutex.RLock()
	defer fake.insertComposeCursorMutex.RUnlock()
	return len(fake.insertComposeCursorArgsForCall)
}

func (fake *FakeSQLClient) InsertComposeCursorArgsForCall(i int) *string {
	fake.insertComposeCursorMutex.RLock()
	defer fake.insertComposeCursorMutex.RUnlock()
	return fake.insertComposeCursorArgsForCall[i].arg1
}

func (fake *FakeSQLClient) InsertComposeCursorReturns(result1 error) {
	fake.InsertComposeCursorStub = nil
	fake.insertComposeCursorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) InsertComposeCursorReturnsOnCall(i int, result1 error) {
	fake.InsertComposeCursorStub = nil
	if fake.insertComposeCursorReturnsOnCall == nil {
		fake.insertComposeCursorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertComposeCursorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) FetchComposeCursor() (*string, error) {
	fake.fetchComposeCursorMutex.Lock()
	ret, specificReturn := fake.fetchComposeCursorReturnsOnCall[len(fake.fetchComposeCursorArgsForCall)]
	fake.fetchComposeCursorArgsForCall = append(fake.fetchComposeCursorArgsForCall, struct{}{})
	fake.recordInvocation("FetchComposeCursor", []interface{}{})
	fake.fetchComposeCursorMutex.Unlock()
	if fake.FetchComposeCursorStub != nil {
		return fake.FetchComposeCursorStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchComposeCursorReturns.result1, fake.fetchComposeCursorReturns.result2
}

func (fake *FakeSQLClient) FetchComposeCursorCallCount() int {
	fake.fetchComposeCursorMutex.RLock()
	defer fake.fetchComposeCursorMutex.RUnlock()
	return len(fake.fetchComposeCursorArgsForCall)
}

func (fake *FakeSQLClient) FetchComposeCursorReturns(result1 *string, result2 error) {
	fake.FetchComposeCursorStub = nil
	fake.fetchComposeCursorReturns = struct {
		result1 *string
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) FetchComposeCursorReturnsOnCall(i int, result1 *string, result2 error) {
	fake.FetchComposeCursorStub = nil
	if fake.fetchComposeCursorReturnsOnCall == nil {
		fake.fetchComposeCursorReturnsOnCall = make(map[int]struct {
			result1 *string
			result2 error
		})
	}
	fake.fetchComposeCursorReturnsOnCall[i] = struct {
		result1 *string
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) InsertComposeLatestEventID(arg1 string) error {
	fake.insertComposeLatestEventIDMutex.Lock()
	ret, specificReturn := fake.insertComposeLatestEventIDReturnsOnCall[len(fake.insertComposeLatestEventIDArgsForCall)]
	fake.insertComposeLatestEventIDArgsForCall = append(fake.insertComposeLatestEventIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("InsertComposeLatestEventID", []interface{}{arg1})
	fake.insertComposeLatestEventIDMutex.Unlock()
	if fake.InsertComposeLatestEventIDStub != nil {
		return fake.InsertComposeLatestEventIDStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.insertComposeLatestEventIDReturns.result1
}

func (fake *FakeSQLClient) InsertComposeLatestEventIDCallCount() int {
	fake.insertComposeLatestEventIDMutex.RLock()
	defer fake.insertComposeLatestEventIDMutex.RUnlock()
	return len(fake.insertComposeLatestEventIDArgsForCall)
}

func (fake *FakeSQLClient) InsertComposeLatestEventIDArgsForCall(i int) string {
	fake.insertComposeLatestEventIDMutex.RLock()
	defer fake.insertComposeLatestEventIDMutex.RUnlock()
	return fake.insertComposeLatestEventIDArgsForCall[i].arg1
}

func (fake *FakeSQLClient) InsertComposeLatestEventIDReturns(result1 error) {
	fake.InsertComposeLatestEventIDStub = nil
	fake.insertComposeLatestEventIDReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) InsertComposeLatestEventIDReturnsOnCall(i int, result1 error) {
	fake.InsertComposeLatestEventIDStub = nil
	if fake.insertComposeLatestEventIDReturnsOnCall == nil {
		fake.insertComposeLatestEventIDReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertComposeLatestEventIDReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) FetchComposeLatestEventID() (*string, error) {
	fake.fetchComposeLatestEventIDMutex.Lock()
	ret, specificReturn := fake.fetchComposeLatestEventIDReturnsOnCall[len(fake.fetchComposeLatestEventIDArgsForCall)]
	fake.fetchComposeLatestEventIDArgsForCall = append(fake.fetchComposeLatestEventIDArgsForCall, struct{}{})
	fake.recordInvocation("FetchComposeLatestEventID", []interface{}{})
	fake.fetchComposeLatestEventIDMutex.Unlock()
	if fake.FetchComposeLatestEventIDStub != nil {
		return fake.FetchComposeLatestEventIDStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchComposeLatestEventIDReturns.result1, fake.fetchComposeLatestEventIDReturns.result2
}

func (fake *FakeSQLClient) FetchComposeLatestEventIDCallCount() int {
	fake.fetchComposeLatestEventIDMutex.RLock()
	defer fake.fetchComposeLatestEventIDMutex.RUnlock()
	return len(fake.fetchComposeLatestEventIDArgsForCall)
}

func (fake *FakeSQLClient) FetchComposeLatestEventIDReturns(result1 *string, result2 error) {
	fake.FetchComposeLatestEventIDStub = nil
	fake.fetchComposeLatestEventIDReturns = struct {
		result1 *string
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) FetchComposeLatestEventIDReturnsOnCall(i int, result1 *string, result2 error) {
	fake.FetchComposeLatestEventIDStub = nil
	if fake.fetchComposeLatestEventIDReturnsOnCall == nil {
		fake.fetchComposeLatestEventIDReturnsOnCall = make(map[int]struct {
			result1 *string
			result2 error
		})
	}
	fake.fetchComposeLatestEventIDReturnsOnCall[i] = struct {
		result1 *string
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLClient) InsertComposeAuditEvents(arg1 []composeapi.AuditEvent) error {
	var arg1Copy []composeapi.AuditEvent
	if arg1 != nil {
		arg1Copy = make([]composeapi.AuditEvent, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.insertComposeAuditEventsMutex.Lock()
	ret, specificReturn := fake.insertComposeAuditEventsReturnsOnCall[len(fake.insertComposeAuditEventsArgsForCall)]
	fake.insertComposeAuditEventsArgsForCall = append(fake.insertComposeAuditEventsArgsForCall, struct {
		arg1 []composeapi.AuditEvent
	}{arg1Copy})
	fake.recordInvocation("InsertComposeAuditEvents", []interface{}{arg1Copy})
	fake.insertComposeAuditEventsMutex.Unlock()
	if fake.InsertComposeAuditEventsStub != nil {
		return fake.InsertComposeAuditEventsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.insertComposeAuditEventsReturns.result1
}

func (fake *FakeSQLClient) InsertComposeAuditEventsCallCount() int {
	fake.insertComposeAuditEventsMutex.RLock()
	defer fake.insertComposeAuditEventsMutex.RUnlock()
	return len(fake.insertComposeAuditEventsArgsForCall)
}

func (fake *FakeSQLClient) InsertComposeAuditEventsArgsForCall(i int) []composeapi.AuditEvent {
	fake.insertComposeAuditEventsMutex.RLock()
	defer fake.insertComposeAuditEventsMutex.RUnlock()
	return fake.insertComposeAuditEventsArgsForCall[i].arg1
}

func (fake *FakeSQLClient) InsertComposeAuditEventsReturns(result1 error) {
	fake.InsertComposeAuditEventsStub = nil
	fake.insertComposeAuditEventsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) InsertComposeAuditEventsReturnsOnCall(i int, result1 error) {
	fake.InsertComposeAuditEventsStub = nil
	if fake.insertComposeAuditEventsReturnsOnCall == nil {
		fake.insertComposeAuditEventsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertComposeAuditEventsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initSchemaMutex.RLock()
	defer fake.initSchemaMutex.RUnlock()
	fake.insertUsageEventListMutex.RLock()
	defer fake.insertUsageEventListMutex.RUnlock()
	fake.fetchLastGUIDMutex.RLock()
	defer fake.fetchLastGUIDMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	fake.queryJSONMutex.RLock()
	defer fake.queryJSONMutex.RUnlock()
	fake.queryRowJSONMutex.RLock()
	defer fake.queryRowJSONMutex.RUnlock()
	fake.updateViewsMutex.RLock()
	defer fake.updateViewsMutex.RUnlock()
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	fake.insertComposeCursorMutex.RLock()
	defer fake.insertComposeCursorMutex.RUnlock()
	fake.fetchComposeCursorMutex.RLock()
	defer fake.fetchComposeCursorMutex.RUnlock()
	fake.insertComposeLatestEventIDMutex.RLock()
	defer fake.insertComposeLatestEventIDMutex.RUnlock()
	fake.fetchComposeLatestEventIDMutex.RLock()
	defer fake.fetchComposeLatestEventIDMutex.RUnlock()
	fake.insertComposeAuditEventsMutex.RLock()
	defer fake.insertComposeAuditEventsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSQLClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.SQLClient = new(FakeSQLClient)