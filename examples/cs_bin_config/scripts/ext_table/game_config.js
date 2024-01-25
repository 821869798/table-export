const {eFieldType} = require("../common/table-lib");
const strings = require('../common/strings');

const ArraySplit = tableEngine.TableConfig.ArraySplit;

function PostTable(tableModel,context) {
    // 处理离散全局表的数据，解析成一个一个字段
    tableModel.ClearRecord = true
    const memTable = tableModel.MemTable
    const tableMeta = tableModel.Meta
    const rawDataMap = memTable.RawDataMapping()
    // 英雄等级上限
    const record1 = memTable.GetRecordByKey(1)
    tableMeta.AddExtraField("heroLevelMax",tableEngine.NewTableFieldType(eFieldType.Int),record1.remark)
    memTable.AddExtraData("heroLevelMax",parseInt(record1.data))
    // 初始队伍
    const record2 = memTable.GetRecordByKey(2)
    tableMeta.AddExtraField("heroInitTeam",tableEngine.NewTableFieldArrayType(tableEngine.NewTableFieldType(eFieldType.Int)),record2.remark)
    memTable.AddExtraData("heroInitTeam",strings.trimSplitIntArray(record2.data,ArraySplit))
    // 初始道具
    const record3 = memTable.GetRecordByKey(3)
    tableMeta.AddExtraField("initItems",tableEngine.GetExtFieldType("KVList_IntInt"),record3.remark)
    memTable.AddExtraData("initItems",tableEngine.ParseExtFieldValue("KVList_IntInt",record3.data))
    // 开启debug
    const record4 = memTable.GetRecordByKey(4)
    tableMeta.AddExtraField("openDebug",tableEngine.NewTableFieldType(eFieldType.Bool),record4.remark)
    memTable.AddExtraData("openDebug",strings.stringToBool(record4.data))
}

module.exports = PostTable;