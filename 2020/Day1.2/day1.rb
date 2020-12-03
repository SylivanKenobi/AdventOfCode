file = File.open("input.txt")
numbers = file.read.split
result1 = ''
result2 = ''
result3 = ''
numbers.each do |number1|
    numbers.each do |number2|
        numbers.each do |number3|
            if (number1.to_i + number2.to_i + number3.to_i) == 2020
                result1 = number1
                result2 = number2
                result3 = number3
                break
            end
        end
    end
end
puts result1.to_i * result2.to_i * result3.to_i