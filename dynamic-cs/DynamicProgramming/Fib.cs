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

    private static int TabSolve(int n)
    {
        var table = new int[n+1];
        table[1] = 1;
        for(var i = 0; i <= n; i++)
        {
            if (i + 1 <= n) table[i + 1] += table[i];
            if (i + 2 <= n) table[i + 2] += table[i];
        }
        return table[n];
    }

    public static void Test()
    {
        Console.WriteLine("Testing Fib");
        Console.WriteLine(Solve(3));
        Console.WriteLine(Solve(5));
        Console.WriteLine(Solve(25, []));
        Console.WriteLine("\n\n");
        Console.WriteLine(TabSolve(6));
        Console.WriteLine(TabSolve(7));
        Console.WriteLine(TabSolve(8));
        Console.WriteLine(TabSolve(26));
    }
}