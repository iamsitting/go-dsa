namespace DynamicProgramming;

public static class Fib
{
    private static int Solve(int n)
    {
        if (n < 2) return 1;

        return Solve(n - 1) + Solve(n - 2);
    }

    private static int Solve(int n, Dictionary<int, int> memo)
    {
        if (memo.TryGetValue(n, out int value)) return value;
        if (n < 2) return 1;
        var res = Solve(n - 1, memo) + Solve(n - 2, memo);
        memo[n] = res;
        return res;
    }

    public static void Test()
    {
        Console.WriteLine("Testing Fib");
        Console.WriteLine(Solve(3));
        Console.WriteLine(Solve(5));
        Console.WriteLine(Solve(25, []));
        Console.WriteLine("\n\n");
    }
}