public static class Extensions
{
    public static string Dump(this int[]? array)
    {
        if(array == null) return "null";
        return $"[{string.Join(",", array)}]";
    }
}