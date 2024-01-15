namespace DynamicProgramming;

public static class GridTraveler
{

    private static int TabSolve(int x, int y)
    {
        var table = new int[x+1, y+1];
        table[1,1] = 1;
        for(var i = 0; i <= x; i++)
        {
            for(var j = 0; j <= y; j++)
            {
                if (j+1 <= y) table[i, j+1] += table[i, j];
                if (i+1 <= x) table[i+1, j] += table[i, j];
            }
        }

        return table[x, y];

    }
    private static int Solve(int x, int y)
    {
        if (x == 0 || y == 0)
        {
            return 0;
        }

        if (x == 1 && y == 1)
        {
            return 1;
        }

        return Solve(x - 1, y) + Solve(x, y - 1);
    }

    private static int Solve(int x, int y, Dictionary<string, int> memo)
    {
        var key = $"{x}+{y}";
        if(memo.TryGetValue(key, out int value)) return value;
        if (x == 0 || y == 0)
        {
            return 0;
        }

        if (x == 1 && y == 1)
        {
            return 1;
        }
        var res = Solve(x - 1, y, memo) + Solve(x, y - 1, memo);
        memo[key] = res;
        return res;
    }

    public static void Test()
    {
        Console.WriteLine("Testing GridTraveler");
        Console.WriteLine(Solve(3, 3));
        Console.WriteLine(Solve(10, 4, []));
        Console.WriteLine("\n\n");
        Console.WriteLine(TabSolve(3, 3));
        Console.WriteLine(TabSolve(10, 4));
    }
}