package meta

import (
	"github.com/821869798/fankit/fanpath"
	"github.com/BurntSushi/toml"
	"github.com/gookit/slog"
	"os"
	"path/filepath"
	"text/template"
)

type RawTableMeta struct {
	Target     string
	Mode       string
	SourceType string            `toml:"source_type"`
	PostScript string            `toml:"post_script"`
	Sources    []*RawTableSource `toml:"sources"`
	Fields     []*RawTableField  `toml:"fields"`
	Checks     []*RawTableCheck  `toml:"checks"`
}

type RawTableSource struct {
	Table string `toml:"file_name"`
	Sheet string `toml:"sheet_name"`
}

type RawTableField struct {
	Active bool   `toml:"active"`
	Source string `toml:"sname"`
	Target string `toml:"tname"`
	Type   string `toml:"type"`
	Key    int    `toml:"key"`
	Desc   string `toml:"desc"`
}

type RawTableCheck struct {
	Active bool   `tomm:"active"`
	Global bool   `toml:"global"`
	Code   string `toml:"code"`
}

func NewRawTableMeta() *RawTableMeta {
	r := &RawTableMeta{}
	return r
}

func NewRawTableField(source, desc string) *RawTableField {
	rtf := &RawTableField{
		Active: false,
		Source: source,
		Target: source,
		Type:   "",
		Key:    0,
		Desc:   desc,
	}
	return rtf
}

func LoadTableMetasByDir(fullPath string) ([]*RawTableMeta, error) {
	fileLists, err := fanpath.GetFileListByExt(fullPath, ".toml")
	if err != nil {
		return nil, err
	}
	tableMetas := make([]*RawTableMeta, len(fileLists), len(fileLists))
	for index, file := range fileLists {
		slog.Debug(file)
		if _, err = toml.DecodeFile(file, &tableMetas[index]); err != nil {
			return nil, err
		}

	}
	return tableMetas, nil
}

func (rtm *RawTableMeta) SaveTableMetaByDir(filePath string) error {
	parentDir := filepath.Dir(filePath)
	//不存在就创建
	if !fanpath.ExistDir(parentDir) {
		_ = os.MkdirAll(parentDir, os.ModePerm)
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	err = toml.NewEncoder(file).Encode(rtm)
	return err
}

func (rtm *RawTableMeta) SaveTableMetaTemplateByDir(filePath string) error {
	tmpl, err := template.New("lua").Parse(`target = "{{.Target}}"
mode = "{{.Mode}}"
source_type = "{{.SourceType}}"
post_script = ""

sources = [
{{range $i, $v := .Sources }}	{ file_name = "{{$v.Table}}",    sheet_name = "{{$v.Sheet}}" },
{{end}}]

fields = [
{{range $i, $v := .Fields }}	{ active = {{$v.Active}},   sname = "{{$v.Source}}" ,      tname = "{{$v.Target}}" ,      type = "{{$v.Type}}" ,  key = {{$v.Key}},    desc = "{{$v.Desc}}" },
{{end}}]

checks = [
]
`)
	if err != nil {
		return err
	}

	parentDir := filepath.Dir(filePath)
	//不存在就创建
	if !fanpath.ExistDir(parentDir) {
		_ = os.MkdirAll(parentDir, os.ModePerm)
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	//渲染输出
	err = tmpl.Execute(file, rtm)
	if err != nil {
		return err
	}
	return nil
}
