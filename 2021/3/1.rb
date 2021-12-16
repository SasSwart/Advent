numbers = File.read('input')
              .split("\n")

def gamma(numbers, acc = nil)
  return acc if numbers.empty?

  acc ||= 0

  numbers[0].chars.each_index do |i|
    acc += numbers[0][i].to_i
  end

  gamma(numbers[1..-1], acc)
end

gamma_bits = gamma(numbers)

p gamma_bits
