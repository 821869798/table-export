
namespace CfgTable
{

    public enum ItemType
    {
        // None : 
        None = 0,
        // Resource : 资源类
        Resource = 1,
        // Package : 礼包
        Package = 2,
    }

    public enum ConditionType
    {
        // None : 
        None = 0,
        // Equal : 等于
        Equal = 1,
        // Greater : 大于
        Greater = 2,
        // Less : 小于
        Less = 3,
        // GreaterEqual : 大于等于
        GreaterEqual = 4,
        // LessEqual : 大于等于
        LessEqual = 5,
        // NotEqual : 不等于
        NotEqual = 10,
    }

}
