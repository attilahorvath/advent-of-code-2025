#! /usr/bin/env ruby

def output(ratings, batteries)
  max_outputs = Array.new(ratings.length, 0)

  ratings.chars.map(&:to_i).each do |rating|
    curr_batteries = batteries
    
    while curr_batteries >= 1
      max_outputs[curr_batteries] = [max_outputs[curr_batteries], max_outputs[curr_batteries - 1] * 10 + rating].max
      curr_batteries -= 1
    end
  end

  max_outputs[batteries]
end

def total_output(batteries: 2)
  File.open('input.txt').each_line.map do |line|
    output(line.chomp, batteries)
  end.sum
end

puts "total output with 2 batteries: #{total_output}"
puts "total output with 12 batteries: #{total_output(batteries: 12)}"
