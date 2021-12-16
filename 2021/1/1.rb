content = File.read('input')
              .split("\n")

p(content[0..-2].zip(content[1..-1])
  .map { |p| p[0].to_i - p[1].to_i }
  .select(&:negative?).length)
