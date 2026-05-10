package airdata

import (
	"errors"
	"fmt"
	"maps"
	"strings"
	"testing"
)

type Table1 struct {
	RecordId             string
	Attachment           AttachmentField           `field:"Attachment"`
	Barcode              BarcodeField              `field:"Barcode"`
	Button               ButtonField               `field:"Button"`
	Checkbox             CheckboxField             `field:"Checkbox"`
	Collaborator         CollaboratorField         `field:"Collaborator"`
	CreatedBy            CreatedByField            `field:"CreatedBy"`
	CreatedTime          CreatedTimeField          `field:"CreatedTime"`
	Currency             CurrencyField             `field:"Currency"`
	Date                 DateField                 `field:"Date"`
	DateTime             DateTimeField             `field:"DateTime"`
	Duration             DurationField             `field:"Duration"`
	Email                EmailField                `field:"Email"`
	LastModifiedBy       LastModifiedByField       `field:"LastModifiedBy"`
	LastModifiedTime     LastModifiedTimeField     `field:"LastModifiedTime"`
	LongText             LongTextField             `field:"LongText"`
	MultipleCollaborator MultipleCollaboratorField `field:"MultipleCollaborator"`
	MultipleSelect       MultipleSelectField       `field:"MultipleSelect"`
	Int                  IntField                  `field:"Int"`
	Float                FloatField                `field:"Float"`
	Percent              PercentField              `field:"Percent"`
	Phone                PhoneField                `field:"Phone"`
	Rating               RatingField               `field:"Rating"`
	Raw                  RawField                  `field:"Raw"`
	RichText             RichTextField             `field:"RichText"`
	SingleLineText       SingleLineTextField       `field:"SingleLine"`
	SingleSelect         SingleSelectField         `field:"SingleSelect"`
	Url                  UrlField                  `field:"Url"`
	SingleLink           *Table1                   `field:"SingleLink"`
	MultiLink            []*Table1                 `field:"MultiLink"`
}

func (t *Table1) GetRecordId() string {
	return t.RecordId
}

func (t *Table1) RetrieveRecord(id string) (map[string]any, error) {
	for _, record := range table1Records {
		if record.RecordId == id {
			return record.Fields, nil
		}
	}
	return nil, fmt.Errorf("record with id '%s' not found", id)
}

func (t *Table1) Marshal() (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Table1) Unmarshal(fields map[string]any) error {
	return UnmarshalRecord(t, fields, 1)
}

type Table1Ptr struct {
	RecordId             string
	Attachment           *AttachmentField           `field:"Attachment"`
	Barcode              *BarcodeField              `field:"Barcode"`
	Button               *ButtonField               `field:"Button"`
	Checkbox             *CheckboxField             `field:"Checkbox"`
	Collaborator         *CollaboratorField         `field:"Collaborator"`
	CreatedBy            *CreatedByField            `field:"CreatedBy"`
	CreatedTime          *CreatedTimeField          `field:"CreatedTime"`
	Currency             *CurrencyField             `field:"Currency"`
	Date                 *DateField                 `field:"Date"`
	DateTime             *DateTimeField             `field:"DateTime"`
	Duration             *DurationField             `field:"Duration"`
	Email                *EmailField                `field:"Email"`
	LastModifiedBy       *LastModifiedByField       `field:"LastModifiedBy"`
	LastModifiedTime     *LastModifiedTimeField     `field:"LastModifiedTime"`
	LongText             *LongTextField             `field:"LongText"`
	MultipleCollaborator *MultipleCollaboratorField `field:"MultipleCollaborator"`
	MultipleSelect       *MultipleSelectField       `field:"MultipleSelect"`
	Int                  *IntField                  `field:"Int"`
	Float                *FloatField                `field:"Float"`
	Percent              *PercentField              `field:"Percent"`
	Phone                *PhoneField                `field:"Phone"`
	Rating               *RatingField               `field:"Rating"`
	Raw                  *RawField                  `field:"Raw"`
	RichText             *RichTextField             `field:"RichText"`
	SingleLineText       *SingleLineTextField       `field:"SingleLine"`
	SingleSelect         *SingleSelectField         `field:"SingleSelect"`
	Url                  *UrlField                  `field:"Url"`
	SingleLink           *Table1Ptr                 `field:"SingleLink"`
	MultiLink            []*Table1Ptr               `field:"MultiLink"`
}

