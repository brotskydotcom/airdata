package airdata

import (
	"fmt"
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
	SingleLineText       SingleLineTextField       `field:"SingleLineText"`
	SingleSelect         SingleSelectField         `field:"SingleSelect"`
	Url                  UrlField                  `field:"Url"`
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
	return UnmarshalRecord(t, fields)
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
	SingleLineText       *SingleLineTextField       `field:"SingleLineText"`
	SingleSelect         *SingleSelectField         `field:"SingleSelect"`
	Url                  *UrlField                  `field:"Url"`
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
	return UnmarshalRecord(t, fields)
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
			"From field: Table1": []any{
				"recZPad51auofCat0",
			},
			"Int": 600,
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
			"From field: Table1": []any{
				"recZPad51auofCat0",
				"recfhYBZ9vM0r7cvs",
			},
			"Int": 3,
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
			"From field: Table1": []any{
				"recfhYBZ9vM0r7cvs",
			},
			"Int": -35,
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
		},
	},
}

func TestUnmarshalRecord(t *testing.T) {
	for _, record := range table1Records {
		t.Run("Table1:"+record.RecordId, func(t *testing.T) {
			target := new(Table1)
			err := UnmarshalRecord(target, record.Fields)
			if err != nil {
				t.Errorf("Failed to unmarshal record %s: %v", record.RecordId, err)
				return
			}
			target.RecordId = record.RecordId
		})
		t.Run("Table1Ptr:"+record.RecordId, func(t *testing.T) {
			target := new(Table1Ptr)
			err := UnmarshalRecord(target, record.Fields)
			if err != nil {
				t.Errorf("Failed to unmarshal record %s: %v", record.RecordId, err)
				return
			}
			target.RecordId = record.RecordId
		})
	}
}

func TestUnmarshalCollaboratorData(t *testing.T) {
	good := map[string]any{"email": "dan@brotsky.com", "id": "usrh3gJSuNh5I0l2T", "name": "Dan Brotsky"}
	c, err := unmarshalCollaboratorData(good)
	if err != nil {
		t.Fatal(err)
	}
	if c.email != "dan@brotsky.com" {
		t.Errorf("Expected email to be 'dan@brotsky.com', got %s", c.email)
	}
	if c.id != "usrh3gJSuNh5I0l2T" {
		t.Errorf("Expected id to be 'usrh3gJSuNh5I0l2T', got %s", c.id)
	}
	if c.name != "Dan Brotsky" {
		t.Errorf("Expected name to be 'Dan Brotsky', got %s", c.name)
	}
	if c.permissionLevel != "" {
		t.Errorf("Expected permissionLevel to be '', got %s", c.permissionLevel)
	}
	if c.profilePicUrl != "" {
		t.Errorf("Expected profilePicUrl to be '', got %s", c.profilePicUrl)
	}
	var bad1 map[string]any = nil
	if _, err := unmarshalCollaboratorData(bad1); err == nil {
		t.Error("Expected error unmarshaling nil")
	}
	bad2 := map[string]string{"a": "a"}
	if _, err := unmarshalCollaboratorData(bad2); err == nil {
		t.Error("Expected error unmarshaling wrong type of map")
	}
	bad3 := map[string]any{"email": "dan@brotsky.com"}
	if _, err := unmarshalCollaboratorData(bad3); err == nil {
		t.Error("Expected error unmarshaling map without id field")
	}
}
