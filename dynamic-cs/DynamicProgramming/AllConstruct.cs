namespace DynamicProgramming;
public static class AllConstruct
{
    /// <summary>
    /// funciton should return a 2D array containing all the ways the
    /// target can be constructed from the wordBank
    /// </summary>
    private static string[][] Solve(string target, string[] wordBank)
    {
        if (string.IsNullOrEmpty(target))
        {
            return [[]];
        }

        string[][] result = [];
        foreach (var word in wordBank)
        {
            if (target.StartsWith(word))
            {
                string suffix = target[word.Length..];
                string[][] waysToBuildSuffix = Solve(suffix, wordBank);
                string[][] waysToBuildTarget = waysToBuildSuffix.Select(way => new string[] { word }.Concat(way).ToArray()).ToArray();
                result = [.. result, .. waysToBuildTarget];
            }
        }

        return result;
    }

    private static string[][] Solve(string target, string[] wordBank, Dictionary<string, string[][]> memo)
    {
        if (string.IsNullOrEmpty(target))
        {
            return [[]];
        }

        if (memo.TryGetValue(target, out var value)) return value;

        string[][] result = [];
        foreach (var word in wordBank)
        {
            if (target.StartsWith(word))
            {
                string suffix = target[word.Length..];
                string[][] waysToBuildSuffix = Solve(suffix, wordBank, memo);
                string[][] waysToBuildTarget = waysToBuildSuffix.Select(way => new string[] { word }.Concat(way).ToArray()).ToArray();
                result = [.. result, .. waysToBuildTarget];
            }
        }

        memo[target] = result;
        return result;
    }

    public static void Test()
    {
        Console.WriteLine(Solve("abcdef", ["ab", "abc", "cd", "def", "abcd"]).Dump());
        Console.WriteLine(Solve("abcdef", ["ab", "abc", "cd", "def", "abcd", "ef"], []).Dump());
        Console.WriteLine(Solve("skateboard", ["bo", "rd", "ate", "t", "boar", "sk"], []).Dump());
    }
}