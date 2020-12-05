seats = (1..127*8+7).to_a
$<.each_line {|r|
    row = r[0..6].tr("FB", "01").to_i(2)
    column = r[7..-1].tr("LR", "01").to_i(2)
    sid = row*8+column
    seats.delete sid
}
puts "my seat: #{seats.each_cons(3).filter_map {|a,b,c|   b if b-a>1&&c-b>1}.first}"