func (t *Table1Ptr) GetRecordId() string {
	return t.RecordId
}

func (t *Table1Ptr) RetrieveRecord(id string) (map[string]any, error) {
	for _, record := range table1Records {
		if record.RecordId == id {
			return record.Fields, nil
		}
	}
	return nil, fmt.Errorf("record with id '%s' not found", id)
}

func (t *Table1Ptr) Marshal() (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Table1Ptr) Unmarshal(fields map[string]any) error {
	return UnmarshalRecord(t, fields, 1)
}

type Table1Record struct {
	RecordId string
	Fields   map[string]any
}

//goland:noinspection SpellCheckingInspection
var table1Records = []Table1Record{
	{
		RecordId: "recZPad51auofCat0",
		Fields: map[string]any{
			"Attachment": []any{
				map[string]any{
					"filename": "Twilight.pdf",
					"id":       "attyxnHVzxgLVejeF",
					"size":     190671,
					"thumbnails": map[string]any{
						"large": map[string]any{
							"height": 663,
							"url":    "https://v5.airtableusercontent.com/v3/u/52/52/1777687200000/O_nVQMPL9KHMKsp7lZ1rkA/nZ7pfM1EOe07zLExAdY6prY6lfsfRgUXy2CdUxx53tPdwO6cD9UymaFnWIw8-yIhZ4K6-lkfDcghvDve4EmdrEfaB4536ObuJAd8bnrQtjruFCAbSY0cAPEiavC1xxaG6Aa-8Tn0Cu92M8opeH-mNw/-kpAmiNQjfCiFq3aP3eR7NmYMqSRxqk178WGx-4TOjI",
							"width":  512,
						},
						"small": map[string]any{
							"height": 36,
							"url":    "https://v5.airtableusercontent.com/v3/u/52/52/1777687200000/O0K-odzqjLNgGnHVHjLJ0A/iS6joVzlRQBPwi_-CCMn8UfHkem25rJD2-2H58b-picOywaLPu7iL6tA9NDHxm5SBj_eDqYWVpHEw1HpUAToV0IfRG9KQFuxhWWkDTUrBa30D80UVEii9yVxqdhUUFkX2utK94Y5T9tVKY9FDQfU6w/96UXCKy8ANYoMlfTsg_HMJSq-Eyn3Q2PUhOivegfOcI",
							"width":  28,
						},
					},
					"type": "application/pdf",
					"url":  "https://v5.airtableusercontent.com/v3/u/52/52/1777687200000/L9a_UEHe6JpgQ4Ob7irx-Q/VwDwY1qrh2jgAi8DR2VupMs8f1kL6LRLyBgwroLQJaUoYrxRZxcP5xIvHdBlwJkgYUcY6652wYQRN8aN2OdlezGqI0Z1keE2Ie0Y59gDzzLYs9-obtomNJwQmS0iuzl-4D8jgHgXTfQo8PfmzdN4ExH7slGF5_TQade7xYseH38/wpb6BYuZVsddGVlmw-2qO08ETS4y1oTq5WtPRrU2X8Q",
				},
			},
			"Barcode": map[string]any{
				"text": "071701303202",
				"type": "upca",
			},
			"Button": map[string]any{
				"label": "Button",
				"url":   "https://brotsky.com",
			},
			"Checkbox": true,
			"Collaborator": map[string]any{
				"email": "dan@brotsky.com",
				"id":    "usrh3gJSuNh5I0l2T",
				"name":  "Dan Brotsky",
			},
			"CreatedBy": map[string]any{
				"email": "dev@brotsky.com",
				"id":    "usraydLTb610fi477",
				"name":  "Daniel Brotsky",
			},
			"CreatedTime": "2026-04-29T20:38:33.000Z",
			"Currency":    750.4587,
			"Date":        "2026-05-02",
			"DateTime":    "2026-05-02T21:30:00.000Z",
			"Duration":    594,
			"Email":       "foo%bar@zotz",
			"Float":       600,
			"Int":         600,
			"LastModifiedBy": map[string]any{
				"email": "dev@brotsky.com",
				"id":    "usraydLTb610fi477",
				"name":  "Daniel Brotsky",
			},
			"LastModifiedTime": "2026-05-01T23:29:07.000Z",
			"LongText":         "This has three lines.\nThis is the second line.\nThis is the third line.",
			"MultipleCollaborator": []any{
				map[string]any{
					"email": "dan@brotsky.com",
					"id":    "usrh3gJSuNh5I0l2T",
					"name":  "Dan Brotsky",
				},
				map[string]any{
					"email": "dev@brotsky.com",
					"id":    "usraydLTb610fi477",
					"name":  "Daniel Brotsky",
				},
			},
			"MultipleSelect": []any{
				"SelectOption1",
				"SelectOption2",
				"SelectOption3",
			},
			"Percent": -0.04,
			"Phone":   "+45 (698) 2000-34",
			"Rating":  1,
			"Raw": []any{
				"recfhYBZ9vM0r7cvs",
				"recZPad51auofCat0",
			},
			"RichText":     "Without\n",
			"SingleLine":   "Record3",
			"SingleSelect": "Todo",
			"Url":          "https://brotsky.com",
			"SingleLink":   []any{},
			"MultiLink":    []any{"recfhYBZ9vM0r7cvs", "rechSfyOe7jT3Ub5a"},
		},
	},
	{
		RecordId: "recfhYBZ9vM0r7cvs",
		Fields: map[string]any{
			"Attachment": []any{
				map[string]any{
					"filename": "Anker PowerPort III info.jpeg",
					"height":   1341,
					"id":       "attZXkEkGf1sIyKtR",
					"size":     710238,
					"thumbnails": map[string]any{
						"full": map[string]any{
							"height": 1341,
							"url":    "https://v5.airtableusercontent.com/v3/u/52/52/1777687200000/uYxXPN3E9hDBKl2hWbokvA/4c7-NSYG4E3nIJMHb5jc81HuyGoMILeHksCbLgoRJQ9bRTeDSdRwFJaUwTnNhLQdnZaVyhIueVC2JoUj_P25j9zCcBDdhFNGLdEtryBP1fWAF1aWnI3uKdZSSSJKSCohf1NbmaDTem5baCtB9iX10g/2tHXHWb-duIDO8pYUfmbSgbHckCCDjEzOGAgaYZnmT0",
							"width":  2294,
						},
						"large": map[string]any{
							"height": 512,
							"url":    "https://v5.airtableusercontent.com/v3/u/52/52/1777687200000/7bToiypueJ0--tvqH2wkJA/ZJNL76nheLsEWqLObQD_r0xsQ5yqMBLUjHbERq3jA7The7zMtnj0CpDkJZTVbvopc9OZWR1VYhSuRIOXq7nl37zDBSKMn97if0tTk-tgna4Zlfx5I6lc9XMuGe3lwHiQzAV0wdQ9tyYmL6rcHKMPUg/KfgiyDO09Wgbia4a8vr1ouumEcbX3HYjFvwjO4b9Vxc",
							"width":  876,
						},
						"small": map[string]any{
							"height": 36,
							"url":    "https://v5.airtableusercontent.com/v3/u/52/52/1777687200000/I9xqIuvQaXjhTUiyeLeTYw/n5ouLHFHEAKg4uOclSycPrrRphh8A4iphJMbZEKUNRznLk29OJgJqOHX6ks48MUfwWQW3eLRXIaa7rHoFJUrz1NDrGWqthrM-G35zFakEpUP-IsBW9NiUGi2-iCJGgVALWtEo8MmUa42Do1rbLlq8Q/WJWPl_DyQDuazTyDw2FhWdSCC1ivPWQS0dDC-DzQL3Y",
							"width":  62,
						},
					},
					"type":  "image/jpeg",
					"url":   "https://v5.airtableusercontent.com/v3/u/52/52/1777687200000/i4NqrJgiyb0d3GwU2kG1Sw/tGyGKmAztuf6JAqwhW1y2FZd5QaMcRdVXX41Gl9fd-LIFA3Ncr0IAKMUZHNAmzinflwsLYBjH-QzZtADC-alsrt9i_vKyt81WkHi7TBlIQmrQ5cs3sRABwaENrGCI_R6XohE3EXkCkI7NF2V2uTKNM7VDwbZX25KzIC0qACvYLM810mvHN4gmquwIIU-dBQv/ajc0pOaBznMVaDau-h95rUBU0kliRWF4LNux6WHU3lA",
					"width": 2294,
				},
			},
			"Button": map[string]any{
				"label": "Button",
				"url":   "https://clickonetwo.io",
			},
			"Collaborator": map[string]any{
				"email": "dev@brotsky.com",
				"id":    "usraydLTb610fi477",
				"name":  "Daniel Brotsky",
			},
			"CreatedBy": map[string]any{
				"email": "dev@brotsky.com",
				"id":    "usraydLTb610fi477",
				"name":  "Daniel Brotsky",
			},
			"CreatedTime": "2026-04-29T20:38:33.000Z",
			"Currency":    365.24,
			"Date":        "2026-04-30",
			"DateTime":    "2026-05-01T22:47:00.000Z",
			"Duration":    5034.307,
			"Email":       "this is a tewst",
			"Float":       3.5,
			"Int":         3,
			"LastModifiedBy": map[string]any{
				"email": "dev@brotsky.com",
				"id":    "usraydLTb610fi477",
				"name":  "Daniel Brotsky",
			},
			"LastModifiedTime": "2026-05-01T23:25:07.000Z",
			"LongText":         "This is short.",
			"MultipleCollaborator": []any{
				map[string]any{
					"email": "dev@brotsky.com",
					"id":    "usraydLTb610fi477",
					"name":  "Daniel Brotsky",
				},
				map[string]any{
					"email": "dan@brotsky.com",
					"id":    "usrh3gJSuNh5I0l2T",
					"name":  "Dan Brotsky",
				},
			},
			"MultipleSelect": []any{
				"SelectOption1",
			},
			"Percent": 0.34,
			"Phone":   "(510) 926-0499",
			"Rating":  3,
			"Raw": []any{
				"recfhYBZ9vM0r7cvs",
				"rechSfyOe7jT3Ub5a",
			},
			"RichText":     "With \\*\\*bold\\*\\*\n",
			"SingleLine":   "Record1",
			"SingleSelect": "In progress",
			"Url":          "https://clickonetwo.io",
			"SingleLink":   []any{"recfhYBZ9vM0r7cvs"},
			"MultiLink":    []any{},
		},
	},
	{
		RecordId: "rechSfyOe7jT3Ub5a",
		Fields: map[string]any{
			"Barcode": map[string]any{
				"text": "888462315968",
				"type": "upca",
			},
			"Button": map[string]any{
				"label": "Button",
				"url":   "",
			},
			"Checkbox": true,
			"CreatedBy": map[string]any{
				"email": "dev@brotsky.com",
				"id":    "usraydLTb610fi477",
				"name":  "Daniel Brotsky",
			},
			"CreatedTime": "2026-04-29T20:38:33.000Z",
			"Currency":    234,
			"Date":        "2026-05-01",
			"DateTime":    "2026-04-30T22:47:00.000Z",
			"Duration":    34.1,
			"Email":       "valid@format.com",
			"Float":       -6.9,
			"Int":         -35,
			"LastModifiedBy": map[string]any{
				"email": "dev@brotsky.com",
				"id":    "usraydLTb610fi477",
				"name":  "Daniel Brotsky",
			},
			"LastModifiedTime": "2026-05-01T23:28:18.000Z",
			"LongText":         "This has two lines.\nThis is the second line.",
			"MultipleSelect": []any{
				"SelectOption1",
				"SelectOption2",
			},
			"Percent":    5.98,
			"Phone":      "crap oh crap",
			"Rating":     5,
			"RichText":   "With \\_italics\\_\n",
			"SingleLine": "Record2",
			"SingleLink": []any{"recZPad51auofCat0"},
			"MultiLink":  []any{"recfhYBZ9vM0r7cvs", "rechSfyOe7jT3Ub5a"},
		},
	},
	{
		RecordId: "bad",
		Fields:   map[string]any{"Percent": "data"},
	},
}

