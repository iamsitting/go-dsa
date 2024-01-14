namespace DynamicProgramming;

public static class CanConstruct
{
    private static bool Solve(string target, string[] wordBank)
    {
        if (target == "") return true;

        foreach (var word in wordBank)
        {
            if (target.StartsWith(word))
            {
                var suffix = target[word.Length..];
                if (Solve(suffix, wordBank))
                {
                    return true;
                }
            }
        }
        return false;
    }
    private static bool Solve(string target, string[] wordBank, Dictionary<string, bool> memo)
    {
        if(memo.TryGetValue(target, out var value)) return value;
        if (target == "") return true;

        foreach (var word in wordBank)
        {
            if (target.StartsWith(word))
            {
                var suffix = target[word.Length..];
                if (Solve(suffix, wordBank, memo))
                {
                    memo[target] = true;
                    return true;
                }
            }
        }
        memo[target] = false;
        return false;
    }

    public static void Test()
    {
        Console.WriteLine(Solve("abcdef", ["ab", "abc", "cd", "def", "abcd"]));
        Console.WriteLine(Solve("skateboard", ["bo", "rd", "ate", "t", "boar", "sk"]));
        Console.WriteLine(Solve("eeeeeeeeeeeeeeeeeeeeeeeeeeee", ["e", "eee", "eeeee"], []));
    }
}