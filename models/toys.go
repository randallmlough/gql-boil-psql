// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Toy is an object representing the database table.
type Toy struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt   null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt   null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Description string      `boil:"description" json:"description" toml:"description" yaml:"description"`
	Color       null.String `boil:"color" json:"color,omitempty" toml:"color" yaml:"color,omitempty"`
	PetID       null.Int    `boil:"pet_id" json:"pet_id,omitempty" toml:"pet_id" yaml:"pet_id,omitempty"`

	R *toyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L toyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ToyColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	Description string
	Color       string
	PetID       string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	Description: "description",
	Color:       "color",
	PetID:       "pet_id",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var ToyWhere = struct {
	ID          whereHelperint
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
	DeletedAt   whereHelpernull_Time
	Description whereHelperstring
	Color       whereHelpernull_String
	PetID       whereHelpernull_Int
}{
	ID:          whereHelperint{field: "\"gqlboil\".\"toys\".\"id\""},
	CreatedAt:   whereHelpernull_Time{field: "\"gqlboil\".\"toys\".\"created_at\""},
	UpdatedAt:   whereHelpernull_Time{field: "\"gqlboil\".\"toys\".\"updated_at\""},
	DeletedAt:   whereHelpernull_Time{field: "\"gqlboil\".\"toys\".\"deleted_at\""},
	Description: whereHelperstring{field: "\"gqlboil\".\"toys\".\"description\""},
	Color:       whereHelpernull_String{field: "\"gqlboil\".\"toys\".\"color\""},
	PetID:       whereHelpernull_Int{field: "\"gqlboil\".\"toys\".\"pet_id\""},
}

// ToyRels is where relationship names are stored.
var ToyRels = struct {
	Pet string
}{
	Pet: "Pet",
}

// toyR is where relationships are stored.
type toyR struct {
	Pet *Pet `boil:"Pet" json:"Pet" toml:"Pet" yaml:"Pet"`
}

// NewStruct creates a new relationship struct
func (*toyR) NewStruct() *toyR {
	return &toyR{}
}

// toyL is where Load methods for each relationship are stored.
type toyL struct{}

var (
	toyAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "description", "color", "pet_id"}
	toyColumnsWithoutDefault = []string{"deleted_at", "description", "color", "pet_id"}
	toyColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	toyPrimaryKeyColumns     = []string{"id"}
)

type (
	// ToySlice is an alias for a slice of pointers to Toy.
	// This should generally be used opposed to []Toy.
	ToySlice []*Toy
	// ToyHook is the signature for custom Toy hook methods
	ToyHook func(context.Context, boil.ContextExecutor, *Toy) error

	toyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	toyType                 = reflect.TypeOf(&Toy{})
	toyMapping              = queries.MakeStructMapping(toyType)
	toyPrimaryKeyMapping, _ = queries.BindMapping(toyType, toyMapping, toyPrimaryKeyColumns)
	toyInsertCacheMut       sync.RWMutex
	toyInsertCache          = make(map[string]insertCache)
	toyUpdateCacheMut       sync.RWMutex
	toyUpdateCache          = make(map[string]updateCache)
	toyUpsertCacheMut       sync.RWMutex
	toyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var toyBeforeInsertHooks []ToyHook
var toyBeforeUpdateHooks []ToyHook
var toyBeforeDeleteHooks []ToyHook
var toyBeforeUpsertHooks []ToyHook

var toyAfterInsertHooks []ToyHook
var toyAfterSelectHooks []ToyHook
var toyAfterUpdateHooks []ToyHook
var toyAfterDeleteHooks []ToyHook
var toyAfterUpsertHooks []ToyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Toy) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range toyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Toy) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range toyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Toy) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range toyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Toy) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range toyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Toy) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range toyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Toy) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range toyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Toy) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range toyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Toy) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range toyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Toy) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range toyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddToyHook registers your hook function for all future operations.
func AddToyHook(hookPoint boil.HookPoint, toyHook ToyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		toyBeforeInsertHooks = append(toyBeforeInsertHooks, toyHook)
	case boil.BeforeUpdateHook:
		toyBeforeUpdateHooks = append(toyBeforeUpdateHooks, toyHook)
	case boil.BeforeDeleteHook:
		toyBeforeDeleteHooks = append(toyBeforeDeleteHooks, toyHook)
	case boil.BeforeUpsertHook:
		toyBeforeUpsertHooks = append(toyBeforeUpsertHooks, toyHook)
	case boil.AfterInsertHook:
		toyAfterInsertHooks = append(toyAfterInsertHooks, toyHook)
	case boil.AfterSelectHook:
		toyAfterSelectHooks = append(toyAfterSelectHooks, toyHook)
	case boil.AfterUpdateHook:
		toyAfterUpdateHooks = append(toyAfterUpdateHooks, toyHook)
	case boil.AfterDeleteHook:
		toyAfterDeleteHooks = append(toyAfterDeleteHooks, toyHook)
	case boil.AfterUpsertHook:
		toyAfterUpsertHooks = append(toyAfterUpsertHooks, toyHook)
	}
}

