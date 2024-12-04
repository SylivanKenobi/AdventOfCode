def valid_number(preamble, number)
    additions = []
    preamble.each_with_index do |num, i|
        for j in i..preamble.length-1 do
            additions.append(num + preamble[j])
        end 
    end
    return additions.include?(number)
end

file = File.open('input.txt')
numbers = file.read.split("\n").map(&:to_i)

pre = 25
i = pre
while i < numbers.length 
    if !valid_number(numbers[i-pre..i-1], numbers[i])
        pp numbers[i]
        break
    end
    i = i + 1
end
