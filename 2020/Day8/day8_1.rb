file = File.open('input.txt')
input = file.read.split("\n")

i = 0
result = 0
command_history = []
while i < input.length
    pp "#{input[i]} #{i} #{result}"
    if command_history.include?(i)
        break
    end
    command = input[i].scan(/[a-zA-Z]+/)[0]
    number = input[i].scan(/[0-9]+/)[0].to_i
    operator = input[i].scan(/\+|\-/)[0]
    case command
    when "acc"
        command_history.append(i)
        if operator == "+"
            result = result + number
        else
            result = result - number
        end
        i = i + 1
    when "jmp"
        command_history.append(i)
        if operator == "+"
            i = i + number
        else
            i = i - number
        end
    when "nop"
        command_history.append(i)
        i = i + 1
    end
end
pp result