file = File.open("input.txt")
lines = file.read.split("\n")
x = 3
trees = 0
i = 0
lines.each do |line|
    trees += 1 if line[i].eql?('#')
    i = (i + x) >= line.length - 1 ? (i + x) - line.length : i + x
end
pp trees