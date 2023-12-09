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
    impossible = false
    game[:reveals].each do |reveal|
        if reveal[:red] > max_red
            impossible = true
        end
        if reveal[:blue] > max_blue
            impossible = true
        end
        if reveal[:green] > max_green
            impossible = true
        end
    end
    unless impossible
        result += game[:id]
    end
end
pp result