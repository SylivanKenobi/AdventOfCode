require 'pry'

file = File.open("input.txt")
lines = file.read.split("\n")
result = 0
lines.each do |line|
    pp line
    bla = line.gsub!(/Card\s*[0-9]*:/, "").split("|")
    points = 0
    (0..((bla[0].split(" ") & bla[1].split(" ")).length - 1)).each do |e| 
        pp points
        points = e == 0 ? 1 : points * 2 
    end
    result += points
end
pp result