require 'pry'

file = File.open("input.txt")
lines = file.read.split("\n")
result = 0
lines.each_with_index do |line, i|
    matches = []
    line.scan(/([0-9]{1,3})/).flatten().each do |number|
        line.scan(/(?:^|(?<=[^\d]))#{number}(?=[^\d]|$)/) do |match|
            matches.append({num: number, index: Regexp.last_match.begin(0)})
        end
    end
    matches = matches.uniq()
    matches.each do |match|
        number = match[:num]
        num_index = match[:index]
        proximity = false
        # binding.pry
        start_index = (num_index - 1) >= 0 ? (num_index - 1) : 0
        end_index = start_index + number.length + 1
        ((start_index)..(end_index)).to_a.each do |index|
            if index >= 0 && index < line.length && line[index].match?(/[^A-Za-z0-9\.]/)
                proximity = true
            end
            if i > 0 && index < lines[i-1].length && lines[i-1][index].match?(/[^A-Za-z0-9\.]/)
                proximity = true
            end
            if i < lines.length() - 1 && index < lines[i+1].length && lines[i+1][index].match?(/[^A-Za-z0-9\.]/)
                proximity = true
            end
        end
        if proximity
            result += number.to_i
        end
    end
end
pp result