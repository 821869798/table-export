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
        var cfg1 = table.Get(1);
        Console.WriteLine(JsonSerializer.Serialize(cfg1));
        Console.WriteLine(table.DataCount);
        
        bytes = File.ReadAllBytes("table_bytes/base_test.bytes");
        ByteBuf buf2 = new ByteBuf(bytes);
        var table2 = new CfgTable.Tblbase_test(buf2);
        var cfg2 = table2.Get(2);
        Console.WriteLine(JsonSerializer.Serialize(cfg2));
        
        bytes = File.ReadAllBytes("table_bytes/complex_test.bytes");
        ByteBuf buf3 = new ByteBuf(bytes);
        var table3 = new CfgTable.Tblcomplex_test(buf3);
        var dataRecord = table3.Get(1, "k3");
        if (dataRecord != null)
        {
            Console.WriteLine(dataRecord.content);
        }
        
        //var summary = BenchmarkRunner.Run<Program>();

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