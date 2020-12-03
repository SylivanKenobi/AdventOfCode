file = File.open("input.txt")
numbers = file.read.split
result1 = ''
result2 = ''
numbers.each_with_index do |number1, i|
    numbers.each_with_index do |number2, j|
        if (number1.to_i + number2.to_i) == 2020
            result1 = number1
            result2 = number2
            break
        end
    end
end
puts result1.to_i * result2.to_i