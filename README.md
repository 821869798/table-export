# table-export

#### 介绍
table-export 是一个游戏导表工具，目前支持源数据excel和csv，导出格式支持C# Proto，lua，json，可以根据自己的需求自己修改扩展。

#### 使用方法

**工具的配置文件目录为./conf/config.toml**

**测试时可以把build出来的执行文件放在examples目录下**

**导出表meta配置文件**

table_export -g="目标名字,excel(.xlsx)文件名字不带后缀,Sheet名"

例如：table_export -g="base_test,base_test,Sheet1"

导出meta模板后，会放在temp_meta中，需要修改后放入对应的目标，例如cs的放在cs_meta中

**导表**

table_export.exe -m="模式1|模式2|..."

例如 :  table_export.exe -m="lua|cs_proto"