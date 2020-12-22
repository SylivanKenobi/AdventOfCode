file = File.open('input.txt')
input_lines = file.read.split("\n\n")
counter = 0
input_lines.each do |input_line|
    map = {}
    block = input_line.split.map { |l| l.split('') }
    block[0].each { |l|  map.merge!("#{l}" => 1) }
    block[1, block.length - 1].each do |line|
        line.each { |l| map[l] += 1 if map.key?(l) }
    end
    map.each { |k, v| counter += 1 if v == block.length }
end
pp counter