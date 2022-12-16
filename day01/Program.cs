int sumCalories(string elf) =>
    elf.Split("\n").Where(l => l.Length > 0).Select(line =>
    {
        return int.Parse(line);
    }).Sum();

string input;
using (var reader = new StreamReader("input.txt"))
{
    input = reader.ReadToEnd();
}

IEnumerable<string> elves = input.Split("\n\n").Where(str => str.Length > 0);

var calories = elves.Select(sumCalories);
var maxCalories = calories.Max();
var topCalories = calories.OrderDescending().Take(3).Sum();

Console.WriteLine($"Max calories: {maxCalories}");
Console.WriteLine($"Top 3 calories: {topCalories}");
