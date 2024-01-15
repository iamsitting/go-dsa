namespace DynamicProgramming;

public static class CanConstruct
{

    private static bool TabSolve(string target, string[] wordBank)
    {
        var table = new bool[target.Length + 1];
        table[0] = true;
        for(var i = 0; i <= target.Length; i++)
        {
            if(table[i])
            {
                foreach(var word in wordBank)
                {
                    if(i + word.Length <= target.Length)
                    {
                        var prefix = target.Substring(i, word.Length);
                        if(prefix == word)
                        {
                            table[i + word.Length] = true;
                        }
                    }                    
                }
            }
        }
        return table[target.Length];
    }
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
        Console.WriteLine("Tab Solve");
        Console.WriteLine(TabSolve("abcdef", ["ab", "abc", "cd", "def", "abcd"]));
        Console.WriteLine(TabSolve("skateboard", ["bo", "rd", "ate", "t", "boar", "sk"]));
        Console.WriteLine(TabSolve("eeeeeeeeeeeeeeeeeeeeeeeeeeee", ["e", "eee", "eeeee"]));
    }
}