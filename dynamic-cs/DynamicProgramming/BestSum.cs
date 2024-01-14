namespace DynamicProgramming;

public static class BestSum
{
    private static int[]? Solve(int target, int[] numbers)
    {
        if (target == 0) return [];
        if (target < 0) return null;

        int[]? bestSum = null;
        foreach (var number in numbers)
        {
            var newArray = Solve(target - number, numbers);
            if (newArray != null)
            {
                int[] aSum = [..newArray, number];
                if(bestSum == null || aSum.Length < bestSum.Length)
                {
                    bestSum = aSum;
                }
            }
        }

        return bestSum;
    }

    private static int[]? Solve(int target, int[] numbers, Dictionary<int, int[]?> memo)
    {
        if(memo.TryGetValue(target, out var value)) return value;
        if (target == 0) return [];
        if (target < 0) return null;

        int[]? bestSum = null;
        foreach (var number in numbers)
        {
            var newArray = Solve(target - number, numbers, memo);
            if (newArray != null)
            {
                int[] aSum = [..newArray, number];
                if(bestSum == null || aSum.Length < bestSum.Length)
                {
                    bestSum = aSum;
                }
            }
        }
        memo[target] = bestSum;
        return bestSum;
    }

    public static void Test()
    {
        Console.WriteLine(Solve(7, [2, 3, 4, 8]).Dump());
        Console.WriteLine(Solve(7, [2, 4]).Dump());
        Console.WriteLine(Solve(13, [2, 4, 25, 3]).Dump());
        Console.WriteLine(Solve(300, [7, 14], []).Dump());
    }
}