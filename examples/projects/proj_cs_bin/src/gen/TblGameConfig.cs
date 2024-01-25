using Serialization;
    
using System.Collections.Generic;

namespace CfgTable 
{

	public partial class TblGameConfig
	{

		private readonly Dictionary<int, CfgGameConfig> _dataMap;
        private readonly List<CfgGameConfig> _dataList;
 
		
  		public TblGameConfig(ByteBuf _buf)
        {
			//first read common data
            _TbCommonGameConfig _commonData = null;
            var commonSize = _buf.ReadSize();
            if( commonSize > 0)
            {
                _commonData = new _TbCommonGameConfig(_buf);
            }

			var size = _buf.ReadSize();
            _dataMap = new Dictionary<int, CfgGameConfig>(16);
 			_dataList = new List<CfgGameConfig>(size);

  			for (int i = 0; i < size; i++)
            {
                CfgGameConfig _v;
                _v = CfgGameConfig.DeserializeCfgGameConfig(_buf, _commonData);
                _dataList.Add(_v);
				_dataMap[_v.id] = _v;
            }

			heroLevelMax = _buf.ReadInt();
			{int __n0 = _buf.ReadSize(); heroInitTeam = new int[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ int __v0; __v0 = _buf.ReadInt(); heroInitTeam[__i0] = __v0; } }
			initItems = new KVList_IntInt(_buf);
			openDebug = _buf.ReadBool();

            PostInit();
		}

		public int DataCount => _dataList.Count;
		public Dictionary<int, CfgGameConfig> DataMap => _dataMap;
		public List<CfgGameConfig> DataList => _dataList;
		
        public CfgGameConfig Get(int __k0) 
        {
            if (_dataMap.TryGetValue(__k0, out var __tmpv0)) { return __tmpv0; }
            #if UNITY_EDITOR
            Debug.LogError($"[TblGameConfig] config id not found,id:{__k0.ToString()}");
            #endif
            return null; 
        } 

        public CfgGameConfig GetWithoutError(int __k0) 
        {
            if (_dataMap.TryGetValue(__k0, out var __tmpv0)) { return __tmpv0; }
            return null; 
        } 


        /// <summary>
        /// 
        /// </summary>
		public int heroLevelMax { get; private set; }

        /// <summary>
        /// 
        /// </summary>
		public int[] heroInitTeam { get; private set; }

        /// <summary>
        /// 
        /// </summary>
		public KVList_IntInt initItems { get; private set; }

        /// <summary>
        /// 
        /// </summary>
		public bool openDebug { get; private set; }

        /// <summary>
        /// table data file
        /// </summary>
        public static string TableFileName { get; } = "GameConfig";

        /// <summary>
        /// post process table
        /// </summary>
		partial void PostInit();

	}

	public partial class CfgGameConfig 
	{
		private CfgGameConfig(ByteBuf _buf, _TbCommonGameConfig _commonData)
        {

			id = _buf.ReadInt();
			{ int dataIndex = _buf.ReadInt() - 1; data = _commonData._field0[dataIndex]; }
			{ int dataIndex = _buf.ReadInt() - 1; 备注 = _commonData._field1[dataIndex]; }
            PostInit();
        }

        internal static CfgGameConfig DeserializeCfgGameConfig(ByteBuf _buf, _TbCommonGameConfig _commonData)
        {
            return new CfgGameConfig(_buf, _commonData);
        }

        /// <summary>
        /// id
        /// </summary>
		public int id { get; private set; }

        /// <summary>
        /// 内容值
        /// </summary>
		public string data { get; private set; }

        /// <summary>
        /// 描述
        /// </summary>
		public string 备注 { get; private set; }


        /// <summary>
        /// post process table
        /// </summary>
        partial void PostInit();
	}

    /// <summary>
    /// internal common data optimize
    /// </summary>
    internal class _TbCommonGameConfig
    {

        internal string[] _field0 { get; private set; }
        internal string[] _field1 { get; private set; }
        internal _TbCommonGameConfig(ByteBuf _buf)
        {

			{int __n0 = _buf.ReadSize(); _field0 = new string[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ string __v0; __v0 = _buf.ReadString(); _field0[__i0] = __v0; } }
			{int __n0 = _buf.ReadSize(); _field1 = new string[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ string __v0; __v0 = _buf.ReadString(); _field1[__i0] = __v0; } }
        }

    }

}
