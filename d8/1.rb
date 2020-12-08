require 'set'

visited = [].to_set
register = 0

lines = $<.each_line.to_a
puts "lines len: #{lines.length}"
i = 0
while i < lines.length do
    r = lines[i]
    if visited.include?i
        puts "already visited"
        break
    end
    visited << i

    op, val = r.split
    # puts "#{op} #{val} #{register}"

    if op == "acc"
        register += val.to_i
    end
    
    if op == "jmp"
        i += val.to_i
        next
    end

    i+=1
end

puts "register is #{register}"
