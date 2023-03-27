using System;
using System.Collections.Generic;

namespace Serialization
{
    public class ByteBufPool
    {
        private readonly Stack<ByteBuf> _bufs = new Stack<ByteBuf>();

        private readonly int _maxCacheNum;

        private readonly Action<ByteBuf> _freeAction;

        public ByteBufPool(int maxCacheNum)
        {
            _maxCacheNum = maxCacheNum;
            _freeAction = Free;
        }

        public ByteBuf Alloc(int? hintSize)
        {
            if (_bufs.TryPop(out var result))
            {
                return result;
            }

            return new ByteBuf(hintSize ?? 64, _freeAction);
        }

        public void Free(ByteBuf buf)
        {
            buf.Clear();
            if (_bufs.Count < _maxCacheNum)
            {
                _bufs.Push(buf);
            }
        }

        public Stack<ByteBuf> GetByteBufsForTest()
        {
            return _bufs;
        }
    }
}