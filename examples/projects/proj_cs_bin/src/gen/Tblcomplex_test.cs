using Serialization;
        
using System.Collections.Generic;

namespace CfgTable 
{

	public partial class Tblcomplex_test
	{

		private readonly Dictionary<int, Dictionary<string, Cfgcomplex_test>> _dataMap;
        private readonly List<Cfgcomplex_test> _dataList;
 
		
  		public Tblcomplex_test(ByteBuf _buf)
        {
			//first read common data
            _TbCommoncomplex_test _commonData = null;
            var commonSize = _buf.ReadSize();
            if( commonSize > 0)
            {
                _commonData = new _TbCommoncomplex_test(_buf);
            }

			var size = _buf.ReadSize();
            _dataMap = new Dictionary<int, Dictionary<string, Cfgcomplex_test>>(16);
 			_dataList = new List<Cfgcomplex_test>(size);

  			for (int i = 0; i < size; i++)
            {
                Cfgcomplex_test _v;
                _v = Cfgcomplex_test.DeserializeCfgcomplex_test(_buf, _commonData);
                _dataList.Add(_v);
				if (!_dataMap.TryGetValue(_v.key1, out var _tmpuk0))
				{
					_tmpuk0 = new Dictionary<string, Cfgcomplex_test>();
					_dataMap[_v.key1] = _tmpuk0;
				}
				_tmpuk0[_v.key2] = _v;
            }

            PostInit();
		}

		public int DataCount => _dataList.Count;
		public Dictionary<int, Dictionary<string, Cfgcomplex_test>> DataMap => _dataMap;
		public List<Cfgcomplex_test> DataList => _dataList;
		
		public Cfgcomplex_test GetDataById(int __k0, string __k1) { if (_dataMap.TryGetValue(__k0, out var __tmpv0) && __tmpv0.TryGetValue(__k1, out var __tmpv1)) { return __tmpv1; } return null; } 
		
        /// <summary>
        /// post process table
        /// </summary>
		partial void PostInit();

	}

	public partial class Cfgcomplex_test 
	{
		private Cfgcomplex_test(ByteBuf _buf, _TbCommoncomplex_test _commonData)
        {

			key1 = _buf.ReadInt();
			key2 = _buf.ReadString();
			content = _buf.ReadString();
			number = _buf.ReadFloat();
			{ int dataIndex = _buf.ReadInt() - 1; test_list = _commonData._field0[dataIndex]; }
			{ int dataIndex = _buf.ReadInt() - 1; test_map = _commonData._field1[dataIndex]; }
            PostInit();
        }

        internal static Cfgcomplex_test DeserializeCfgcomplex_test(ByteBuf _buf, _TbCommoncomplex_test _commonData)
        {
            return new Cfgcomplex_test(_buf, _commonData);
        }

        /// <summary>
        /// id
        /// </summary>
		public int key1 { get; private set; }

        /// <summary>
        /// id2
        /// </summary>
		public string key2 { get; private set; }

        /// <summary>
        /// 内容
        /// </summary>
		public string content { get; private set; }

        /// <summary>
        /// 数字
        /// </summary>
		public float number { get; private set; }

        /// <summary>
        /// 列表
        /// </summary>
		public int[] test_list { get; private set; }

        /// <summary>
        /// 字典
        /// </summary>
		public Dictionary<int, string> test_map { get; private set; }


        /// <summary>
        /// post process table
        /// </summary>
        partial void PostInit();
	}

    /// <summary>
    /// internal common data.Optimize memory
    /// </summary>
    internal class _TbCommoncomplex_test
    {

        internal int[][] _field0 { get; private set; }
        internal Dictionary<int, string>[] _field1 { get; private set; }
        internal _TbCommoncomplex_test(ByteBuf _buf)
        {

			{int __n0 = _buf.ReadSize(); _field0 = new int[__n0][]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ int[] __v0; {int __n1 = _buf.ReadSize(); __v0 = new int[__n1]; for(var __i1 = 0 ; __i1 < __n1 ; __i1++ ){ int __v1; __v1 = _buf.ReadInt(); __v0[__i1] = __v1; } } _field0[__i0] = __v0; } }
			{int __n0 = _buf.ReadSize(); _field1 = new Dictionary<int, string>[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ Dictionary<int, string> __v0; { int __n1 = _buf.ReadSize(); __v0 = new Dictionary<int, string> (__n1 * 3 / 2); for(var __i1 = 0 ; __i1 < __n1 ; __i1++ ) {int __k1; __k1 = _buf.ReadInt(); string __v1; __v1 = _buf.ReadString(); __v0.Add(__k1, __v1); } } _field1[__i0] = __v0; } }
        }

    }

}
