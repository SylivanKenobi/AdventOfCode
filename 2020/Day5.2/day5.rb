def get_row(row_letters)
    lower_row = 0
    higher_row = 127
    row_letters.split('').each do |letter|
        if letter.eql?('F')
            higher_row = ((higher_row + lower_row) / 2).floor
        elsif letter.eql?('B')
            lower_row = ((higher_row + lower_row) / 2).ceil
        end
    end
    return higher_row
end

def get_column(column_letters)
    lower_column =  0
    higher_column = 7
    column_letters.split('').each do |letter|
        if letter.eql?('L')
            higher_column = ((higher_column + lower_column) / 2).floor
        elsif letter.eql?('R')
            lower_column = ((higher_column + lower_column) / 2).ceil
        end
    end
    return higher_column
end

file = File.open('input.txt')
lines = file.read.split("\n")
ids = []
lines.each do |line|
    row = get_row(line[0..6])
    column = get_column(line[7..-1])
    id = row * 8 + column
    ids << id
end
ids.sort!
ids.each_with_index do |id, i|
    if ids[i+1] - id == 2
        pp id
        pp ids[i+1]
    end
end