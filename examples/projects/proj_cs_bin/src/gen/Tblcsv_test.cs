using Serialization;
    
using System.Collections.Generic;

namespace CfgTable 
{

	public partial class Tblcsv_test
	{

		private readonly Dictionary<int, Cfgcsv_test> _dataMap;
        private readonly List<Cfgcsv_test> _dataList;
 
		
  		public Tblcsv_test(ByteBuf _buf)
        {
			//first read common data
            _TbCommoncsv_test _commonData = null;
            var commonSize = _buf.ReadSize();
            if( commonSize > 0)
            {
                _commonData = new _TbCommoncsv_test(_buf);
            }

			var size = _buf.ReadSize();
            _dataMap = new Dictionary<int, Cfgcsv_test>(16);
 			_dataList = new List<Cfgcsv_test>(size);

  			for (int i = 0; i < size; i++)
            {
                Cfgcsv_test _v;
                _v = Cfgcsv_test.DeserializeCfgcsv_test(_buf, _commonData);
                _dataList.Add(_v);
				_dataMap[_v.id] = _v;
            }

            PostInit();
		}

		public int DataCount => _dataList.Count;
		public Dictionary<int, Cfgcsv_test> DataMap => _dataMap;
		public List<Cfgcsv_test> DataList => _dataList;
		
		public Cfgcsv_test GetDataById(int __k0) { if (_dataMap.TryGetValue(__k0, out var __tmpv0)) { return __tmpv0; } return null; } 
		
        /// <summary>
        /// post process table
        /// </summary>
		partial void PostInit();

	}

	public partial class Cfgcsv_test 
	{
		private Cfgcsv_test(ByteBuf _buf, _TbCommoncsv_test _commonData)
        {

			id = _buf.ReadInt();
			{ int dataIndex = _buf.ReadInt() - 1; name = _commonData._field0[dataIndex]; }
			age = _buf.ReadInt();
            PostInit();
        }

        internal static Cfgcsv_test DeserializeCfgcsv_test(ByteBuf _buf, _TbCommoncsv_test _commonData)
        {
            return new Cfgcsv_test(_buf, _commonData);
        }

        /// <summary>
        /// id
        /// </summary>
		public int id { get; private set; }

        /// <summary>
        /// 名字
        /// </summary>
		public string name { get; private set; }

        /// <summary>
        /// 年龄
        /// </summary>
		public int age { get; private set; }


        /// <summary>
        /// post process table
        /// </summary>
        partial void PostInit();
	}

    /// <summary>
    /// internal common data optimize
    /// </summary>
    internal class _TbCommoncsv_test
    {

        internal string[] _field0 { get; private set; }
        internal _TbCommoncsv_test(ByteBuf _buf)
        {

			{int __n0 = _buf.ReadSize(); _field0 = new string[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ string __v0; __v0 = _buf.ReadString(); _field0[__i0] = __v0; } }
        }

    }

}
