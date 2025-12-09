#! /usr/bin/env ruby

class Diagram
  def initialize
    @rolls = Set.new
  end

  def add_roll(x, y)
    @rolls.add([x, y])
  end

  def accessible_rolls
    @rolls.select do |roll|
      (-1..1).map do |x|
        (-1..1).map do |y|
          next 0 if x == 0 && y == 0
          next 0 unless @rolls.include?([roll[0] + x, roll[1] + y])

          1
        end.sum
      end.sum < 4
    end
  end

  def remove_accessible_rolls
    removed = 0

    loop do
      accessible = accessible_rolls
      return removed if accessible.count == 0

      @rolls -= accessible
      removed += accessible.count
    end
  end
end

def load_diagram
  diagram = Diagram.new

  File.open('input.txt').each_line.each_with_index do |line, y|
    line.chars.each_with_index do |char, x|
      diagram.add_roll(x, y) if char == '@'
    end
  end

  diagram
end

diagram = load_diagram

puts "accessible rolls: #{diagram.accessible_rolls.count}"
puts "rolls that can be removed: #{diagram.remove_accessible_rolls}"
