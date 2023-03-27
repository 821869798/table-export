using Serialization;
using System.Diagnostics;
using System.Text.Json;

class Program
{
    static void Main(string[] args)
    {

        var bytes = File.ReadAllBytes("table_bytes/base_test.bytes");
        ByteBuf buf = new ByteBuf(bytes);
        var table = new CfgTable.Tblbase_test(buf);
        Console.WriteLine(table.DataCount);

        bytes = File.ReadAllBytes("table_bytes/complex_test.bytes");
        ByteBuf buf2 = new ByteBuf(bytes);
        var table2 = new CfgTable.Tblcomplex_test(buf2);
        var dataRecord = table2.GetDataById(1, "k3");
        if (table2 != null)
        {
            Console.WriteLine(dataRecord.content);
        }

    }

}