sid_max = 0
$<.each_line {|r|
    row = r[0..6].tr("FB", "01").to_i(2)
    column = r[7..-1].tr("LR", "01").to_i(2)
    sid = row*8+column
    sid_max = sid if sid > sid_max
}
puts "max seat id: #{sid_max}"
