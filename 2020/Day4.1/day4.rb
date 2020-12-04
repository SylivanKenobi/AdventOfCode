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
vaild_passports = 0
passports.each do |passport|
    vaild_passports += 1 if ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"].all? { |key| passport.has_key?(key) }
end
pp vaild_passports