// One returns a single toy record from the query.
func (q toyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Toy, error) {
	o := &Toy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for toys")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Toy records from the query.
func (q toyQuery) All(ctx context.Context, exec boil.ContextExecutor) (ToySlice, error) {
	var o []*Toy

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Toy slice")
	}

	if len(toyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Toy records in the query.
func (q toyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count toys rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q toyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if toys exists")
	}

	return count > 0, nil
}

// Pet pointed to by the foreign key.
func (o *Toy) Pet(mods ...qm.QueryMod) petQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PetID),
	}

	queryMods = append(queryMods, mods...)

	query := Pets(queryMods...)
	queries.SetFrom(query.Query, "\"gqlboil\".\"pet\"")

	return query
}

// LoadPet allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (toyL) LoadPet(ctx context.Context, e boil.ContextExecutor, singular bool, maybeToy interface{}, mods queries.Applicator) error {
	var slice []*Toy
	var object *Toy

	if singular {
		object = maybeToy.(*Toy)
	} else {
		slice = *maybeToy.(*[]*Toy)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &toyR{}
		}
		if !queries.IsNil(object.PetID) {
			args = append(args, object.PetID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &toyR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.PetID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.PetID) {
				args = append(args, obj.PetID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`gqlboil.pet`),
		qm.WhereIn(`gqlboil.pet.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pet")
	}

	var resultSlice []*Pet
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pet")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for pet")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for pet")
	}

	if len(toyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Pet = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.PetID, foreign.ID) {
				local.R.Pet = foreign
				break
			}
		}
	}

	return nil
}

// SetPet of the toy to the related item.
// Sets o.R.Pet to related.
// Adds o to related.R.Toys.
func (o *Toy) SetPet(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Pet) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"gqlboil\".\"toys\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pet_id"}),
		strmangle.WhereClause("\"", "\"", 2, toyPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.PetID, related.ID)
	if o.R == nil {
		o.R = &toyR{
			Pet: related,
		}
	} else {
		o.R.Pet = related
	}

	if related.R == nil {
		related.R = &petR{
			Toys: ToySlice{o},
		}
	} else {
		related.R.Toys = append(related.R.Toys, o)
	}

	return nil
}

// RemovePet relationship.
// Sets o.R.Pet to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Toy) RemovePet(ctx context.Context, exec boil.ContextExecutor, related *Pet) error {
	var err error

	queries.SetScanner(&o.PetID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("pet_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Pet = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Toys {
		if queries.Equal(o.PetID, ri.PetID) {
			continue
		}

		ln := len(related.R.Toys)
		if ln > 1 && i < ln-1 {
			related.R.Toys[i] = related.R.Toys[ln-1]
		}
		related.R.Toys = related.R.Toys[:ln-1]
		break
	}
	return nil
}

// Toys retrieves all the records using an executor.
func Toys(mods ...qm.QueryMod) toyQuery {
	mods = append(mods, qm.From("\"gqlboil\".\"toys\""))
	return toyQuery{NewQuery(mods...)}
}

// FindToy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindToy(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Toy, error) {
	toyObj := &Toy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"gqlboil\".\"toys\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, toyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from toys")
	}

	return toyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Toy) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no toys provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(toyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	toyInsertCacheMut.RLock()
	cache, cached := toyInsertCache[key]
	toyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			toyAllColumns,
			toyColumnsWithDefault,
			toyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(toyType, toyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(toyType, toyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"gqlboil\".\"toys\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"gqlboil\".\"toys\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into toys")
	}

	if !cached {
		toyInsertCacheMut.Lock()
		toyInsertCache[key] = cache
		toyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Toy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Toy) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	toyUpdateCacheMut.RLock()
	cache, cached := toyUpdateCache[key]
	toyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			toyAllColumns,
			toyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update toys, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"gqlboil\".\"toys\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, toyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(toyType, toyMapping, append(wl, toyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update toys row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for toys")
	}

	if !cached {
		toyUpdateCacheMut.Lock()
		toyUpdateCache[key] = cache
		toyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q toyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for toys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for toys")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ToySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), toyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"gqlboil\".\"toys\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, toyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in toy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all toy")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Toy) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no toys provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(toyColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	toyUpsertCacheMut.RLock()
	cache, cached := toyUpsertCache[key]
	toyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			toyAllColumns,
			toyColumnsWithDefault,
			toyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			toyAllColumns,
			toyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert toys, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(toyPrimaryKeyColumns))
			copy(conflict, toyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"gqlboil\".\"toys\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(toyType, toyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(toyType, toyMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert toys")
	}

	if !cached {
		toyUpsertCacheMut.Lock()
		toyUpsertCache[key] = cache
		toyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Toy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Toy) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Toy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), toyPrimaryKeyMapping)
	sql := "DELETE FROM \"gqlboil\".\"toys\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from toys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for toys")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q toyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no toyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from toys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for toys")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ToySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(toyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), toyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"gqlboil\".\"toys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, toyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from toy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for toys")
	}

	if len(toyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Toy) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindToy(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ToySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ToySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), toyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"gqlboil\".\"toys\".* FROM \"gqlboil\".\"toys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, toyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ToySlice")
	}

	*o = slice

	return nil
}

// ToyExists checks if the Toy row exists.
func ToyExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"gqlboil\".\"toys\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if toys exists")
	}

	return exists, nil
}
