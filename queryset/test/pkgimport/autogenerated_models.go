// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	forex "github.com/jirfag/go-queryset/queryset/test/pkgimport/forex/v1"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set ExampleQuerySet

// ExampleQuerySet is an queryset type for Example
type ExampleQuerySet struct {
	db *gorm.DB
}

// NewExampleQuerySet constructs new ExampleQuerySet
func NewExampleQuerySet(db *gorm.DB) ExampleQuerySet {
	return ExampleQuerySet{
		db: db.Model(&Example{}),
	}
}

func (qs ExampleQuerySet) w(db *gorm.DB) ExampleQuerySet {
	return NewExampleQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) All(ret *[]Example) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Example) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Currency1Eq is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Eq(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency1` = ?", currency1))
}

// Currency1Gt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Gt(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency1` > ?", currency1))
}

// Currency1Gte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Gte(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency1` >= ?", currency1))
}

// Currency1In is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1In(currency1 forex.Currency1, currency1Rest ...forex.Currency1) ExampleQuerySet {
	iArgs := []interface{}{currency1}
	for _, arg := range currency1Rest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("`currency1` IN (?)", iArgs))
}

// Currency1Lt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Lt(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency1` < ?", currency1))
}

// Currency1Lte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Lte(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency1` <= ?", currency1))
}

// Currency1Ne is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Ne(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency1` != ?", currency1))
}

// Currency1NotIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1NotIn(currency1 forex.Currency1, currency1Rest ...forex.Currency1) ExampleQuerySet {
	iArgs := []interface{}{currency1}
	for _, arg := range currency1Rest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("`currency1` NOT IN (?)", iArgs))
}

// Currency2Eq is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Eq(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency2` = ?", currency2))
}

// Currency2In is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2In(currency2 forex.Currency2, currency2Rest ...forex.Currency2) ExampleQuerySet {
	iArgs := []interface{}{currency2}
	for _, arg := range currency2Rest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("`currency2` IN (?)", iArgs))
}

// Currency2Ne is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Ne(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency2` != ?", currency2))
}

// Currency2NotIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2NotIn(currency2 forex.Currency2, currency2Rest ...forex.Currency2) ExampleQuerySet {
	iArgs := []interface{}{currency2}
	for _, arg := range currency2Rest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("`currency2` NOT IN (?)", iArgs))
}

// Currency3Eq is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Eq(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency3` = ?", currency3))
}

// Currency3In is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3In(currency3 forex.Currency3, currency3Rest ...forex.Currency3) ExampleQuerySet {
	iArgs := []interface{}{currency3}
	for _, arg := range currency3Rest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("`currency3` IN (?)", iArgs))
}

// Currency3Ne is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Ne(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("`currency3` != ?", currency3))
}

// Currency3NotIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3NotIn(currency3 forex.Currency3, currency3Rest ...forex.Currency3) ExampleQuerySet {
	iArgs := []interface{}{currency3}
	for _, arg := range currency3Rest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("`currency3` NOT IN (?)", iArgs))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Example) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Delete() error {
	return qs.db.Delete(Example{}).Error
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) GetUpdater() ExampleUpdater {
	return NewExampleUpdater(qs.db)
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Limit(limit int) ExampleQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ExampleQuerySet) One(ret *Example) error {
	return qs.db.First(ret).Error
}

// OrderAscByCurrency1 is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderAscByCurrency1() ExampleQuerySet {
	return qs.w(qs.db.Order("currency1 ASC"))
}

// OrderAscByPriceID is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderAscByPriceID() ExampleQuerySet {
	return qs.w(qs.db.Order("price_id ASC"))
}

// OrderDescByCurrency1 is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderDescByCurrency1() ExampleQuerySet {
	return qs.w(qs.db.Order("currency1 DESC"))
}

// OrderDescByPriceID is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderDescByPriceID() ExampleQuerySet {
	return qs.w(qs.db.Order("price_id DESC"))
}

// PriceIDEq is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDEq(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("`price_id` = ?", priceID))
}

// PriceIDGt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDGt(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("`price_id` > ?", priceID))
}

// PriceIDGte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDGte(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("`price_id` >= ?", priceID))
}

// PriceIDIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDIn(priceID int64, priceIDRest ...int64) ExampleQuerySet {
	iArgs := []interface{}{priceID}
	for _, arg := range priceIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("`price_id` IN (?)", iArgs))
}

// PriceIDLt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDLt(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("`price_id` < ?", priceID))
}

// PriceIDLte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDLte(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("`price_id` <= ?", priceID))
}

// PriceIDNe is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDNe(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("`price_id` != ?", priceID))
}

// PriceIDNotIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDNotIn(priceID int64, priceIDRest ...int64) ExampleQuerySet {
	iArgs := []interface{}{priceID}
	for _, arg := range priceIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("`price_id` NOT IN (?)", iArgs))
}

// SetCurrency1 is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) SetCurrency1(currency1 forex.Currency1) ExampleUpdater {
	u.fields[string(ExampleDBSchema.Currency1)] = currency1
	return u
}

// SetCurrency2 is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) SetCurrency2(currency2 forex.Currency2) ExampleUpdater {
	u.fields[string(ExampleDBSchema.Currency2)] = currency2
	return u
}

// SetCurrency3 is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) SetCurrency3(currency3 forex.Currency3) ExampleUpdater {
	u.fields[string(ExampleDBSchema.Currency3)] = currency3
	return u
}

// SetPriceID is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) SetPriceID(priceID int64) ExampleUpdater {
	u.fields[string(ExampleDBSchema.PriceID)] = priceID
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set ExampleQuerySet

// ===== BEGIN of Example modifiers

// ExampleDBSchemaField describes database schema field. It requires for method 'Update'
type ExampleDBSchemaField string

func (f ExampleDBSchemaField) String() string {
	return string(f)
}

// ExampleDBSchema stores db field names of Example
var ExampleDBSchema = struct {
	PriceID   ExampleDBSchemaField
	Currency1 ExampleDBSchemaField
	Currency2 ExampleDBSchemaField
	Currency3 ExampleDBSchemaField
}{

	PriceID:   ExampleDBSchemaField("price_id"),
	Currency1: ExampleDBSchemaField("currency1"),
	Currency2: ExampleDBSchemaField("currency2"),
	Currency3: ExampleDBSchemaField("currency3"),
}

// Update updates Example fields by primary key
func (o *Example) Update(db *gorm.DB, fields ...ExampleDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"price_id":  o.PriceID,
		"currency1": o.Currency1,
		"currency2": o.Currency2,
		"currency3": o.Currency3,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Example %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// ExampleUpdater is an Example updates manager
type ExampleUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewExampleUpdater creates new Example updater
func NewExampleUpdater(db *gorm.DB) ExampleUpdater {
	return ExampleUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Example{}),
	}
}

// ===== END of Example modifiers

// ===== END of all query sets
