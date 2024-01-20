using Serialization;
    
using System.Collections.Generic;

namespace CfgTable 
{

	public partial class Tblbase_test
	{

		private readonly Dictionary<int, Cfgbase_test> _dataMap;
        private readonly List<Cfgbase_test> _dataList;
 
		
  		public Tblbase_test(ByteBuf _buf)
        {
			//first read common data
            _TbCommonbase_test _commonData = null;
            var commonSize = _buf.ReadSize();
            if( commonSize > 0)
            {
                _commonData = new _TbCommonbase_test(_buf);
            }

			var size = _buf.ReadSize();
            _dataMap = new Dictionary<int, Cfgbase_test>(16);
 			_dataList = new List<Cfgbase_test>(size);

  			for (int i = 0; i < size; i++)
            {
                Cfgbase_test _v;
                _v = Cfgbase_test.DeserializeCfgbase_test(_buf, _commonData);
                _dataList.Add(_v);
				_dataMap[_v.id] = _v;
            }

            PostInit();
		}

		public int DataCount => _dataList.Count;
		public Dictionary<int, Cfgbase_test> DataMap => _dataMap;
		public List<Cfgbase_test> DataList => _dataList;
		
        public Cfgbase_test Get(int __k0) 
        {
            if (_dataMap.TryGetValue(__k0, out var __tmpv0)) { return __tmpv0; }
            #if UNITY_EDITOR
            Debug.LogError($"[Tblbase_test] config id not found,id:{__k0.ToString()}");
            #endif
            return null; 
        } 

        public Cfgbase_test GetWithoutError(int __k0) 
        {
            if (_dataMap.TryGetValue(__k0, out var __tmpv0)) { return __tmpv0; }
            return null; 
        } 
		
        /// <summary>
        /// post process table
        /// </summary>
		partial void PostInit();

	}

	public partial class Cfgbase_test 
	{
		private Cfgbase_test(ByteBuf _buf, _TbCommonbase_test _commonData)
        {

			id = _buf.ReadInt();
			{ int dataIndex = _buf.ReadInt() - 1; name = _commonData._field0[dataIndex]; }
			age = _buf.ReadInt();
			{ int dataIndex = _buf.ReadInt() - 1; course = _commonData._field1[dataIndex]; }
			bigDataId = _buf.ReadInt();
			itemType = (ItemType)_buf.ReadInt();
			conditionType = (ConditionType)_buf.ReadInt();
            PostInit();
        }

        internal static Cfgbase_test DeserializeCfgbase_test(ByteBuf _buf, _TbCommonbase_test _commonData)
        {
            return new Cfgbase_test(_buf, _commonData);
        }

        /// <summary>
        /// 主id
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
        /// 学科
        /// </summary>
		public int[] course { get; private set; }

        /// <summary>
        /// 对应big_data表id
        /// </summary>
		public int bigDataId { get; private set; }

        /// <summary>
        /// 道具类型，测试枚举1
        /// </summary>
		public ItemType itemType { get; private set; }

        /// <summary>
        /// 条件类型，测试枚举2
        /// </summary>
		public ConditionType conditionType { get; private set; }


        /// <summary>
        /// post process table
        /// </summary>
        partial void PostInit();
	}

    /// <summary>
    /// internal common data optimize
    /// </summary>
    internal class _TbCommonbase_test
    {

        internal string[] _field0 { get; private set; }
        internal int[][] _field1 { get; private set; }
        internal _TbCommonbase_test(ByteBuf _buf)
        {

			{int __n0 = _buf.ReadSize(); _field0 = new string[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ string __v0; __v0 = _buf.ReadString(); _field0[__i0] = __v0; } }
			{int __n0 = _buf.ReadSize(); _field1 = new int[__n0][]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ int[] __v0; {int __n1 = _buf.ReadSize(); __v0 = new int[__n1]; for(var __i1 = 0 ; __i1 < __n1 ; __i1++ ){ int __v1; __v1 = _buf.ReadInt(); __v0[__i1] = __v1; } } _field1[__i0] = __v0; } }
        }

    }

}
