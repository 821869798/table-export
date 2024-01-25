const {eFieldType} = require("../common/table-lib");
const strings = require('../common/strings');

const ArraySplit = tableEngine.TableConfig.ArraySplit;

function PostGlobal(tableMap,context) {
    console.log("tableCount:" + tableEngine.len(tableMap));

    // test code
    const gameConfig = tableEngine.getMapValue(tableMap,"GameConfig")
    console.log("heroLevelMax:" + gameConfig.MemTable.GetExtraDataMap()["heroLevelMax"]);

    const baseTest = tableEngine.getMapValue(tableMap,"base_test")
    console.log("base_test[1]:" + tableEngine.toJson(gameConfig.MemTable.GetRecordByKey(1)));

    tableEngine.forEach(tableMap,(value, key) => {
        //console.log(key);
    });
}

module.exports = PostGlobal;