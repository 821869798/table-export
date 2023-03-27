
namespace Serialization
{


    public class ThreadLocalTemporalByteBufPool
    {
        [ThreadStatic]
        private static ByteBufPool s_pool;

        [ThreadStatic]
        private static ByteBuf s_fastBuf;

        public static int MaxPoolCacheNum { get; set; } = 100;


        private static ByteBufPool Pool => s_pool ?? (s_pool = new ByteBufPool(MaxPoolCacheNum));

        public static TemporalByteBuf GuardAlloc(int? hintSize)
        {
            return new TemporalByteBuf(Alloc(hintSize));
        }

        public static ByteBuf Alloc(int? hintSize)
        {
            ByteBuf byteBuf = s_fastBuf;
            if (byteBuf != null)
            {
                s_fastBuf = null;
                return byteBuf;
            }

            return Pool.Alloc(hintSize);
        }

        public static void Free(ByteBuf buf)
        {
            buf.Clear();
            if (s_fastBuf == null)
            {
                s_fastBuf = buf;
            }
            else
            {
                Pool.Free(buf);
            }
        }

        public static byte[] CopyDataThenFree(ByteBuf buf)
        {
            byte[] result = buf.CopyData();
            Free(buf);
            return result;
        }

        internal static void ResetForTest()
        {
            s_fastBuf = null;
            s_pool = null;
        }

        internal static ByteBuf GetFastBufForTest()
        {
            return s_fastBuf;
        }

        internal static Stack<ByteBuf> GetByteBufsForTest()
        {
            return Pool.GetByteBufsForTest();
        }
    }


    public readonly struct TemporalByteBuf : IDisposable
    {
        public ByteBuf Buf { get; }

        public TemporalByteBuf(ByteBuf buf)
        {
            Buf = buf;
        }

        public void Dispose()
        {
            ThreadLocalTemporalByteBufPool.Free(Buf);
        }
    }

}
