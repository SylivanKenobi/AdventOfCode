def parse_contents(contents)
    output = []

    return output if contents == "no other bags."

    contents.split(", ").each do |content|
        parts = content.split(" ")
        count = parts[0].to_i
        bag_name = parts[1..2].join(" ")
        output.push({ :name => bag_name, :count => count })
    end

    return output
end

def add_children(bags)
    bag_count = 0
    bags_to_check = ["shiny gold"]

    while !bags_to_check.empty?
        bag_to_check = bags_to_check.shift
        for bag in bags[bag_to_check]
            bag_count += bag[:count]
            bags_to_check += [bag[:name]] * bag[:count]
        end
    end

    return bag_count
end


file = File.open('input.txt')
input = file.read.split("\n")
bags = {}
children_count = {}
input.each do |line|
    bag = line.chomp.split(" ")
    parent_bag = bag[0..1].join(" ")
    contents = parse_contents(bag[4..].join(" "))
    bags[parent_bag] = contents
end
pp add_children(bags)


