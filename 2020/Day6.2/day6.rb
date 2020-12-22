file = File.open('input.txt')
input_lines = file.read.split("\n\n")
counter = 0
input_lines.each do |input_line|
    split_line = input_line.split.map { |l| l.split('') }
    map = {}
    split_line[0].each { |l|  map.merge!("#{l}" => 0) }
    split_line.each do |variable|
        variable.each do |letter|
            if map.key?(letter)
                map[letter] += 1
            end
        end
    end
    map.each do |k, v|
        if v == split_line.length
            counter += 1
        end
    end
end
pp counter