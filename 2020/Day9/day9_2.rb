file = File.open('input.txt')
numbers = file.read.split("\n").map(&:to_i)

bad_num = 258585477
i = 0

numbers.each_with_index do |num, i|
    sum = 0
    for j in i..numbers.length-1 do
        sum = sum + numbers[j]
        if sum == bad_num
            res = numbers[i..j].sort
            pp res[0] + res[-1]
            exit
        elsif sum > bad_num 
            break
        end
    end
end
