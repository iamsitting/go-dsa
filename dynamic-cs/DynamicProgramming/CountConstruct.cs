namespace DynamicProgramming;

public static class CountConstruct
{
    private static int Solve(string target, string[] wordBank)
    {
        if (target == "") return 1;

        var count = 0;
        foreach(var word in wordBank)
        {
            if(target.StartsWith(word))
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
        foreach(var word in wordBank)
        {
            if(target.StartsWith(word))
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
        Console.WriteLine(Solve("skateboard", ["bo", "rd", "ate", "t", "boar", "sk"]));
        Console.WriteLine(Solve("eeeeeeeeeeeeeeeeeeeeeeeeeeee", ["e", "eee", "eeeee"], []));
    }
}