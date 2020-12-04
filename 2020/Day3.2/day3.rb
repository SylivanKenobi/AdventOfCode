file = File.open("input.txt")
lines = file.read.split("\n")
x = 1
trees = 0
i = 0
lines.each_with_index do |line, j|
    if j % 2 == 0
        trees += 1 if line[i].eql?('#')
        i = (i + x) >= line.length - 1 ? (i + x) - line.length : i + x
    end
end
pp trees

# 3:1 294
# 1:1 75
# 5:1 79
# 7:1 85
# 1:2 39
