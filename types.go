// Copyright © 2026 Daniel C. Brotsky
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package airdata

import (
	"fmt"
	"reflect"
	"time"
)

// RecordData is an interface implemented by structs that can be
// marshaled to and unmarshaled from Airtable records. (Because this interface
// allows unmarshaling into the receiver, the method receiver must be
// a pointer to the struct type.)
//
// Every non-link field in a RecordData struct that is mapped to an Airtable field
// must have one of the named `...Field` types (or be a pointer to one of them)
// and a `field` tag that gives the Airtable field name or ID.
//
// The standard `...Field` definitions use fixed types, so you must know the type
// of the data being received from the Airtable side to use them.
// For polymorphic, read-only fields where you are not sure of the received data type
// (such as Lookup or Formula fields),
// you can use the `RawField` type which allows any data shape.
//
// RecordData fields that link to other records, rather than using one of the predefined
// `...Field` types, should be defined as a pointer to or a slice of pointers to a
// struct type that implements RecordData.
// When marshaling, the record IDs of the
// referenced structures will be included in the marshaled data. When unmarshaling,
// if the unmarshaling depth is greater than 0, structures will be allocated and
// populated with data from the linked records.
//
// A pointer field allows only one link and will give an unmarshaling error if the
// received data has more than one.
// A slice field allows any number of links (including 0) and,
// if empty, will be marshaled as an empty array (not nil).
type RecordData interface {
	GetRecordId() string                              // called when marshaling links
	RetrieveRecord(id string) (map[string]any, error) // called when unmarshaling links
	Marshal() (map[string]any, error)
	Unmarshal(map[string]any) error
}

var recordDataType = reflect.TypeFor[RecordData]()

func isRecordData(t reflect.Type) bool {
	return t.Implements(recordDataType)
}

type UnmarshalError struct {
	Message string
	Field   string
	Err     error
}

func (e UnmarshalError) Is(target error) bool {
	//goland:noinspection GoTypeAssertionOnErrors
	_, ok := target.(UnmarshalError)
	return ok
}

func (e UnmarshalError) Error() string {
	if e.Err == nil {
		if e.Field == "" {
			return fmt.Sprintf("error unmarshaling: %s", e.Message)
		}
		return fmt.Sprintf("error unmarshaling field %q: %s", e.Field, e.Message)
	}
	if e.Field == "" {
		if e.Message != "" {
			return fmt.Sprintf("error unmarshaling: %s: %v", e.Message, e.Err)
		}
		return fmt.Sprintf("error unmarshaling: %v", e.Err)
	}
	if e.Message != "" {
		return fmt.Sprintf("error unmarshaling field %q: %s: %v", e.Field, e.Message, e.Err)
	}
	return fmt.Sprintf("error unmarshaling field %q: %v", e.Field, e.Err)
}

func (e UnmarshalError) Unwrap() error { return e.Err }

// AttachmentField is a RecordData field type that maps to an Airtable
// [attachment field](https://airtable.com/developers/web/api/field-model#multipleattachment).
//
// When unmarshaling, the url will be a temporary Airtable-based URL.
//
// When marshaling, specify only the id or only the url and (optional) filename.
// When providing an url, it must be a publicly accessible URL.
type AttachmentField []AttachmentData

// BarcodeField is a RecordData field type that maps to an Airtable
// [barcode field](https://airtable.com/developers/web/api/field-model#barcode).
type BarcodeField struct {
	kind string // Airtable `type`, that is, the barcode type.
	data string // Airtable `text`, that is, the barcode data.
}

// ButtonField is a RecordData field type that maps to an Airtable
// [button field](https://airtable.com/developers/web/api/field-model#button).
type ButtonField struct {
	label string
	url   string
}

// CheckboxField is a RecordData field type that maps to an Airtable
// [checkbox field](https://airtable.com/developers/web/api/field-model#checkbox).
//
// When unmarshaling, a missing value is taken to be _false_. When marshaling,
// a _false_ is marshaled to `false`.
type CheckboxField bool

// CollaboratorField is a RecordData field that maps to an Airtable
// [collaborator field](https://airtable.com/developers/web/api/field-model#collaborator).
//
// When marshaling, only the id or email field is used (id has priority).
type CollaboratorField CollaboratorData

// CreatedByField is a RecordData field that maps to an Airtable
// [created-by field](https://airtable.com/developers/web/api/field-model#createdby).
//
// This field is read-only, so it is not marshaled.
type CreatedByField CollaboratorData

// CreatedTimeField is a RecordData field that maps to an Airtable
// [created-time field](https://airtable.com/developers/web/api/field-model#createdtime).
//
// This field is read-only, so it is not marshaled.
type CreatedTimeField time.Time

// CurrencyField is a RecordData field that maps to an Airtable
// [currency field](https://airtable.com/developers/web/api/field-model#currencynumber).
type CurrencyField float64

