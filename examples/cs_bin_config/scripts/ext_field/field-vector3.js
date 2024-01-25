const {eFieldType} = require('../common/table-lib');
const strings = require('../common/strings');

const TypeName = "TestVector3";

var ftc = tableEngine.NewTableFieldClass(TypeName);
ftc.AddField("x", tableEngine.NewTableFieldType(eFieldType.Float));
ftc.AddField("y", tableEngine.NewTableFieldType(eFieldType.Float));
ftc.AddField("z", tableEngine.NewTableFieldType(eFieldType.Float));

const fieldType = tableEngine.NewTableFieldClassType(ftc);

const ArraySplit = tableEngine.TableConfig.ArraySplit;

const extFieldVector3 = {
    Name() {
        return TypeName;
    },
    DefineFile() {
        return "CfgMath";
    },
    TableFieldType() {
        return fieldType;
    },
    ParseOriginData(origin, context) {
        const stringArray = strings.trimSplit(origin,ArraySplit);
        if (stringArray.length === 0) {
            return { x : 0, y: 0, z: 0};
        }
        if (stringArray.length !== 3) {
            context.Error = "Vector3 Parse Error,param count must be 3 or 0: " + origin;
            //console.error("js error:" + origin);
            return null
        }
        const x = parseFloat(stringArray[0]);
        const y = parseFloat(stringArray[1]);
        const z = parseFloat(stringArray[2]);
        return { x : x, y: y, z: z};
    }
}

module.exports = extFieldVector3;