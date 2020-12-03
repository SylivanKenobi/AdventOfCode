file = File.open("input.txt")
lines = file.read.split("\n")
vaild_passwords = 0
lines.map do |line|
    parts = line.split(" ")
    letter = parts[1][0]
    numbers = parts[0].split("-")
    if (parts[2][numbers[0].to_i - 1].eql?(letter) && !parts[2][numbers[1].to_i - 1].eql?(letter)) || (!parts[2][numbers[0].to_i - 1].eql?(letter) && parts[2][numbers[1].to_i - 1].eql?(letter))
        vaild_passwords += 1
    end
end
pp vaild_passwords
