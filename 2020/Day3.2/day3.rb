file = File.open("input.txt")
lines = file.read.split("\n")
x = 7
trees = 0
i = 0
for lines in lines.length do
    
end
lines.each do |line|
    trees += 1 if line[i].eql?('#')
    i = (i + x) >= line.length - 1 ? (i + x) - line.length : i + x
end
pp trees

# 3:1 294
# 1:1 75
# 5:1 79
# 7:1 85
