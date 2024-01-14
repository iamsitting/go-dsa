namespace DynamicProgramming;

public static class CanSum
{
    private static bool Solve(int target, int[] numbers)
    {
        if(target == 0) return true;
        if(target < 0) return false;

        foreach(var number in numbers)
        {
            var res = Solve(target - number, numbers);
            if(res) return true;
        }

        return false;
    }

    private static bool Solve(int target, int[] numbers, Dictionary<int, bool> memo)
    {
        if (memo.TryGetValue(target, out var value)) return value;
        if(target == 0) return true;
        if(target < 0) return false;

        foreach(var number in numbers)
        {
            var res = Solve(target - number, numbers, memo);
            if(res)
            {
                memo[target] = res;
                return res;
            }
        }

        memo[target] = false;
        return false;
    }
    
    public static void Test()
    {
        Console.WriteLine(Solve(7, [2, 3, 4, 8]));
        Console.WriteLine(Solve(7, [2, 4]));
        Console.WriteLine(Solve(13, [2, 4, 25, 3]));
        Console.WriteLine(Solve(300, [7, 14], []));
    }
}