func TestUnmarshalRecord(t *testing.T) {
	for _, record := range table1Records {
		if record.RecordId == "bad" {
			continue
		}
		t.Run("Table1:"+record.RecordId, func(t *testing.T) {
			target := new(Table1)
			err := UnmarshalRecord(target, record.Fields, 1)
			if err != nil {
				t.Errorf("Failed to unmarshal record %s: %v", record.RecordId, err)
				return
			}
			target.RecordId = record.RecordId
		})
		t.Run("Table1Ptr:"+record.RecordId, func(t *testing.T) {
			target := new(Table1Ptr)
			err := UnmarshalRecord(target, record.Fields, 1)
			if err != nil {
				t.Errorf("Failed to unmarshal record %s: %v", record.RecordId, err)
				return
			}
			target.RecordId = record.RecordId
		})
	}
}

type BadTable1 Table1

func (b *BadTable1) GetRecordId() string {
	return "bad"
}

func (b *BadTable1) RetrieveRecord(id string) (map[string]any, error) {
	if id == "bad" {
		for _, record := range table1Records {
			if record.RecordId == id {
				return record.Fields, nil
			}
		}
	}
	return nil, errors.New("not found")
}

func (b *BadTable1) Marshal() (map[string]any, error) {
	panic("implement me")
}

