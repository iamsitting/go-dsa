namespace DynamicProgramming;
public static class Extensions
{
    public static string Dump(this int[]? array)
    {
        if (array == null) return "null";
        return $"[{string.Join(",", array)}]";
    }

    public static string Dump(this string[][]? array)
    {
        if (array == null) return "null";

        return $"[{string.Join(",", array.Select(row => $"[{string.Join(",", row)}]") )}]";
    }

}