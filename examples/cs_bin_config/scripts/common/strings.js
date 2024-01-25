const EmptyArray = [];

const strings = {
    trimSplit(origin, sep) {
        origin = origin.trim();
        if (origin.length === 0) {
            return EmptyArray
        }
        return origin.split(sep);
    },
    trimSplitIntArray(origin, sep) {
        let array = this.trimSplit(origin, sep);
        for (let i = 0; i < array.length; i++) {
            array[i] = parseInt(array[i]);
        }
        return array;
    },
    trimSplitFloatArray(origin, sep) {
        let array = this.trimSplit(origin, sep);
        for (let i = 0; i < array.length; i++) {
            array[i] = parseFloat(array[i]);
        }
        return array;
    },

    stringToBool(str) {
        switch (str.toLowerCase().trim()) {
            case "true":
            case "1":
                return true;
            case "false":
            case "0":
            case "":
                return false;
            default:
                // 如果输入不是预期的字符串，可以抛出错误或返回null/undefined
                throw new Error("Invalid string for conversion to boolean");
        }
    }

}

module.exports = strings