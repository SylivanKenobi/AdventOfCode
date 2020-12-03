file = File.open("input.txt")
lines = file.read.split("\n")
vaild_passwords = 0
lines.map do |line|
    parts = line.split(" ")
    letter = parts[1][0]
    numbers = parts[0].split("-")
    counter = 0
    parts[2].split('').each do |part|
        if part == letter
            counter += 1
        end
    end
    if counter.between?(numbers[0].to_i, numbers[1].to_i)
        vaild_passwords += 1
    end
end
pp vaild_passwords
