file = File.open("input.txt")
lines = file.read.split("\n")

max_red = 12
max_green = 13
max_blue = 14

games = []
lines.each do |line|
    parts = line.split(": ")
    id = parts[0].tr('^0-9', '').to_i
    reveals_dirty = parts[1].split("; ")
    reveals = []
    reveals_dirty.each do |r|
        red = r.split(", ").select {|i| i =~ /red/}[0]&.tr('^0-9', '').to_i
        blue = r.split(", ").select {|i| i =~ /blue/}[0]&.tr('^0-9', '').to_i
        green = r.split(", ").select {|i| i =~ /green/}[0]&.tr('^0-9', '').to_i
        reveals.append({red: red, blue: blue, green: green})
    end
    games.append({id: id, reveals: reveals})
end
result = 0
games.each do |game|
    lowest_red = 0
    lowest_blue = 0
    lowest_green = 0
    game[:reveals].each do |reveal|
        lowest_red = reveal[:red] > lowest_red ? reveal[:red] : lowest_red
        lowest_blue = reveal[:blue] > lowest_blue ? reveal[:blue] : lowest_blue
        lowest_green = reveal[:green] > lowest_green ? reveal[:green] : lowest_green
    end
    result += (lowest_red * lowest_blue * lowest_green)
end
pp result