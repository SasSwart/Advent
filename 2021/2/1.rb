content = File.read('input')
              .split("\n").map { |line| line.split(' ') }

position = content.reduce([0, 0, 0]) do |acc, i|
  case i[0]
  when 'forward'
    [acc[0] + i[1].to_i, acc[1] + i[1].to_i * acc[2], acc[2]]
  when 'down'
    [acc[0], acc[1], acc[2] + i[1].to_i]
  when 'up'
    [acc[0], acc[1], acc[2] - i[1].to_i]
  end
end

p position[0] * position[1]
