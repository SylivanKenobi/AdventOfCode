def vaildate_byr(byr)
    !byr.nil? && byr.to_i.between?(1920, 2002)
end

def validate_iyr(iyr)
    !iyr.nil? && iyr.to_i.between?(2010, 2020)
end

def validate_eyr(eyr)
    !eyr.nil? && eyr.to_i.between?(2020, 2030)
end

def vaildate_hgt(hgt)
    return false if hgt.nil?
    hgt = hgt.split(/(?<=\d)(?=[A-Za-z])/)
    return false if hgt[1].nil?
    valid = false
    if hgt[1].eql?("cm")
        valid = hgt[0].to_i.between?(150, 193)
    elsif hgt[1].eql?("in")
        valid = hgt[0].to_i.between?(59, 76)
    end
    valid
end

def validate_hcl(hcl)
    hcl =~ /^#(?:[0-9a-fA-F]{6})$/ && !hcl.nil?
end

def validate_ecl(ecl)
    ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].include?(ecl) && !ecl.nil?
end

def validate_pid(pid)
    pid =~ /^\d{9}$/ && !pid.nil?
end

file = File.open('input.txt')
lines = file.read.split("\n\n")
passports = []
lines.each do |line|
    passport_entry = {}
    line.gsub("\n", " ").gsub(" ", ",").split(",").each do |entry|
        passport_entry.merge!(Hash[entry.split(':')[0], entry.split(':')[1]])
    end
    passports.append(passport_entry)
end
valid_passwords = 0
passports.each do |passport|
    if vaildate_byr(passport["byr"]) && vaildate_hgt(passport["hgt"]) && validate_ecl(passport["ecl"]) && validate_eyr(passport["eyr"]) && validate_hcl(passport["hcl"]) && validate_iyr(passport["iyr"]) && validate_pid(passport["pid"])
        valid_passwords += 1
    end
end
pp valid_passwords