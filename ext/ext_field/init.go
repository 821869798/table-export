package ext_field

import "github.com/821869798/table-export/ext"

func init() {
	ext.RegisterExtFieldType(NewExtFieldPointInt())
	ext.RegisterExtFieldType(NewExtFieldKVListIntInt())
	ext.RegisterExtFieldType(NewExtFieldKVListIntFloat())
}
