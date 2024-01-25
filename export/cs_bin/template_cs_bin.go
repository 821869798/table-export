package cs_bin

const templateCSCodeTable string = `{{ .CodeHead }}
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
{{ range $v := .CSCodeExtraFields }}
			{{CSbinFieldReader $v.Field.Type $v.Name "_buf" }}
{{- end }}

            PostInit();
		}

		public int DataCount => _dataList.Count;
		public {{.KeyDefTypeMap}} DataMap => _dataMap;
		public List<{{.RecordClassName}}> DataList => _dataList;
		
{{UniqueMapGetFunc $ "_dataMap" 2}} 

{{UniqueMapGetFuncWithoutError $ "_dataMap" 2}} 

{{ range $v := .CSCodeExtraFields }}
        /// <summary>
        /// {{$v.Desc}}
        /// </summary>
		public {{$v.TypeDef}} {{$v.Name}} { get; private set; }
{{ end }}
        /// <summary>
        /// table data file
        /// </summary>
        public static string TableFileName { get; } = "{{.TableName}}";

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
			{{CSbinFieldReaderEx $v "_buf" "_commonData._field"}}
{{- end }}
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

const templateCSCodeEnum = `
namespace {{.NameSpace}}
{
{{ range $v := .Enums }}
    public enum {{$.EnumDefinePrefix}}{{$v.Name}}
    {
{{- range $enumValue := $v.Values }}
        // {{$enumValue.Name}} : {{$enumValue.Desc}}
        {{$enumValue.Name}} = {{$enumValue.Index}},
{{- end }}
    }
{{ end }}
}
`

const templateCSCodeClass = `{{ .CodeHead }}
using System.Collections.Generic;

namespace {{.NameSpace}} 
{
{{ range $classType := .ClassTypes }}
	public partial class {{$.ClassDefinePrefix}}{{$classType.Name}}
	{
{{- range $v := $classType.Fields }}
		public {{$v.TypeDef}} {{$v.Name}} { get; private set; }
{{- end }}

		public {{$.ClassDefinePrefix}}{{$classType.Name}}(ByteBuf _buf)
        {
{{- range $v := $classType.Fields }}
			{{CSbinFieldReader $v.FieldType $v.Name "_buf" }}
{{- end }}
            PostInit();
        }

        /// <summary>
        /// post process ext class
        /// </summary>
        partial void PostInit();
	}
{{ end }}
}
`
