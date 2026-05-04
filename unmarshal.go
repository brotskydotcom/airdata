package airdata

import (
	"fmt"
	"reflect"
	"time"
)

// UnmarshalRecord unmarshals an airtable.Record into a struct.
func UnmarshalRecord(target any, fields map[string]any) error {
	if target == nil {
		// only true if called with literal nil, not a typed nil pointer
		return fmt.Errorf("cannot unmarshal to nil")
	}
	// Now we have a concrete value that we can examine
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("cannot unmarshal to non-pointer type")
	}
	if v.IsNil() {
		return fmt.Errorf("can only unmarshal to a valid struct pointer, not nil")
	}
	s := v.Elem()
	if s.Kind() != reflect.Struct {
		return fmt.Errorf("can only unmarshal to a struct pointer")
	}
	if !isRecordData(v.Type()) {
		return fmt.Errorf("can only unmarshal to a struct that implements RecordData")
	}
	// set to a zero value before unmarshaling fields
	s.Set(reflect.New(s.Type()).Elem())
	if err := unmarshalStructValue(s, fields); err != nil {
		// reset partially unmarshaled fields
		s.Set(reflect.New(s.Type()).Elem())
		return err
	}
	return nil
}

func unmarshalStructValue(s reflect.Value, fields map[string]any) error {
	st := s.Type()
	// work through the tagged fields
	for i := 0; i < st.NumField(); i++ {
		ft := st.Field(i)
		tag := ft.Tag.Get("field")
		if tag == "" {
			continue
		}
		fv := s.Field(i)
		data, ok := fields[tag]
		if !ok {
			continue
		}
		var v reflect.Value
		var err error
		switch fv.Type().Name() {
		case "AttachmentField":
			v, err = unmarshalAttachmentField(data)
		case "BarcodeField":
			v, err = unmarshalBarcodeField(data)
		case "ButtonField":
			v, err = unmarshalButtonField(data)
		case "CheckboxField":
			v, err = unmarshalCheckboxField(data)
		case "CollaboratorField":
			v, err = unmarshalCollaboratorField(data)
		case "CreatedByField":
			v, err = unmarshalCreatedByField(data)
		case "CreatedTimeField":
			v, err = unmarshalCreatedTimeField(data)
		case "CurrencyField":
			v, err = unmarshalCurrencyField(data)
		case "DateField":
			v, err = unmarshalDateField(data)
		case "DateTimeField":
			v, err = unmarshalDateTimeField(data)
		case "DurationField":
			v, err = unmarshalDurationField(data)
		case "EmailField":
			v, err = unmarshalEmailField(data)
		case "LastModifiedByField":
			v, err = unmarshalLastModifiedByField(data)
		case "LastModifiedTimeField":
			v, err = unmarshalLastModifiedTimeField(data)
		case "LongTextField":
			v, err = unmarshalLongTextField(data)
		case "MultipleCollaboratorField":
			v, err = unmarshalMultipleCollaboratorField(data)
		case "MultipleSelectField":
			v, err = unmarshalMultipleSelectField(data)
		case "IntField":
			v, err = unmarshalIntField(data)
		case "FloatField":
			v, err = unmarshalFloatField(data)
		case "PercentField":
			v, err = unmarshalPercentField(data)
		case "PhoneField":
			v, err = unmarshalPhoneField(data)
		case "RatingField":
			v, err = unmarshalRatingField(data)
		case "RawField":
			v, err = unmarshalRawField(data)
		case "RichTextField":
			v, err = unmarshalRichTextField(data)
		case "SingleLineTextField":
			v, err = unmarshalSingleLineTextField(data)
		case "SingleSelectField":
			v, err = unmarshalSingleSelectField(data)
		case "UrlField":
			v, err = unmarshalUrlField(data)
		}
		if err != nil {
			return fmt.Errorf("error unmarshaling field %s: %w", ft.Name, err)
		}
		if v.IsValid() {
			fv.Set(v)
			continue
		}
		return fmt.Errorf("field %s does not have a known field type", ft.Name)
	}
	return nil
}

func unmarshalAttachmentField(data any) (reflect.Value, error) {
	attachmentsData, ok := data.([]any)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid attachment data: %v", data)
	}
	attachments := make([]AttachmentData, 0, len(attachmentsData))
	for _, attachmentData := range attachmentsData {
		attachment, err := unmarshalAttachmentData(attachmentData)
		if err != nil {
			return reflect.Value{}, err
		}
		attachments = append(attachments, *attachment)
	}
	return reflect.ValueOf(AttachmentField(attachments)), nil
}

func unmarshalBarcodeField(data any) (reflect.Value, error) {
	barcodeData, ok := data.(map[string]any)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid barcode data: %v", data)
	}
	barcode := BarcodeField{
		kind: entryOrZero[string]("type", barcodeData),
		data: entryOrZero[string]("data", barcodeData),
	}
	return reflect.ValueOf(barcode), nil
}

