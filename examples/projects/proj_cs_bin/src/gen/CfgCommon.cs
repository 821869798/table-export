using Serialization;
    
using System.Collections.Generic;

namespace CfgTable 
{

	public partial class KVList_IntInt
	{
		public int[] keys { get; private set; }
		public int[] values { get; private set; }

		public KVList_IntInt(ByteBuf _buf)
        {
			{int __n0 = _buf.ReadSize(); keys = new int[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ int __v0; __v0 = _buf.ReadInt(); keys[__i0] = __v0; } }
			{int __n0 = _buf.ReadSize(); values = new int[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ int __v0; __v0 = _buf.ReadInt(); values[__i0] = __v0; } }
            PostInit();
        }

        /// <summary>
        /// post process ext class
        /// </summary>
        partial void PostInit();
	}

	public partial class KVList_IntFloat
	{
		public int[] keys { get; private set; }
		public float[] values { get; private set; }

		public KVList_IntFloat(ByteBuf _buf)
        {
			{int __n0 = _buf.ReadSize(); keys = new int[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ int __v0; __v0 = _buf.ReadInt(); keys[__i0] = __v0; } }
			{int __n0 = _buf.ReadSize(); values = new float[__n0]; for(var __i0 = 0 ; __i0 < __n0 ; __i0++ ){ float __v0; __v0 = _buf.ReadFloat(); values[__i0] = __v0; } }
            PostInit();
        }

        /// <summary>
        /// post process ext class
        /// </summary>
        partial void PostInit();
	}

}
