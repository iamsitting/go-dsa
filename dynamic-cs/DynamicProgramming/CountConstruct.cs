namespace DynamicProgramming;

public static class CountConstruct
{
    private static int TabSolve(string target, string[] wordBank)
    {
        var table = new int[target.Length + 1];
        table[0] = 1; // there is 1 way to build an empty string
        for (var i = 0; i <= target.Length; i++)
        {
            if (table[i] > 0)
            {
                foreach(var word in wordBank)
                {
                    if(i + word.Length <= target.Length)
                    {
                        var prefix = target.Substring(i, word.Length);
                        if(prefix == word)
                        {
                            table[i + word.Length] += table[i];
                        }
                    }
                }
            }
        }
        return table[target.Length];
    }
    private static int Solve(string target, string[] wordBank)
    {
        if (target == "") return 1;

        var count = 0;
        foreach (var word in wordBank)
        {
            if (target.StartsWith(word))
            {
                var suffix = target[word.Length..];
                count += Solve(suffix, wordBank);
            }
        }
        return count;
    }

    private static int Solve(string target, string[] wordBank, Dictionary<string, int> memo)
    {
        if (memo.TryGetValue(target, out var value)) return value;

        if (target == "") return 1;

        var count = 0;
        foreach (var word in wordBank)
        {
            if (target.StartsWith(word))
            {
                var suffix = target[word.Length..];
                count += Solve(suffix, wordBank, memo);
            }
        }
        memo[target] = count;
        return count;
    }


    public static void Test()
    {
        Console.WriteLine(Solve("abcdef", ["ab", "abc", "cd", "def", "abcd"]));
        Console.WriteLine(Solve("abcdef", ["ab", "abc", "cd", "def", "abcd", "ef"]));
        Console.WriteLine(Solve("skateboard", ["bo", "rd", "ate", "t", "boar", "sk"]));
        Console.WriteLine(Solve("eeeeeeeeeeeeeeeeeeeeeeeeeeee", ["e", "eee", "eeeee"], []));
        Console.WriteLine("Tab Solve");
        Console.WriteLine(TabSolve("abcdef", ["ab", "abc", "cd", "def", "abcd"]));
        Console.WriteLine(TabSolve("abcdef", ["ab", "abc", "cd", "def", "abcd", "ef"]));
        Console.WriteLine(TabSolve("skateboard", ["bo", "rd", "ate", "t", "boar", "sk"]));
        Console.WriteLine(TabSolve("eeeeeeeeeeeeeeeeeeeeeeeeeeee", ["e", "eee", "eeeee"]));
    }
}