using BenchmarkDotNet.Attributes;
using Serialization;
using System.Text.Json;


[MemoryDiagnoser]
public class Program
{
    static void Main(string[] args)
    {

        var bytes = File.ReadAllBytes("table_bytes/big_data.bytes");
        ByteBuf buf = new ByteBuf(bytes);
        var table = new CfgTable.Tblbig_data(buf);
        var cfg1 = table.GetDataById(1);
        Console.WriteLine(JsonSerializer.Serialize(cfg1));
        Console.WriteLine(table.DataCount);

        //var summary = BenchmarkRunner.Run<Program>();

        bytes = File.ReadAllBytes("table_bytes/complex_test.bytes");
        ByteBuf buf2 = new ByteBuf(bytes);
        var table2 = new CfgTable.Tblcomplex_test(buf2);
        var dataRecord = table2.GetDataById(1, "k3");
        if (table2 != null)
        {
            Console.WriteLine(dataRecord.content);
        }

    }


    byte[] bytes;
    [GlobalSetup]
    public void Setup()
    {
        bytes = File.ReadAllBytes("D:\\program\\go\\table-export\\examples\\projects\\proj_cs_bin\\table_bytes\\big_data.bytes");
    }

    [Benchmark]
    public void TestTableLoad()
    {
        ByteBuf buf = new ByteBuf(bytes);
        var table = new CfgTable.Tblbig_data(buf);
        int count = table.DataCount;
    }

}