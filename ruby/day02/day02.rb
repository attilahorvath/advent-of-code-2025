#! /usr/bin/env ruby

def sum_all_invalid_ids(twice_only: true)
  File.read('input.txt').split(',').map do |range|
    min, max = range.split('-').map(&:to_i)

    (min..max).select do |id|
      id_str = id.to_s

      if twice_only
        id_str[...(id_str.length / 2)] == id_str[(id_str.length / 2)..]
      else
        id_str[/\A(\d+)\1+\z/]
      end
    end.sum
  end.sum
end

puts "sum with sequences repeated twice only: #{sum_all_invalid_ids}"
puts "sum with all sequences: #{sum_all_invalid_ids(twice_only: false)}"
