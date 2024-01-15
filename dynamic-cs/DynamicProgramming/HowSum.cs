namespace DynamicProgramming;

public static class HowSum
{
    private static int[]? TabSolve(int target, int[] numbers)
    {
        int[][]? table = new int[target+1][];
        table[0] = [];
        for(var i = 0; i <= target; i++)
        {
            if(table[i] != null)
            {
                foreach(var number in numbers)
                {
                    if (number + i <= target)
                    {
                        table[number + i] = [..table[i], number];
                    }
                }
            }
        }
        return table[target];
    }
    private static int[]? Solve(int target, int[] numbers)
    {
        if (target == 0) return [];
        if (target < 0) return null;

        foreach (var number in numbers)
        {
            var newArray = Solve(target - number, numbers);
            if (newArray != null)
            {
                return [..newArray, number];

            }
        }

        return null;
    }

    private static int[]? Solve(int target, int[] numbers, Dictionary<int, int[]?> memo)
    {
        if(memo.TryGetValue(target, out var value)) return value;
        if (target == 0) return [];
        if (target < 0) return null;

        foreach (var number in numbers)
        {
            var newArray = Solve(target - number, numbers, memo);
            if (newArray != null)
            {
                memo[target] = [..newArray, number];
                return memo[target];

            }
        }
        
        memo[target] = null;
        return null;
    }


    public static void Test()
    {
        Console.WriteLine(Solve(7, [2, 3, 4, 8]).Dump());
        Console.WriteLine(Solve(7, [2, 4]).Dump());
        Console.WriteLine(Solve(13, [2, 4, 25, 3]).Dump());
        Console.WriteLine(Solve(300, [7, 14], []).Dump());
        Console.WriteLine("Tab Solve");
        Console.WriteLine(TabSolve(7, [2, 3, 4, 8]).Dump());
        Console.WriteLine(TabSolve(7, [2, 4]).Dump());
        Console.WriteLine(TabSolve(13, [2, 4, 25, 3]).Dump());
        Console.WriteLine(TabSolve(300, [7, 14]).Dump());
    }
}