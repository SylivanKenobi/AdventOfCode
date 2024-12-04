def boot(command_list)
    i = 0
    result = 0
    command_history = []
    while i < command_list.length
        if command_history.include?(i)
            result = 0
            break
        end
        command = command_list[i].scan(/[a-zA-Z]+/)[0]
        number = command_list[i].scan(/[0-9]+/)[0].to_i
        operator = command_list[i].scan(/\+|\-/)[0]
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
    return result
end

file = File.open('input.txt')
input = file.read.split("\n")
duplicate = input.map(&:clone)

result = 0
nop = input.each_index.select{|i| input[i].include?('nop')}
jmp = input.each_index.select{|i| input[i].include?('jmp')}

i = 0
for i in nop
    number = input[i].scan(/\+[0-9]+|\-[0-9]+/)[0]
    duplicate[i] = "jmp #{number}"
    result = boot(duplicate)
    duplicate = input.map(&:clone)
    break if result > 0
    i = i + 1
end

i = 0
for i in jmp
    number = input[i].scan(/\+[0-9]+|\-[0-9]+/)[0]
    duplicate[i] = "nop #{number}"
    result = boot(duplicate)
    duplicate = input.map(&:clone)
    break if result > 0
    i = i + 1
end

pp result