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

}
