def get_parents(data, bag_color, top_bags)
    size = top_bags.length
    data.each do |line|
        if line =~ /^.*contain.*#{bag_color} bag.*$/
            top_bags << line.split(" ")[0,2].join(' ')}
            top_bags.uniq!
        end
    end
    return top_bags if size == top_bags.length
    top_bags.each do |bag|
        top_bags + get_parents(data, bag, top_bags)
    end
end

file = File.open('input.txt')
input = file.read.split("\n")
pp get_parents(input, 'shiny gold', []).length


