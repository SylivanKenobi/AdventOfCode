file = File.open("input.txt")
lines = file.read.split
num_map = {
    "1" => "one",
    "2" => "two",
    "3" => "three",
    "4" => "four",
    "5" => "five",
    "6" => "six",
    "7" => "seven",
    "8" => "eight",
    "9" => "nine"
}
claibration = 0
lines.each do |line|
    clean_line = line.scan(/([0-9]|one|two|three|four|five|six|seven|eight|nine)/).join
    num_map.each do |x, y|
        clean_line.gsub! y.to_s, x.to_s
    end
    coordinate = "#{clean_line[0]}#{clean_line[-1]}"
    claibration += coordinate.to_i
end
pp claibration
