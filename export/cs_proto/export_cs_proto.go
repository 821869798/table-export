package cs_proto

import (
	log "github.com/sirupsen/logrus"
	"table-export/config"
	"table-export/data/model"
	"table-export/export/api"
	"table-export/export/common"
	"table-export/meta"
	"table-export/util"
)

type ExportCsProto struct {
	tableMetas []*meta.RawTableMeta
}

func NewExportCsProto(tableMetas []*meta.RawTableMeta, extraArg map[string]string) api.IExport {
	e := &ExportCsProto{
		tableMetas: tableMetas,
	}
	return e
}

func (e *ExportCsProto) TableMetas() []*meta.RawTableMeta {
	return e.tableMetas
}

func (e *ExportCsProto) Export() {
	csRule := config.GlobalConfig.Meta.RuleCSProto

	//清空目录
	if err := util.ClearDirAndCreateNew(csRule.ProtoTempDir); err != nil {
		log.Fatal(err)
	}

	//实际开始转换
	common.CommonMutilExport(e.tableMetas, func(dataModel *model.TableModel) {
		exportCSProtoFile(dataModel, csRule)
	})
}

func exportCSProtoFile(dataModel *model.TableModel, csRule *config.RawMetaRuleCSProto) {

}
