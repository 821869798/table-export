package meta

import (
	"github.com/821869798/table-export/field_type"
)

type TableField struct {
	Source     string
	Target     string
	Type       *field_type.TableFieldType
	TypeString string
	Desc       string
	Key        int
}

func newTableField(rtf *RawTableField) (*TableField, error) {
	tf := &TableField{
		Source:     rtf.Source,
		Target:     rtf.Target,
		TypeString: rtf.Type,
		Desc:       rtf.Desc,
		Key:        rtf.Key,
	}
	tft, err := parseFieldTypeFromString(tf.TypeString)
	if err != nil {
		return nil, err
	}
	tf.Type = tft
	return tf, nil
}
