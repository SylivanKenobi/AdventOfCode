file = File.open("input.txt")
lines = file.read.split
claibration = 0
lines.each do |line|
    numbers = line.tr('^0-9', '')
    coordinate = "#{numbers[0]}#{numbers[-1]}"
    claibration += coordinate.to_i
end
pp claibration
