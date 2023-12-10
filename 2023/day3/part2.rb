require 'pry'

file = File.open("input.txt")
lines = file.read.split("\n")
result = 0
lines.each_with_index do |line, i|
    matches = []
    line.scan(/(\*)/) do |match|
        matches.append(Regexp.last_match.begin(0))
    end
    matches.each do |match|
        num_index = match
        proximity = 0
        ratio = 0
        start_index = (num_index - 1) >= 0 ? (num_index - 1) : 0
        end_index = start_index + 2

        numbers = []
        line.scan(/([0-9]{1,3})/) do |number|
            if (Regexp.last_match.begin(0)..Regexp.last_match.end(0) - 1).any? { |e|  (start_index..end_index).include?(e) }
                proximity += 1
                numbers.append(number[0].to_i)
            end
        end
        lines[i+1].scan(/([0-9]{1,3})/) do |number|
            if (Regexp.last_match.begin(0)..(Regexp.last_match.end(0) - 1)).any? { |e|  (start_index..end_index).include?(e) }
                proximity += 1
                numbers.append(number[0].to_i)
            end
        end
        lines[i-1].scan(/([0-9]{1,3})/) do |number|
            if (Regexp.last_match.begin(0)..Regexp.last_match.end(0) - 1).any? { |e|  (start_index..end_index).include?(e) }
                proximity += 1
                numbers.append(number[0].to_i)
            end
        end
        if proximity == 2
            result += numbers.inject(:*)
        end
    end
end
pp result
