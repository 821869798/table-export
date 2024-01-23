using Serialization;
    
using System.Collections.Generic;

namespace CfgTable 
{

	public partial class PointInt
	{
		public int x { get; private set; }
		public int y { get; private set; }

		public PointInt(ByteBuf _buf)
        {
			x = _buf.ReadInt();
			y = _buf.ReadInt();
            PostInit();
        }

        /// <summary>
        /// post process ext class
        /// </summary>
        partial void PostInit();
	}

	public partial class TestVector3
	{
		public float x { get; private set; }
		public float y { get; private set; }
		public float z { get; private set; }

		public TestVector3(ByteBuf _buf)
        {
			x = _buf.ReadFloat();
			y = _buf.ReadFloat();
			z = _buf.ReadFloat();
            PostInit();
        }

        /// <summary>
        /// post process ext class
        /// </summary>
        partial void PostInit();
	}

}