func unmarshalButtonField(data any) (reflect.Value, error) {
	buttonData, ok := data.(map[string]any)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid button data: %v", data)
	}
	label := entryOrZero[string]("label", buttonData)
	if label == "" {
		return reflect.Value{}, fmt.Errorf("button data missing required field 'label'")
	}
	button := ButtonField{
		label: label,
		url:   entryOrZero[string]("url", buttonData),
	}
	return reflect.ValueOf(button), nil
}

func unmarshalCheckboxField(data any) (reflect.Value, error) {
	val, _ := data.(bool)
	return reflect.ValueOf(CheckboxField(val)), nil
}

func unmarshalCollaboratorField(data any) (reflect.Value, error) {
	collaborator, err := unmarshalCollaboratorData(data)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(CollaboratorField(*collaborator)), nil
}

func unmarshalCreatedByField(data any) (reflect.Value, error) {
	collaborator, err := unmarshalCollaboratorData(data)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(CreatedByField(*collaborator)), nil
}

func unmarshalCreatedTimeField(data any) (reflect.Value, error) {
	dt, err := unmarshalDateTimeData(data)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(CreatedTimeField(*dt)), nil
}

func unmarshalCurrencyField(data any) (reflect.Value, error) {
	val, ok := data.(float64)
	if !ok {
		iv, ok := data.(int)
		if !ok {
			return reflect.Value{}, fmt.Errorf("invalid currency data: %v", data)
		}
		val = float64(iv)
	}
	return reflect.ValueOf(CurrencyField(val)), nil
}

func unmarshalDateField(data any) (reflect.Value, error) {
	dt, err := unmarshalDateTimeData(data)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(DateField(*dt)), nil
}

func unmarshalDateTimeField(data any) (reflect.Value, error) {
	dt, err := unmarshalDateTimeData(data)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(DateTimeField(*dt)), nil
}

func unmarshalDurationField(data any) (reflect.Value, error) {
	duration, err := unmarshalDurationData(data)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(DurationField(*duration)), nil
}

func unmarshalEmailField(data any) (reflect.Value, error) {
	val, ok := data.(string)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid email data: %v", data)
	}
	return reflect.ValueOf(EmailField(val)), nil
}

func unmarshalLastModifiedByField(data any) (reflect.Value, error) {
	collaborator, err := unmarshalCollaboratorData(data)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(LastModifiedByField(*collaborator)), nil
}

func unmarshalLastModifiedTimeField(data any) (reflect.Value, error) {
	dt, err := unmarshalDateTimeData(data)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(LastModifiedTimeField(*dt)), nil
}

func unmarshalLongTextField(data any) (reflect.Value, error) {
	val, ok := data.(string)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid long text data: %v", data)
	}
	return reflect.ValueOf(LongTextField(val)), nil
}

func unmarshalMultipleCollaboratorField(data any) (reflect.Value, error) {
	collaborators, err := unmarshalMultipleCollaboratorData(data)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(MultipleCollaboratorField(collaborators)), nil
}

func unmarshalMultipleSelectField(data any) (reflect.Value, error) {
	vals, ok := data.([]any)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid multiple select data: %v", data)
	}
	options := make([]string, 0, len(vals))
	for _, val := range vals {
		option, ok := val.(string)
		if !ok {
			return reflect.Value{}, fmt.Errorf("invalid select option data: %v", val)
		}
		options = append(options, option)
	}
	return reflect.ValueOf(MultipleSelectField(options)), nil
}

func unmarshalIntField(data any) (reflect.Value, error) {
	val, ok := data.(int)
	if !ok {
		fv, ok := data.(float64)
		if !ok {
			return reflect.Value{}, fmt.Errorf("invalid number data: %v", data)
		}
		val = int(fv)
	}
	return reflect.ValueOf(IntField(val)), nil
}

func unmarshalFloatField(data any) (reflect.Value, error) {
	val, ok := data.(float64)
	if !ok {
		iv, ok := data.(int)
		if !ok {
			return reflect.Value{}, fmt.Errorf("invalid number data: %v", data)
		}
		val = float64(iv)
	}
	return reflect.ValueOf(FloatField(val)), nil
}

func unmarshalPercentField(data any) (reflect.Value, error) {
	val, ok := data.(float64)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid number data: %v", data)
	}
	return reflect.ValueOf(PercentField(val)), nil
}

func unmarshalPhoneField(data any) (reflect.Value, error) {
	val, ok := data.(string)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid string data: %v", data)
	}
	return reflect.ValueOf(PhoneField(val)), nil
}

func unmarshalRatingField(data any) (reflect.Value, error) {
	val, ok := data.(int)
	if !ok || val < 1 || val > 10 {
		return reflect.Value{}, fmt.Errorf("invalid rating data: %v", data)
	}
	return reflect.ValueOf(RatingField(val)), nil
}

func unmarshalRawField(data any) (reflect.Value, error) {
	return reflect.ValueOf(RawField(data)), nil
}

func unmarshalRichTextField(data any) (reflect.Value, error) {
	val, ok := data.(string)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid rich text data: %v", data)
	}
	return reflect.ValueOf(RichTextField(val)), nil
}

func unmarshalSingleLineTextField(data any) (reflect.Value, error) {
	val, ok := data.(string)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid single line text data: %v", data)
	}
	return reflect.ValueOf(SingleLineTextField(val)), nil
}

