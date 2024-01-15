using System.ComponentModel.DataAnnotations.Schema;

namespace DynamicProgramming;

public static class CanSum
{
    // 7 -> [4, 2, 3, 6]
    //init:  [t, f, f, f, f, f, f, f]
    //0:     [t, f, t, t, t, f, t, f]
    //1:     [t, f, t, t, t, t, t, f]


    private static bool TabSolve(int target, int[] numbers)
    {
        var table = new bool[target + 1];
        table[0] = true;
        for (var i = 0; i <= target; i++)
        {
            if (table[i])
            {
                foreach (var number in numbers)
                {
                    if (number + i <= target)
                    {
                        table[number + i] = true;
                    }
                }
            }
        }
        return table[target];
    }
    private static bool Solve(int target, int[] numbers)
    {
        if (target == 0) return true;
        if (target < 0) return false;

        foreach (var number in numbers)
        {
            var res = Solve(target - number, numbers);
            if (res) return true;
        }

        return false;
    }

    private static bool Solve(int target, int[] numbers, Dictionary<int, bool> memo)
    {
        if (memo.TryGetValue(target, out var value)) return value;
        if (target == 0) return true;
        if (target < 0) return false;

        foreach (var number in numbers)
        {
            var res = Solve(target - number, numbers, memo);
            if (res)
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
        Console.WriteLine("Tab Solve");
        Console.WriteLine(TabSolve(7, [2, 3, 4, 8]));
        Console.WriteLine(TabSolve(7, [2, 4]));
        Console.WriteLine(TabSolve(13, [2, 4, 25, 3]));
        Console.WriteLine(TabSolve(300, [7, 14]));
    }
}