// DateField is a RecordData field type that maps to an Airtable
// [date field](https://airtable.com/developers/web/api/field-model#dateonly).
//
// Airtable dates are always interpreted as UTC. When marshaling, any time zone
// in the time will be shifted to UTC and the corresponding date used.
type DateField time.Time

// DateTimeField is a RecordData field type that maps to an Airtable
// [date/time field](https://airtable.com/developers/web/api/field-model#dateandtime).
//
// Airtable times are accurate to the millisecond. When marshaling, times
// are truncated to the nearest millisecond.
type DateTimeField time.Time

// DurationField is a RecordData field type that maps to an Airtable
// [duration field](https://airtable.com/developers/web/api/field-model#durationnumber).
//
// Airtable durations are precise to the millisecond. When marshaling,
// durations are truncated to the nearest millisecond.
type DurationField time.Duration

// EmailField is a RecordData field type that maps to an Airtable
// [email field](https://airtable.com/developers/web/api/field-model#emailtext).
type EmailField string

// LastModifiedByField is a RecordData field that maps to an Airtable
// [last-modified-by field](https://airtable.com/developers/web/api/field-model#lastmodifiedby).
//
// This field is read-only, so it is not marshaled.
type LastModifiedByField CollaboratorData

// LastModifiedTimeField is a RecordData field that maps to an Airtable
// [last-modified-time field](https://airtable.com/developers/web/api/field-model#lastmodifiedtime).
//
// This field is read-only, so it is not marshaled.
type LastModifiedTimeField time.Time

// LongTextField is a RecordData field type that maps to an Airtable
// [long text field](https://airtable.com/developers/web/api/field-model#multilinetext).
type LongTextField string

// MultipleCollaboratorField is a RecordData field that maps to an Airtable
// [multiple collaborator field](https://airtable.com/developers/web/api/field-model#multicollaborator).
type MultipleCollaboratorField []CollaboratorData

// MultipleSelectField is a RecordData field that maps to an Airtable
// [multiple select field](https://airtable.com/developers/web/api/field-model#multiselect).
type MultipleSelectField []string

// IntField is a RecordData field type that maps to an Airtable
// [number field](https://airtable.com/developers/web/api/field-model#decimalorintegernumber).
//
// When unmarshaling, any float value found will be truncated.
// When marshaling, no decimal point will be included.
type IntField int64

// FloatField is a RecordData field type that maps to an Airtable
// [number field](https://airtable.com/developers/web/api/field-model#decimalorintegernumber).
//
// When marshaling, all significant digits will be preserved.
type FloatField float64

// PercentField is a RecordData field type that maps to an Airtable
// [percent field](https://airtable.com/developers/web/api/field-model#percentnumber).
//
// When marshaling, all significant digits will be preserved.
type PercentField float64

// PhoneField is a RecordData field type that maps to an Airtable
// [phone field](https://airtable.com/developers/web/api/field-model#phone).
type PhoneField string

// RatingField is a RecordData field type that maps to an Airtable
// [rating field](https://airtable.com/developers/web/api/field-model#rating).
//
// Airtable ratings are always in the range 1-10 (inclusive).
type RatingField uint64

// RawField is a RecordData field type that maps to any Airtable field.
//
// The value is marshaled/unmarshaled as-is. When marshaling, it must contain data that
// can be marshaled to Airtable JSON. When unmarshaling, it will contain data that has
// been unmarshaled from the Airtable JSON.
type RawField any

// RichTextField is a RecordData field type that maps to an Airtable
// [rich text field](https://airtable.com/developers/web/api/field-model#richtext).
//
// Airtable's rich text field is just a long text field that contains Markdown.
type RichTextField string

// SingleLineTextField is a RecordData field type that maps to an Airtable
// [single line text field](https://airtable.com/developers/web/api/field-model#simpletext).
type SingleLineTextField string

// SingleSelectField is a RecordData field type that maps to an Airtable
// [single select field](https://airtable.com/developers/web/api/field-model#select).
type SingleSelectField string

// UrlField is a RecordData field type that maps to an Airtable
// [Url field](https://airtable.com/developers/web/api/field-model#urltext).
type UrlField string

// AttachmentData is how Airtable represents an attachment.
//
// On unmarshaling, the id field is always populated, and all the fields may be populated.
//
// On marshaling, only the id field is used if it is non-empty. If it is empty, then
// the url field must be non-empty and the filename field is included if it is non-empty.
type AttachmentData struct {
	id         string
	mimeType   string
	filename   string
	url        string
	size       int64
	height     int64
	width      int64
	thumbnails map[string]ThumbnailData
}

// ThumbnailData is how Airtable represents a thumbnail.
type ThumbnailData struct {
	url    string
	height int64
	width  int64
}

// CollaboratorData is how Airtable represents a collaborator.
//
// On unmarshaling, the id field is always populated, and all the fields may be populated.
//
// On marshaling, only the id and email fields are used.
// If the id field is non-empty, the email field is ignored.
// If the id field is empty, the email field must not be.
type CollaboratorData struct {
	id              string
	email           string
	name            string
	permissionLevel string
	profilePicUrl   string
}
