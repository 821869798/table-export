package cs_bin

const template_CS_Table string = `{{ .CodeHead }}
using System.Collections.Generic;

namespace {{.NameSpace}} 
{

	public partial class {{.TableClassName}}
	{

		private readonly {{.KeyDefTypeMap}} _dataMap;
        private readonly List<{{.RecordClassName}}> _dataList;
 
		
  		public {{.TableClassName}}(ByteBuf _buf)
        {
			//first read common data
            {{.TableCommonName}} _commonData = null;
            var commonSize = _buf.ReadSize();
            if( commonSize > 0)
            {
                _commonData = new {{.TableCommonName}}(_buf);
            }

			var size = _buf.ReadSize();
            _dataMap = new {{.KeyDefTypeMap}}(16);
 			_dataList = new List<{{.RecordClassName}}>(size);

  			for (int i = 0; i < size; i++)
            {
                {{.RecordClassName}} _v;
                _v = {{.RecordClassName}}.Deserialize{{.RecordClassName}}(_buf, _commonData);
                _dataList.Add(_v);
{{UniqueMapAssignment $ "_dataMap" "_v" 4}}
            }

            PostInit();
		}

		public int DataCount => _dataList.Count;
		public {{.KeyDefTypeMap}} DataMap => _dataMap;
		public List<{{.RecordClassName}}> DataList => _dataList;
		
{{UniqueMapGetFunc $ "_dataMap" 2}} 

{{UniqueMapGetFuncWithoutError $ "_dataMap" 2}} 
		
        /// <summary>
        /// post process table
        /// </summary>
		partial void PostInit();

	}

	public partial class {{.RecordClassName}} 
	{
		private {{.RecordClassName}}(ByteBuf _buf, {{.TableCommonName}} _commonData)
        {
{{ range $v := .CSCodeWriteFields }}
			{{CSbinFieldReaderEx $v "_buf" "_commonData._field"}}{{ end }}
            PostInit();
        }

        internal static {{.RecordClassName}} Deserialize{{.RecordClassName}}(ByteBuf _buf, {{.TableCommonName}} _commonData)
        {
            return new {{.RecordClassName}}(_buf, _commonData);
        }
{{ range $v := .CSCodeWriteFields }}
        /// <summary>
        /// {{$v.Desc}}
        /// </summary>
		public {{$v.TypeDef}} {{$v.Name}} { get; private set; }
{{ end }}

        /// <summary>
        /// post process table
        /// </summary>
        partial void PostInit();
	}

    /// <summary>
    /// internal common data optimize
    /// </summary>
    internal class {{.TableCommonName}}
    {
{{ if .TableOptimize -}}
{{ range $index,$v := .TableOptimize.OptimizeFields }}
        internal {{GetOutputDefTypeValue $v.OptimizeType $.CollectionReadonly}} _field{{$index}} { get; private set; }{{ end }}
{{- end }}
        internal {{.TableCommonName}}(ByteBuf _buf)
        {
{{ if .TableOptimize -}}
{{ range $index,$v := .TableOptimize.OptimizeFields }}
			{{ $combine := (printf "_field%v" $index) }}{{CSbinFieldReader $v.OptimizeType $combine "_buf"}}{{ end }}
{{- end }}
        }

    }

}
`

const template_CS_Enum = `
namespace {{.NameSpace}}
{
{{ range $v := .Enums }}
    public enum {{$v.Name}}
    {
{{- range $enumValue := $v.Values }}
        // {{$enumValue.Name}} : {{$enumValue.Desc}}
        {{$enumValue.Name}} = {{$enumValue.Index}},
{{- end }}
    }
{{ end }}
}
`

// 废弃
const templateCSCodeNew = `
{{.CodeHead}}
using System.Collections.Generic;

namespace {{.NameSpace}}
{
    public partial class {{.TableClassName}}
    {
        private readonly {{UniqueMapDef}} _dataMap;
        private readonly List<{{RecordClassName}}> _dataList;
{{~if has_unique~}}

        /// <summary>
        /// Multiple unique key map,easy to traverse
        /// </summary>
        private readonly {{unique_map_def_rw}} _uniqueDataMap;
{{~end~}}

        public {{.TableClassName}}(ByteBuf _buf)
        {
            //first read common data
            {{tb_common_name}} _commonData = null;
            var commonSize = _buf.ReadSize();
            if( commonSize > 0)
            {
                _commonData = new {{tb_common_name}}(_buf);
            }

            //then read row data
            var size = _buf.ReadSize();
            _dataMap = new Dictionary<{{major_key_type}}, {{RecordClassName}}>(size * 3 / 2);
            _dataList = new List<{{RecordClassName}}>(size);
{{~if has_unique~}}
            _uniqueDataMap = new {{unique_map_def_rw}}(16);
{{~end~}}

            for (int i = 0; i < size; i++)
            {
                {{RecordClassName}} _v;
                _v = {{RecordClassName}}.Deserialize{{RecordClassName}}(_buf, _commonData);
                _dataList.Add(_v);
                _dataMap.Add(_v.{{major_key_name}}, _v);
{{~if has_unique~}}
                {{cs_unique_key_map_assignment unique_keys '_uniqueDataMap' '_v' RecordClassName}}
{{~end~}}
            }

            PostInit();
        }

        public {{def_map}}<int, {{.RecordClassName}}> DataMap => _dataMap;
        public {{def_list}}<{{.RecordClassName}}> DataList => _dataList;

        public {{.RecordClassName}} GetOrDefault(int key) => _dataMap.TryGetValue(key, out var v) ? v : null;
        public {{.RecordClassName}} Get(int key) => _dataMap[key];
        public {{.RecordClassName}} this[int key] => _dataMap[key];
{{~if has_unique > 0~}}

        public {{unique_map_def_rw}} UniqueKeyDataMap => _uniqueDataMap;
        {{cs_unique_key_map_get_func unique_keys '_uniqueDataMap' .RecordClassName}}
{{~end~}}

        /// <summary>
        /// post process table
        /// </summary>
        partial void PostInit();

    }

    public partial class {{.RecordClassName}}
    {
        private {{.RecordClassName}}(ByteBuf _buf, {{tb_common_name}} _commonData)
        {
{{~ for field in fields~}}
            {{cs_bin_field_reader_ex field optimize '_buf' '_commonData._field'}}
{{~end~}}

            PostInit();
        }

        internal static {{.RecordClassName}} Deserialize{{.RecordClassName}}(ByteBuf _buf, {{tb_common_name}} _commonData)
        {
            return new {{.RecordClassName}}(_buf, _commonData);
        }

        /// <summary>
        /// major key
        /// </summary>
        public {{major_key_type}} {{major_key_name}} { get; private set; }

{{~ for field in fields offset:1~}}
        /// <summary>
        /// {{field.desc}}
        /// </summary>
        public {{cs_get_class_field_def field.type collection_readonly}} {{field.name}} { get; private set; }
{{~end~}}

        /// <summary>
        /// post process table
        /// </summary>
        partial void PostInit();
    }

    /// <summary>
    /// internal common data optimize
    /// </summary>
    internal class {{TableCommonName}}
    {

{{~if optimize_fields != null ~}}
{{~ for opt_field in optimize_fields ~}}
        internal {{cs_get_class_field_def opt_field.optimize_type collection_readonly}} _field{{for.index}} { get; private set; }
{{~end~}}
{{~end~}}
        internal {{TableCommonName}}(ByteBuf _buf)
        {
{{~ for opt_field in optimize_fields~}}
            {{cs_bin_field_reader opt_field.optimize_type '_buf' ('_field' + for.index)}}
{{~end~}}
        }

    }
}
`
