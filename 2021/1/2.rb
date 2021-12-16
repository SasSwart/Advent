content = File.read('input')
              .split("\n").map(&:to_i)

content = content[0..-3].zip(content[1..-2], content[2..-1]).map(&:sum)

p(content[0..-2].zip(content[1..-1])
  .map { |p| p[0].to_i - p[1].to_i }
  .select(&:negative?).length)