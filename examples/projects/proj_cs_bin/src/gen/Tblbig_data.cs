using Serialization;
        
using System.Collections.Generic;

namespace CfgTable 
{

	public partial class Tblbig_data
	{

		private readonly Dictionary<int, Cfgbig_data> _dataMap;
        private readonly List<Cfgbig_data> _dataList;
 
		
  		public Tblbig_data(ByteBuf _buf)
        {
			//first read common data
            _TbCommonbig_data _commonData = null;
            var commonSize = _buf.ReadSize();
            if( commonSize > 0)
            {
                _commonData = new _TbCommonbig_data(_buf);
            }

			var size = _buf.ReadSize();
            _dataMap = new Dictionary<int, Cfgbig_data>(16);
 			_dataList = new List<Cfgbig_data>(size);

  			for (int i = 0; i < size; i++)
            {
                Cfgbig_data _v;
                _v = Cfgbig_data.DeserializeCfgbig_data(_buf, _commonData);
                _dataList.Add(_v);
				_dataMap[_v.id] = _v;
            }

            PostInit();
		}

		public int DataCount => _dataList.Count;
		public Dictionary<int, Cfgbig_data> DataMap => _dataMap;
		public List<Cfgbig_data> DataList => _dataList;
		
		public Cfgbig_data GetDataById(int __k0) { if (_dataMap.TryGetValue(__k0, out var __tmpv0)) { return __tmpv0; } return null; } 
		
        /// <summary>
        /// post process table
        /// </summary>
		partial void PostInit();

	}

	public partial class Cfgbig_data 
	{
		private Cfgbig_data(ByteBuf _buf, _TbCommonbig_data _commonData)
        {

			id = _buf.ReadInt();
			{ int dataIndex = _buf.ReadInt() - 1; name = _commonData._field0[dataIndex]; }
			age = _buf.ReadInt();
			{ int dataIndex = _buf.ReadInt() - 1; course = _commonData._field1[dataIndex]; }
			{ int dataIndex = _buf.ReadInt() - 1; score = _commonData._field2[dataIndex]; }
            PostInit();
        }

        internal static Cfgbig_data DeserializeCfgbig_data(ByteBuf _buf, _TbCommonbig_data _commonData)
        {
            return new Cfgbig_data(_buf, _commonData);
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
        /// 学科id
        /// </summary>
		public int[] course { get; private set; }

        /// <summary>
        /// 成绩组
        /// </summary>
		public Dictionary<int, int> score { get; private set; }


        /// <summary>
        /// post process table
        /// </summary>
        partial void PostInit();
	}

    /// <summary>
    /// internal common data.Optimize memory
    /// </summary>
    internal class _TbCommonbig_data
    {

        internal string[] _field0 { get; private set; }
        internal int[][] _field1 { get; private set; }
        internal Dictionary<int, int>[] _field2 { get; private set; }
        internal _TbCommonbig_data(ByteBuf _buf)
        {

			{int __n0 = _buf.ReadSize(); _field0 = new string[__n0];for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){string __v0; __v0 = _buf.ReadString(); _field0[__i0] = __v0;} }
			{int __n0 = _buf.ReadSize(); _field1 = new int[__n0][];for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){int[] __v0; {int __n1 = _buf.ReadSize(); __v0 = new int[__n1];for(var __i1 = 0 ; __i1 < __n1 ; __i1++ ){int __v1; __v1 = _buf.ReadInt(); __v0[__i1] = __v1;} } _field1[__i0] = __v0;} }
			{int __n0 = _buf.ReadSize(); _field2 = new Dictionary<int, int>[__n0];for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){Dictionary<int, int> __v0; { int __n1 = _buf.ReadSize(); var __f1 = new Dictionary<int, int> (__n1 * 3 / 2) ; __v0 = __f1; for(var __i1 = 0 ; __i1 < __n1 ; __i1++ ) {int __k1; __k1 = _buf.ReadInt(); int __v1; __v1 = _buf.ReadInt(); __f1.Add(__k1, __v1); } } _field2[__i0] = __v0;} }
        }

    }

}
