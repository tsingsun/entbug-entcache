// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entgo.io/bug/ent/schema"
	"entgo.io/bug/ent/securityjournal"
	"entgo.io/bug/ent/securityposition"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	securityjournalFields := schema.SecurityJournal{}.Fields()
	_ = securityjournalFields
	// securityjournalDescAccountID is the schema descriptor for account_id field.
	securityjournalDescAccountID := securityjournalFields[2].Descriptor()
	// securityjournal.DefaultAccountID holds the default value on creation for the account_id field.
	securityjournal.DefaultAccountID = securityjournalDescAccountID.Default.(int)
	// securityjournalDescBizNo is the schema descriptor for biz_no field.
	securityjournalDescBizNo := securityjournalFields[3].Descriptor()
	// securityjournal.BizNoValidator is a validator for the "biz_no" field. It is called by the builders before save.
	securityjournal.BizNoValidator = securityjournalDescBizNo.Validators[0].(func(string) error)
	// securityjournalDescBizType is the schema descriptor for biz_type field.
	securityjournalDescBizType := securityjournalFields[4].Descriptor()
	// securityjournal.BizTypeValidator is a validator for the "biz_type" field. It is called by the builders before save.
	securityjournal.BizTypeValidator = securityjournalDescBizType.Validators[0].(func(string) error)
	// securityjournalDescTradeCode is the schema descriptor for trade_code field.
	securityjournalDescTradeCode := securityjournalFields[6].Descriptor()
	// securityjournal.TradeCodeValidator is a validator for the "trade_code" field. It is called by the builders before save.
	securityjournal.TradeCodeValidator = securityjournalDescTradeCode.Validators[0].(func(string) error)
	// securityjournalDescChangeType is the schema descriptor for change_type field.
	securityjournalDescChangeType := securityjournalFields[7].Descriptor()
	// securityjournal.ChangeTypeValidator is a validator for the "change_type" field. It is called by the builders before save.
	securityjournal.ChangeTypeValidator = securityjournalDescChangeType.Validators[0].(func(string) error)
	// securityjournalDescMaterialNo is the schema descriptor for material_no field.
	securityjournalDescMaterialNo := securityjournalFields[12].Descriptor()
	// securityjournal.MaterialNoValidator is a validator for the "material_no" field. It is called by the builders before save.
	securityjournal.MaterialNoValidator = securityjournalDescMaterialNo.Validators[0].(func(string) error)
	// securityjournalDescMaterialName is the schema descriptor for material_name field.
	securityjournalDescMaterialName := securityjournalFields[13].Descriptor()
	// securityjournal.MaterialNameValidator is a validator for the "material_name" field. It is called by the builders before save.
	securityjournal.MaterialNameValidator = securityjournalDescMaterialName.Validators[0].(func(string) error)
	// securityjournalDescCurrency is the schema descriptor for currency field.
	securityjournalDescCurrency := securityjournalFields[16].Descriptor()
	// securityjournal.CurrencyValidator is a validator for the "currency" field. It is called by the builders before save.
	securityjournal.CurrencyValidator = securityjournalDescCurrency.Validators[0].(func(string) error)
	// securityjournalDescQty is the schema descriptor for qty field.
	securityjournalDescQty := securityjournalFields[17].Descriptor()
	// securityjournal.DefaultQty holds the default value on creation for the qty field.
	securityjournal.DefaultQty = securityjournalDescQty.Default.(float64)
	// securityjournalDescUnit is the schema descriptor for unit field.
	securityjournalDescUnit := securityjournalFields[22].Descriptor()
	// securityjournal.UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	securityjournal.UnitValidator = securityjournalDescUnit.Validators[0].(func(string) error)
	// securityjournalDescSpec is the schema descriptor for spec field.
	securityjournalDescSpec := securityjournalFields[23].Descriptor()
	// securityjournal.SpecValidator is a validator for the "spec" field. It is called by the builders before save.
	securityjournal.SpecValidator = securityjournalDescSpec.Validators[0].(func(string) error)
	// securityjournalDescCount is the schema descriptor for count field.
	securityjournalDescCount := securityjournalFields[27].Descriptor()
	// securityjournal.DefaultCount holds the default value on creation for the count field.
	securityjournal.DefaultCount = securityjournalDescCount.Default.(int)
	// securityjournalDescTamper is the schema descriptor for tamper field.
	securityjournalDescTamper := securityjournalFields[29].Descriptor()
	// securityjournal.TamperValidator is a validator for the "tamper" field. It is called by the builders before save.
	securityjournal.TamperValidator = securityjournalDescTamper.Validators[0].(func(string) error)
	// securityjournalDescSummary is the schema descriptor for summary field.
	securityjournalDescSummary := securityjournalFields[30].Descriptor()
	// securityjournal.SummaryValidator is a validator for the "summary" field. It is called by the builders before save.
	securityjournal.SummaryValidator = securityjournalDescSummary.Validators[0].(func(string) error)
	// securityjournalDescPairSubjectCode is the schema descriptor for pair_subject_code field.
	securityjournalDescPairSubjectCode := securityjournalFields[31].Descriptor()
	// securityjournal.PairSubjectCodeValidator is a validator for the "pair_subject_code" field. It is called by the builders before save.
	securityjournal.PairSubjectCodeValidator = securityjournalDescPairSubjectCode.Validators[0].(func(string) error)
	// securityjournalDescStlCurrency is the schema descriptor for stl_currency field.
	securityjournalDescStlCurrency := securityjournalFields[36].Descriptor()
	// securityjournal.StlCurrencyValidator is a validator for the "stl_currency" field. It is called by the builders before save.
	securityjournal.StlCurrencyValidator = securityjournalDescStlCurrency.Validators[0].(func(string) error)
	// securityjournalDescIsSettlement is the schema descriptor for is_settlement field.
	securityjournalDescIsSettlement := securityjournalFields[39].Descriptor()
	// securityjournal.DefaultIsSettlement holds the default value on creation for the is_settlement field.
	securityjournal.DefaultIsSettlement = securityjournalDescIsSettlement.Default.(string)
	// securityjournal.IsSettlementValidator is a validator for the "is_settlement" field. It is called by the builders before save.
	securityjournal.IsSettlementValidator = securityjournalDescIsSettlement.Validators[0].(func(string) error)
	// securityjournalDescTransNo is the schema descriptor for trans_no field.
	securityjournalDescTransNo := securityjournalFields[41].Descriptor()
	// securityjournal.TransNoValidator is a validator for the "trans_no" field. It is called by the builders before save.
	securityjournal.TransNoValidator = securityjournalDescTransNo.Validators[0].(func(string) error)
	// securityjournalDescAccountingNo is the schema descriptor for accounting_no field.
	securityjournalDescAccountingNo := securityjournalFields[43].Descriptor()
	// securityjournal.AccountingNoValidator is a validator for the "accounting_no" field. It is called by the builders before save.
	securityjournal.AccountingNoValidator = securityjournalDescAccountingNo.Validators[0].(func(string) error)
	// securityjournalDescRefTransNo is the schema descriptor for ref_trans_no field.
	securityjournalDescRefTransNo := securityjournalFields[49].Descriptor()
	// securityjournal.RefTransNoValidator is a validator for the "ref_trans_no" field. It is called by the builders before save.
	securityjournal.RefTransNoValidator = securityjournalDescRefTransNo.Validators[0].(func(string) error)
	// securityjournalDescIsDayBooking is the schema descriptor for is_day_booking field.
	securityjournalDescIsDayBooking := securityjournalFields[50].Descriptor()
	// securityjournal.IsDayBookingValidator is a validator for the "is_day_booking" field. It is called by the builders before save.
	securityjournal.IsDayBookingValidator = securityjournalDescIsDayBooking.Validators[0].(func(string) error)
	// securityjournalDescIsEffectFund is the schema descriptor for is_effect_fund field.
	securityjournalDescIsEffectFund := securityjournalFields[51].Descriptor()
	// securityjournal.DefaultIsEffectFund holds the default value on creation for the is_effect_fund field.
	securityjournal.DefaultIsEffectFund = securityjournalDescIsEffectFund.Default.(string)
	// securityjournal.IsEffectFundValidator is a validator for the "is_effect_fund" field. It is called by the builders before save.
	securityjournal.IsEffectFundValidator = securityjournalDescIsEffectFund.Validators[0].(func(string) error)
	securitypositionFields := schema.SecurityPosition{}.Fields()
	_ = securitypositionFields
	// securitypositionDescParentID is the schema descriptor for parent_id field.
	securitypositionDescParentID := securitypositionFields[1].Descriptor()
	// securityposition.DefaultParentID holds the default value on creation for the parent_id field.
	securityposition.DefaultParentID = securitypositionDescParentID.Default.(int)
	// securitypositionDescMaterialNo is the schema descriptor for material_no field.
	securitypositionDescMaterialNo := securitypositionFields[9].Descriptor()
	// securityposition.MaterialNoValidator is a validator for the "material_no" field. It is called by the builders before save.
	securityposition.MaterialNoValidator = securitypositionDescMaterialNo.Validators[0].(func(string) error)
	// securitypositionDescBalance is the schema descriptor for balance field.
	securitypositionDescBalance := securitypositionFields[11].Descriptor()
	// securityposition.DefaultBalance holds the default value on creation for the balance field.
	securityposition.DefaultBalance = securitypositionDescBalance.Default.(float64)
	// securitypositionDescAvailable is the schema descriptor for available field.
	securitypositionDescAvailable := securitypositionFields[12].Descriptor()
	// securityposition.DefaultAvailable holds the default value on creation for the available field.
	securityposition.DefaultAvailable = securitypositionDescAvailable.Default.(float64)
	// securitypositionDescFreeze is the schema descriptor for freeze field.
	securitypositionDescFreeze := securitypositionFields[13].Descriptor()
	// securityposition.DefaultFreeze holds the default value on creation for the freeze field.
	securityposition.DefaultFreeze = securitypositionDescFreeze.Default.(float64)
	// securitypositionDescAfloat is the schema descriptor for afloat field.
	securitypositionDescAfloat := securitypositionFields[14].Descriptor()
	// securityposition.DefaultAfloat holds the default value on creation for the afloat field.
	securityposition.DefaultAfloat = securitypositionDescAfloat.Default.(float64)
	// securitypositionDescUnit is the schema descriptor for unit field.
	securitypositionDescUnit := securitypositionFields[15].Descriptor()
	// securityposition.UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	securityposition.UnitValidator = securitypositionDescUnit.Validators[0].(func(string) error)
	// securitypositionDescSpec is the schema descriptor for spec field.
	securitypositionDescSpec := securitypositionFields[16].Descriptor()
	// securityposition.SpecValidator is a validator for the "spec" field. It is called by the builders before save.
	securityposition.SpecValidator = securitypositionDescSpec.Validators[0].(func(string) error)
	// securitypositionDescCurrency is the schema descriptor for currency field.
	securitypositionDescCurrency := securitypositionFields[17].Descriptor()
	// securityposition.CurrencyValidator is a validator for the "currency" field. It is called by the builders before save.
	securityposition.CurrencyValidator = securitypositionDescCurrency.Validators[0].(func(string) error)
	// securitypositionDescPrice is the schema descriptor for price field.
	securitypositionDescPrice := securitypositionFields[18].Descriptor()
	// securityposition.DefaultPrice holds the default value on creation for the price field.
	securityposition.DefaultPrice = securitypositionDescPrice.Default.(float64)
	// securitypositionDescStlCurrency is the schema descriptor for stl_currency field.
	securitypositionDescStlCurrency := securitypositionFields[23].Descriptor()
	// securityposition.StlCurrencyValidator is a validator for the "stl_currency" field. It is called by the builders before save.
	securityposition.StlCurrencyValidator = securitypositionDescStlCurrency.Validators[0].(func(string) error)
	// securitypositionDescStlMarginLv is the schema descriptor for stl_margin_lv field.
	securitypositionDescStlMarginLv := securitypositionFields[28].Descriptor()
	// securityposition.DefaultStlMarginLv holds the default value on creation for the stl_margin_lv field.
	securityposition.DefaultStlMarginLv = securitypositionDescStlMarginLv.Default.(float64)
	// securitypositionDescUniqueTag is the schema descriptor for unique_tag field.
	securitypositionDescUniqueTag := securitypositionFields[32].Descriptor()
	// securityposition.UniqueTagValidator is a validator for the "unique_tag" field. It is called by the builders before save.
	securityposition.UniqueTagValidator = securitypositionDescUniqueTag.Validators[0].(func(string) error)
}
