package cs_proto

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"table-export/config"
	"table-export/data/model"
	"table-export/export/api"
	"table-export/export/common"
	"table-export/meta"
	"table-export/util"
	"time"
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
	if err := util.ClearDirAndCreateNew(csRule.BytesDir); err != nil {
		log.Fatal(err)
	}
	if err := util.ClearDirAndCreateNew(csRule.ProtoCSdir); err != nil {
		log.Fatal(err)
	}

	defer util.TimeCost(time.Now(), "export cs_proto time cost = %v\n")

	targetFiles := make([]string, 0, len(e.tableMetas))

	//实际开始转换
	common.CommonMutilExport(e.tableMetas, func(dataModel *model.TableModel) {
		exportCSProtoFile(dataModel, csRule)
		filePath := dataModel.Meta.Target + ".proto"
		targetFiles = append(targetFiles, filePath)
	})

	//创建.cs文件
	if len(targetFiles) > 0 {
		args := []string{"--csharp_out=" + filepath.Join(config.GPath.AbsExeDir(csRule.ProtoCSdir)),
			"--proto_path=" + filepath.Join(config.GPath.AbsExeDir(csRule.ProtoTempDir))}
		args = append(args, targetFiles...)
		var execPath string
		if runtime.GOOS == "windows" {
			execPath = filepath.Join(config.GPath.AbsExeDir(csRule.ProtoCWinDir))
		} else {
			execPath = filepath.Join(config.GPath.AbsExeDir(csRule.ProtoCMacDir))
		}
		protoc := exec.Command(execPath, args...)
		protoc.Stdout = os.Stdout
		protoc.Stderr = os.Stderr
		if err := protoc.Run(); err != nil {
			log.Fatalf("export cs_proto csharp file error:", err)
		}

	}

}

func exportCSProtoFile(dataModel *model.TableModel, csRule *config.RawMetaRuleCSProto) {
	pfd, err := buildProtoFile(dataModel, csRule)
	if err != nil {
		log.Fatal(err)
	}

	buildProtoBytesFile(dataModel, csRule, pfd)

}
