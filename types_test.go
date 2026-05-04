package airdata

import (
	"reflect"
	"testing"
)

type AllTypeStruct struct {
	Field01 string
	Field02 AttachmentField           `field:"AttachmentField"`
	Field03 BarcodeField              `field:"BarcodeField"`
	Field04 ButtonField               `field:"ButtonField"`
	Field05 CheckboxField             `field:"CheckboxField"`
	Field06 CollaboratorField         `field:"CollaboratorField"`
	Field07 CreatedByField            `field:"CreatorField"`
	Field08 CreatedTimeField          `field:"CreatedTimeField"`
	Field09 CurrencyField             `field:"CurrencyField"`
	Field10 DateField                 `field:"DateField"`
	Field11 DateTimeField             `field:"DateTimeField"`
	Field12 DurationField             `field:"DurationField"`
	Field13 EmailField                `field:"EmailField"`
	Field14 LastModifiedByField       `field:"LastModifiedByField"`
	Field15 LastModifiedTimeField     `field:"LastModifiedTimeField"`
	Field16 LongTextField             `field:"LongTextField"`
	Field17 MultipleCollaboratorField `field:"MultipleCollaboratorField"`
	Field18 MultipleSelectField       `field:"MultipleSelectField"`
	Field19 IntField                  `field:"IntField"`
	Field20 FloatField                `field:"FloatField"`
	Field21 PercentField              `field:"PercentField"`
	Field22 PhoneField                `field:"PhoneField"`
	Field23 RatingField               `field:"RatingField"`
	Field24 RawField                  `field:"RawField"`
	Field25 RichTextField             `field:"RichTextField"`
	Field26 SingleLineTextField       `field:"SingleLineTextField"`
	Field27 SingleSelectField         `field:"SingleSelectField"`
	Field28 UrlField                  `field:"UrlField"`
	Field29 *AllTypeStruct            `field:"SingleRecordLink"`
	Field30 []*AllTypeStruct          `field:"MultipleRecordLinks"`
	Field31 *int                      `field:"Invalid"`
	Field32 []string                  `field:"Invalid"`
}

func (a *AllTypeStruct) GetRecordId() string {
	return a.Field01
}

func (a *AllTypeStruct) RetrieveRecord(id string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AllTypeStruct) Marshal() (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AllTypeStruct) Unmarshal(fields map[string]any) error {
	//TODO implement me
	panic("implement me")
}

func TestTypeReflection(t *testing.T) {
	v := reflect.ValueOf(&AllTypeStruct{})
	vt := v.Type()
	// work through the tagged fields
	for i := 0; i < vt.NumField(); i++ {
		ft := vt.Field(i)
		tag := ft.Tag.Get("field")
		if tag == "" {
			t.Logf("Ignoring %s because it has no tag", ft.Name)
			continue
		}
		fv := v.Field(i)
		if tag == ft.Type.Name() {
			t.Logf("%s has known type %s", ft.Name, ft.Type.Name())
			continue
		}
		if fv.Kind() == reflect.Pointer {
			fvt := fv.Type()
			if !isRecordData(fvt) {
				if tag == "Invalid" {
					t.Logf("Ignoring pointer field %s because its tag is 'Invalid'", ft.Name)
					continue
				}
				t.Errorf("%s has type %T which does not implement RecordData", ft.Name, fv.Elem())
				continue
			}
			t.Logf("%s is a single link to %s", ft.Name, fvt.Elem().Name())
			continue
		}
		if fv.Kind() == reflect.Slice {
			fvt := fv.Type().Elem()
			if !isRecordData(fvt) {
				if tag == "Invalid" {
					t.Logf("Ignoring slice field %s because its tag is 'Invalid'", ft.Name)
					continue
				}
				t.Errorf("%s is a slice of a type which does not implement RecordData", ft.Name)
				continue
			}
			t.Logf("%s is a multiple link to %s", ft.Name, fvt.Elem().Name())
			continue
		}
	}
}
