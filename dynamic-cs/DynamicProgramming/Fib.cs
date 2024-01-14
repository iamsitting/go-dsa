public static class Fib
{
    private static int Solve(int n)
    {
        if(n < 2) return 1;

        return Solve(n-1) + Solve(n-2);
    }

    private static int Solve(int n, Dictionary<int, int> memo)
    {
        if(memo.TryGetValue(n, out int value)) return value;
        if(n < 2) return 1;
        return Solve(n-1) + Solve(n-2);
    }

    public static void Test()
    {
        Console.WriteLine(Solve(3));
        Console.WriteLine(Solve(5));
        Console.WriteLine(Solve(25, []));
    }
}