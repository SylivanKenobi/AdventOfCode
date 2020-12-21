file = File.open('input.txt')
lines = file.read.split("\n")
ids = []
lines.each do |line|
    rows = {low: 0, high: 127}
    columns =  {low: 0, high: 7}
    row = 0
    column = 0
    line.split('').each do |letter|
        leftover_rows = (rows[:low]..rows[:high]).to_a.length / 2
        leftover_columns = (columns[:low]..columns[:high]).to_a.length / 2
        if letter.eql?('F')
            row = rows[:low] if leftover_rows == 1
            rows = {low: rows[:low], high: rows[:high] - leftover_rows} if leftover_rows > 1
        elsif letter.eql?('B')
            row = rows[:low] if leftover_rows == 1
            rows = {low: rows[:low] + leftover_rows, high: rows[:high]} if leftover_rows > 1
        elsif letter.eql?('L')
            column = columns[:low] if leftover_columns == 1
            columns = {low: columns[:low], high: columns[:high] - leftover_columns} if leftover_columns > 1
        elsif letter.eql?('R')
            column = columns[:high] if leftover_columns == 1
            columns = {low: columns[:low] + leftover_columns, high: columns[:high]} if leftover_columns > 1
        end
    end
    id = row * 8 + column
    pp "#{line} #{row} #{column} #{id}"
    ids << id
end
pp ids.sort[0]
