namespace DynamicProgramming;
public static class AllConstruct
{
    private static string[][] TabSolve(string target, string[] wordBank)
    {
        var table = new string[target.Length + 1][][];
        table[0] = [[]];
        for (var i = 1; i <= target.Length; i++)
        {
            table[i] = [];
        }
        for (var i = 0; i <= target.Length; i++)
        {
            //Console.WriteLine($"Index: {i} - {table.Dump()}");
            if (table[i].Length > 0)
            {
                foreach (var word in wordBank)
                {
                    if (i + word.Length <= target.Length)
                    {
                        var prefix = target.Substring(i, word.Length);
                        if (prefix == word)
                        {
                            //table[i + word.Length] = [.. table[i], [word]];

                            var waysToBuildWord = table[i].Select(way => way.Concat(new string[] { word }).ToArray()).ToArray();
                            table[i + word.Length] = [.. table[i + word.Length], ..waysToBuildWord];
                        }
                    }
                }
            }
        }
        return table[target.Length];
    }
    /// <summary>
    /// function should return a 2D array containing all the ways the
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
        Console.WriteLine("Tab Solve");
        Console.WriteLine(TabSolve("abcdef", ["ab", "abc", "cd", "def", "abcd"]).Dump());
        Console.WriteLine(TabSolve("abcdef", ["ab", "abc", "cd", "def", "abcd", "ef"]).Dump());
        Console.WriteLine(TabSolve("skateboard", ["bo", "rd", "ate", "t", "boar", "sk"]).Dump());
    }
}