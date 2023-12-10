require 'pry'

file = File.open("input.txt")
lines = file.read.split("\n")
og_lines = lines.dup
result = 0
lines.each_with_index do |line, i|
    bla = line.gsub(/Card\s*[0-9]*:/, "").split("|")
    card = line.match(/Card\s*([0-9]{1,3}):/)[1].to_i
    wins = (bla[0].split(" ") & bla[1].split(" ")).length
    (0..(wins - 1)).each do |w|
        lines.append(og_lines[card + w])
    end
end
pp lines.length