require 'set'

$lines = $<.each_line.to_a
puts "lines len: #{$lines.length}"

def iterate_lines(corrupted=0)
    i = 0
    visited = [].to_set
    register = 0

    while i < $lines.length do
        r = $lines[i]
        if visited.include?i
            break
        end
        visited << i

        op, val = r.split
        # puts "#{op} #{val} #{register}"

        if i == corrupted
            op = "jmp" if op == "nop"
            op = "nop" if op == "jmp"
        end

        if op == "acc"
            register += val.to_i
        end

        
        if op == "jmp"
            i += val.to_i
            break if val.to_i == 0
            next
        end

        i+=1
        if i == $lines.length-1
            puts "#{corrupted} ound the corrupted one"
            break
        end
    end

    puts "#{corrupted} register is #{register}"

end

for corrupted in (1..$lines.length)
    iterate_lines corrupted
end