func (b *BadTable1) Unmarshal(fields map[string]any) error {
	return UnmarshalRecord(b, fields, 1)
}

//goland:noinspection SpellCheckingInspection
func TestUnmarshalInvalidRecordData(t *testing.T) {
	bad1 := maps.Clone(table1Records[0].Fields)
	test1 := func(name string, data interface{}) {
		t.Helper()
		previous := bad1[name]
		bad1[name] = data
		if err := UnmarshalRecord(new(BadTable1), bad1, 1); !errors.Is(err, UnmarshalError{}) {
			t.Errorf("Expected UnmarshalError, got %T", err)
		} else if !strings.Contains(err.Error(), name) {
			t.Errorf("Expected error message to contain field name '%s', got %v", name, err)
		} else {
			//t.Logf("Message for %s: %v", name, err)
		}
		bad1[name] = previous
	}
	for name := range bad1 {
		switch name {
		case "Attachment":
			test1("Attachment", "not map[string]any")
			badAttach := []any{
				"not a map[string]any",
			}
			test1("Attachment", badAttach)
			noId := []any{
				map[string]any{
					"filename": "Twilight.pdf",
					"size":     190671,
				},
			}
			test1("Attachment", noId)
			badThumbnails := []any{
				map[string]any{
					"filename":   "Twilight.pdf",
					"id":         "attyxnHVzxgLVejeF",
					"size":       190671,
					"thumbnails": "not a map[string]any",
				},
			}
			test1("Attachment", badThumbnails)
			badThumbnail := []any{
				map[string]any{
					"filename": "Twilight.pdf",
					"id":       "attyxnHVzxgLVejeF",
					"size":     190671,
					"thumbnails": map[string]any{
						"full": "not a map[string]any",
					},
				},
			}
			test1("Attachment", badThumbnail)
			thumbNoId := []any{
				map[string]any{
					"filename": "Twilight.pdf",
					"id":       "attyxnHVzxgLVejeF",
					"size":     190671,
					"thumbnails": map[string]any{
						"full": map[string]any{
							"height": 1341,
							"width":  1024,
						},
					},
				},
			}
			test1("Attachment", thumbNoId)
		case "Barcode":
			test1("Barcode", "not map[string]any")
		case "Button":
			test1("Button", "not map[string]any")
			noId := map[string]any{
				"url": "https://brotsky.com",
			}
			test1("Button", noId)
		case "Checkbox":
			test1("Checkbox", "not bool")
		case "Collaborator":
			test1("Collaborator", "not map[string]any")
		case "CreatedBy":
			test1("CreatedBy", "not map[string]any")
		case "CreatedTime":
			test1("CreatedTime", 34)
		case "Currency":
			test1("Currency", "not float64")
		case "Date":
			test1("Date", 34)
		case "DateTime":
			test1("DateTime", 34)
			test1("DateTime", "not a date")
		case "Duration":
			test1("Duration", "not float64")
		case "Email":
			test1("Email", 34)
		case "Float":
			test1("Float", "not float64")
		case "Int":
			test1("Int", "not int")
		case "LastModifiedBy":
			test1("LastModifiedBy", "not map[string]any")
		case "LastModifiedTime":
			test1("LastModifiedTime", 34)
		case "LongText":
			test1("LongText", 34)
		case "MultipleCollaborator":
			test1("MultipleCollaborator", "not []any")
			badData := []any{
				"not a map[string]any",
			}
			test1("MultipleCollaborator", badData)
			badCollaborator := []any{
				map[string]any{
					"email": "dan@brotsky.com",
					"name":  "Dan Brotsky",
				},
			}
			test1("MultipleCollaborator", badCollaborator)
		case "MultipleSelect":
			test1("MultipleSelect", "not []any")
			badOption := []any{
				"not a string",
				34,
			}
			test1("MultipleSelect", badOption)
		case "Percent":
			test1("Percent", "not float64")
		case "Phone":
			test1("Phone", 34)
		case "Rating":
			test1("Rating", "not int")
		case "Raw":
			continue
		case "RichText":
			test1("RichText", 34)
		case "SingleLine":
			test1("SingleLine", 34)
		case "SingleSelect":
			test1("SingleSelect", 34)
		case "Url":
			test1("Url", 34)
		case "SingleLink":
			test1("SingleLink", "not []any")
			test1("SingleLink", []any{"not a valid record id", "anotherLink"})
			test1("SingleLink", []any{34})
		case "MultiLink":
			test1("MultiLink", "not []any")
			test1("MultiLink", []any{"bad"})
			test1("MultiLink", []any{"not a valid record id", "anotherLink"})
		default:
			t.Errorf("Unexpected field name '%s'", name)
		}
	}
}

