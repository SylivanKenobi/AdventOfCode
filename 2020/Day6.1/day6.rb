file = File.open('input.txt')
lines = file.read.split("\n\n")
yesses = 0
lines.each do |line|
    yesses += line.gsub("\n", "").split('').uniq.length
end
pp yesses