func unmarshalSingleSelectField(data any) (reflect.Value, error) {
	val, ok := data.(string)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid single select data: %v", data)
	}
	return reflect.ValueOf(SingleSelectField(val)), nil
}

func unmarshalUrlField(data any) (reflect.Value, error) {
	val, ok := data.(string)
	if !ok {
		return reflect.Value{}, fmt.Errorf("invalid url data: %v", data)
	}
	return reflect.ValueOf(UrlField(val)), nil
}

func unmarshalAttachmentData(data any) (*AttachmentData, error) {
	attachmentData, ok := data.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("attachment data in incorrect format")
	}
	id := entryOrZero[string]("id", attachmentData)
	if id == "" {
		return nil, fmt.Errorf("attachment data missing required field 'id'")
	}
	attachment := AttachmentData{
		id:       id,
		url:      entryOrZero[string]("url", attachmentData),
		filename: entryOrZero[string]("filename", attachmentData),
		mimeType: entryOrZero[string]("type", attachmentData),
		size:     int64(entryOrZero[float64]("size", attachmentData)),
		height:   int64(entryOrZero[float64]("height", attachmentData)),
		width:    int64(entryOrZero[float64]("width", attachmentData)),
	}
	if thumbnails, ok := attachmentData["thumbnails"]; ok {
		thumbnailsData, err := unmarshalThumbnails(thumbnails)
		if err != nil {
			return nil, err
		}
		attachment.thumbnails = thumbnailsData
	}
	return &attachment, nil
}

func unmarshalThumbnails(data any) (map[string]ThumbnailData, error) {
	thumbnailsData, ok := data.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("thumbnails data in incorrect format")
	}
	thumbnails := make(map[string]ThumbnailData, len(thumbnailsData))
	for thumbnailName, thumbnailData := range thumbnailsData {
		thumbnail, err := unmarshalThumbnail(thumbnailData)
		if err != nil {
			return nil, err
		}
		thumbnails[thumbnailName] = *thumbnail
	}
	return thumbnails, nil
}

func unmarshalThumbnail(data any) (*ThumbnailData, error) {
	thumbnailData, ok := data.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("thumbnail data in incorrect format")
	}
	url := entryOrZero[string]("url", thumbnailData)
	if url == "" {
		return nil, fmt.Errorf("thumbnail data missing required field 'url'")
	}
	thumbnail := ThumbnailData{
		url:    url,
		height: int64(entryOrZero[float64]("height", thumbnailData)),
		width:  int64(entryOrZero[float64]("width", thumbnailData)),
	}
	return &thumbnail, nil
}

func unmarshalMultipleCollaboratorData(data any) ([]CollaboratorData, error) {
	collaboratorsData, ok := data.([]any)
	if !ok {
		return nil, fmt.Errorf("multiple collaborator data in incorrect format")
	}
	collaborators := make([]CollaboratorData, 0, len(collaboratorsData))
	for _, collaboratorData := range collaboratorsData {
		collaborator, err := unmarshalCollaboratorData(collaboratorData)
		if err != nil {
			return nil, err
		}
		collaborators = append(collaborators, *collaborator)
	}
	return collaborators, nil
}

func unmarshalCollaboratorData(data any) (*CollaboratorData, error) {
	collabData, ok := data.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("collaborator data in incorrect format")
	}
	id := entryOrZero[string]("id", collabData)
	if id == "" {
		return nil, fmt.Errorf("collaborator data doesn't have required 'id' value")
	}
	c := CollaboratorData{
		id:              id,
		email:           entryOrZero[string]("email", collabData),
		name:            entryOrZero[string]("name", collabData),
		permissionLevel: entryOrZero[string]("permissionLevel", collabData),
		profilePicUrl:   entryOrZero[string]("profilePicUrl", collabData),
	}
	return &c, nil
}

func unmarshalDateTimeData(data any) (*time.Time, error) {
	s, ok := data.(string)
	if !ok {
		return nil, fmt.Errorf("datetime data in incorrect format")
	}
	// first try full date and time, then just date
	dt, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", s, time.UTC)
	if err != nil {
		dt, err = time.ParseInLocation("2006-01-02", s, time.UTC)
		if err != nil {
			return nil, fmt.Errorf("datetime data in incorrect format: %w", err)
		}
	}
	return &dt, nil
}

func unmarshalDurationData(data any) (*time.Duration, error) {
	seconds, ok := data.(float64)
	if !ok {
		secs, ok := data.(int)
		if !ok {
			return nil, fmt.Errorf("duration data in invalid format")
		}
		seconds = float64(secs)
	}
	return new(time.Duration(seconds * float64(time.Second))), nil
}

func entryOrZero[T any](key string, m map[string]any) T {
	if val, ok := m[key]; ok {
		if v, ok := val.(T); ok {
			return v
		}
	}
	return *new(T)
}

func ensureTarget(target reflect.Value) {
	if target.Kind() == reflect.Ptr && target.IsNil() {
		np := reflect.New(target.Type().Elem())
		target.Set(np)
	}
}