type BadTable2 struct {
	Field1 string `field:"Field1"`
}

func (b *BadTable2) GetRecordId() string {
	panic("implement me")
}

func (b *BadTable2) RetrieveRecord(string) (map[string]any, error) {
	panic("implement me")
}

func (b *BadTable2) Marshal() (map[string]any, error) {
	panic("implement me")
}

func (b *BadTable2) Unmarshal(fields map[string]any) error {
	return UnmarshalRecord(b, fields, 1)
}

type BadTable3 struct {
	Field1 *string `field:"Field1"`
}

func (b *BadTable3) GetRecordId() string {
	panic("implement me")
}

func (b *BadTable3) RetrieveRecord(string) (map[string]any, error) {
	panic("implement me")
}

func (b *BadTable3) Marshal() (map[string]any, error) {
	panic("implement me")
}

func (b *BadTable3) Unmarshal(fields map[string]any) error {
	return UnmarshalRecord(b, fields, 1)
}

func TestUnmarshalInvalidField(t *testing.T) {
	if err := UnmarshalRecord(new(BadTable2), map[string]any{"Field1": "value"}, 1); !errors.Is(err, UnmarshalError{}) {
		t.Errorf("Expected UnmarshalError for invalid field, got %v", err)
	} else if !strings.Contains(err.Error(), "Field1") {
		t.Errorf("Expected error message to contain 'Field1', got %v", err)
	}
	if err := UnmarshalRecord(new(BadTable3), map[string]any{"Field1": "value"}, 1); !errors.Is(err, UnmarshalError{}) {
		t.Errorf("Expected UnmarshalError for invalid field, got %v", err)
	} else if !strings.Contains(err.Error(), "Field1") {
		t.Errorf("Expected error message to contain 'Field1', got %v", err)
	}
}

