package meta

type TableField struct {
	Source     string
	Target     string
	Type       *TableFieldType
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
	tft, err := getFieldTypeFromString(tf.TypeString)
	if err != nil {
		return nil, err
	}
	tf.Type = tft
	return tf, nil
}