func TestUnmarshalInvalidType(t *testing.T) {
	fields := table1Records[0].Fields
	if err := UnmarshalRecord(nil, fields, 1); !errors.Is(err, UnmarshalError{}) {
		t.Errorf("Expected UnmarshalError, got %T", err)
	} else if !strings.Contains(err.Error(), "nil") {
		t.Errorf("Expected error message to contain 'nil', got %v", err)
	}
	if err := UnmarshalRecord("test", fields, 1); !errors.Is(err, UnmarshalError{}) {
		t.Errorf("Expected UnmarshalError, got %T", err)
	} else if !strings.Contains(err.Error(), "pointer") {
		t.Errorf("Expected error message to contain 'pointer', got %v", err)
	}
	if err := UnmarshalRecord(&fields, fields, 1); !errors.Is(err, UnmarshalError{}) {
		t.Errorf("Expected UnmarshalError, got %v", err)
	} else if !strings.Contains(err.Error(), "struct") {
		t.Errorf("Expected error message to contain 'struct', got %v", err)
	}
	var p *Table1
	if err := UnmarshalRecord(p, fields, 1); !errors.Is(err, UnmarshalError{}) {
		t.Errorf("Expected UnmarshalError, got %v", err)
	} else if !strings.Contains(err.Error(), "valid struct") {
		t.Errorf("Expected error message to contain 'valid struct', got %v", err)
	}
	if err := UnmarshalRecord(&struct{}{}, fields, 1); !errors.Is(err, UnmarshalError{}) {
		t.Errorf("Expected UnmarshalError, got %v", err)
	} else if !strings.Contains(err.Error(), "RecordData") {
		t.Errorf("Expected error message to contain 'RecordData', got %v", err)
	}
}
